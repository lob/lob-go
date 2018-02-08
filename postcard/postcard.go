package postcard

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"reflect"

	"github.com/lob/lob-go/address"
	"github.com/lob/lob-go/scaffold"
	"github.com/lob/lob-go/thumbnail"
	"github.com/lob/lob-go/tracking"

	"github.com/lob/lob-go/api"
)

const (
	basePath = "v1/postcards"
)

// Postcard represents a postcard object as returned from the Lob API.
type Postcard struct {
	ID                     string                `json:"id"`
	Description            *string               `json:"description"`
	To                     address.Address       `json:"to"`
	From                   *address.Address      `json:"from"`
	URL                    string                `json:"url"`
	FrontTemplateID        *string               `json:"front_template_id"`
	BackTemplateID         *string               `json:"back_template_id"`
	FrontTemplateVersionID *string               `json:"front_template_version_id"`
	BackTemplateVersionID  *string               `json:"back_template_version_id"`
	Carrier                string                `json:"carrier"`
	Size                   string                `json:"size"`
	MailType               string                `json:"mail_type"`
	ExpectedDeliveryDate   string                `json:"expected_delivery_date"`
	SendDate               string                `json:"send_date"`
	TrackingEvents         []tracking.Event      `json:"tracking_events"`
	Thumbnail              []thumbnail.Thumbnail `json:"thumbnail"`
	DateCreated            string                `json:"date_created"`
	DateModified           string                `json:"date_modified"`
	Metadata               map[string]string     `json:"metadata"`
	Deleted                bool                  `json:"deleted"`
}

func (p Postcard) String() string {
	psc, _ := json.Marshal(&p)
	return string(psc)
}

// Handler specifies methods for interacting with postcards via the API.
type Handler interface {
	Create(context.Context, *scaffold.Postcard) (*Postcard, error)
	Retrieve(context.Context, string) (*Postcard, error)
	Cancel(context.Context, string) error
	List(context.Context, int, int) ([]Postcard, error)
}

// Client implements the Handler interface and manages client state.
type Client struct {
	client api.Client
}

// NewClient creates a Client
func NewClient(client api.Client) *Client {
	return &Client{client}
}

// Create attempts to create a new Postcard.
func (pc *Client) Create(ctx context.Context, scfld *scaffold.Postcard) (*Postcard, error) {
	var req *http.Request
	var err error

	if reflect.TypeOf(scfld.Front) == reflect.TypeOf([]byte{}) || reflect.TypeOf(scfld.Back) == reflect.TypeOf([]byte{}) {
		req, err = pc.newFileUploadRequest(ctx, scfld)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = pc.client.NewRequest(ctx, "POST", basePath, scfld)
		if err != nil {
			return nil, err
		}
	}

	resp, err := pc.client.Do(req)
	if err != nil {
		return nil, err
	}

	outPsc := &Postcard{}
	err = api.Deserialize(resp, outPsc)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	return outPsc, err
}

// Retrieve gets a Postcard corresponding to the given id.
func (pc *Client) Retrieve(ctx context.Context, id string) (*Postcard, error) {
	relPath := fmt.Sprintf("%s/%s", basePath, id)

	resp, err := pc.client.Get(ctx, relPath)
	if err != nil {
		return nil, err
	}

	psc := &Postcard{}
	err = api.Deserialize(resp, psc)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	return psc, err
}

// Cancel attempts to cancel the specified Postcard.
func (pc *Client) Cancel(ctx context.Context, id string) error {
	relPath := fmt.Sprintf("%s/%s", basePath, id)

	req, err := pc.client.NewRequest(ctx, "DELETE", relPath, nil)
	if err != nil {
		return err
	}

	resp, err := pc.client.Do(req)
	if err != nil {
		return err
	}

	err = resp.Body.Close()
	return err
}

// List returns all Postcards subject to the given limit and offset.
func (pc *Client) List(ctx context.Context, limit int, offset int) ([]Postcard, error) {
	relPath := fmt.Sprintf("%s/?limit=%d&offset=%d", basePath, limit, offset)

	resp, err := pc.client.Get(ctx, relPath)
	if err != nil {
		return nil, err
	}

	msg := new(struct {
		Data  []Postcard `json:"data"`
		Count int        `json:"count"`
	})

	err = api.Deserialize(resp, msg)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	return msg.Data, err
}

func (pc *Client) newFileUploadRequest(ctx context.Context, scfld *scaffold.Postcard) (*http.Request, error) {
	formFields, err := scaffold.MarshalAsFormValues(scfld)
	if err != nil {
		return nil, err
	}

	buffer := &bytes.Buffer{}
	w := multipart.NewWriter(buffer)

	api.WriteFormFile(scfld.Front, "front", w)
	if err != nil {
		return nil, err
	}

	api.WriteFormFile(scfld.Back, "back", w)
	if err != nil {
		return nil, err
	}

	// Add non-file fields
	for key, val := range formFields {
		_ = w.WriteField(key, val)
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}

	req, err := pc.client.NewMultiformRequest(ctx, basePath, w.FormDataContentType(), buffer)
	return req, err
}
