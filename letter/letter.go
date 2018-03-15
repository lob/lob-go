package letter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"reflect"

	"github.com/lob/lob-go/address"
	"github.com/lob/lob-go/api"
	"github.com/lob/lob-go/scaffold"
	"github.com/lob/lob-go/thumbnail"
	"github.com/lob/lob-go/tracking"
)

const (
	basePath = "v1/letters"
)

// Letter represents a letter object as returned from the Lob API.
type Letter struct {
	ID                   string                `json:"id"`
	Description          *string               `json:"description"`
	To                   address.Address       `json:"to"`
	From                 address.Address       `json:"from"`
	Color                bool                  `json:"color"`
	DoubleSided          bool                  `json:"double_sided"`
	AddressPlacement     string                `json:"address_placement"`
	ReturnEnvelope       bool                  `json:"return_envelope"`
	PerforatedPage       *int                  `json:"perforated_page"`
	ExtraService         *string               `json:"extra_service"`
	MailType             string                `json:"mail_type"`
	URL                  string                `json:"url"`
	TemplateID           *string               `json:"template_id"`
	TemplateVersionID    *string               `json:"template_version_id"`
	Carrier              string                `json:"carrier"`
	TrackingNumber       *string               `json:"tracking_number"`
	ExpectedDeliveryDate string                `json:"expected_delivery_date"`
	SendDate             string                `json:"send_date"`
	TrackingEvents       []tracking.Event      `json:"tracking_events"`
	Thumbnails           []thumbnail.Thumbnail `json:"thumbnails"`
	DateCreated          string                `json:"date_created"`
	DateModified         string                `json:"date_modified"`
	Metadata             map[string]string     `json:"metadata"`
	Deleted              bool                  `json:"deleted"`
}

func (l Letter) String() string {
	ltr, _ := json.Marshal(&l)
	return string(ltr)
}

// Handler specifies methods for interacting with letters via the API.
type Handler interface {
	Create(context.Context, *scaffold.Letter) (*Letter, error)
	Retrieve(context.Context, string) (*Letter, error)
	Cancel(context.Context, string) error
	List(context.Context, int, int) ([]Letter, error)
}

// Client implements the LetterHandler interface and manages client state.
type Client struct {
	client api.Client
}

// NewClient creates a LetterClient
func NewClient(client api.Client) *Client {
	return &Client{client}
}

// Create attempts to create a new Letter.
func (lc *Client) Create(ctx context.Context, scfld *scaffold.Letter) (*Letter, error) {
	var req *http.Request
	var err error

	if reflect.TypeOf(scfld.File) == reflect.TypeOf([]byte{}) {
		req, err = lc.newFileUploadRequest(ctx, scfld)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = lc.client.NewRequest(ctx, "POST", basePath, scfld)
		if err != nil {
			return nil, err
		}
	}

	resp, err := lc.client.Do(req)
	if err != nil {
		return nil, err
	}

	ltr := &Letter{}
	err = api.Deserialize(resp, ltr)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	return ltr, err
}

// Retrieve gets a Letter corresponding to the given id.
func (lc *Client) Retrieve(ctx context.Context, id string) (*Letter, error) {
	relPath := fmt.Sprintf("%s/%s", basePath, id)

	resp, err := lc.client.Get(ctx, relPath)
	if err != nil {
		return nil, err
	}

	ltr := &Letter{}
	err = api.Deserialize(resp, ltr)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	return ltr, err
}

// Cancel attempts to cancel the specified Letter.
func (lc *Client) Cancel(ctx context.Context, id string) error {
	relPath := fmt.Sprintf("%s/%s", basePath, id)

	resp, err := lc.client.Delete(ctx, relPath)
	if err != nil {
		return err
	}

	err = resp.Body.Close()
	return err
}

// List returns all Letters subject to the given limit and offset.
func (lc *Client) List(ctx context.Context, limit int, offset int) ([]Letter, error) {
	relPath := fmt.Sprintf("%s/?limit=%d&offset=%d", basePath, limit, offset)

	resp, err := lc.client.Get(ctx, relPath)
	if err != nil {
		return nil, err
	}

	msg := new(struct {
		Data  []Letter `json:"data"`
		Count int      `json:"count"`
	})

	err = api.Deserialize(resp, msg)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	return msg.Data, err
}

func (lc *Client) newFileUploadRequest(ctx context.Context, scfld *scaffold.Letter) (*http.Request, error) {

	formFields, err := scaffold.MarshalAsFormValues(scfld)
	if err != nil {
		return nil, err
	}

	buffer := &bytes.Buffer{}
	w := multipart.NewWriter(buffer)

	err = api.WriteFormFile(scfld.File, "file", w)
	if err != nil {
		return nil, err
	}

	// Add the other fields
	for key, val := range formFields {
		_ = w.WriteField(key, val)
	}

	err = w.Close()
	if err != nil {
		return nil, err
	}

	req, err := lc.client.NewMultiformRequest(ctx, basePath, w.FormDataContentType(), buffer)
	return req, err
}
