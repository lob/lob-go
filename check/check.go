package check

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"reflect"

	"github.com/lob/lob-go/address"
	"github.com/lob/lob-go/bankaccount"
	"github.com/lob/lob-go/scaffold"
	"github.com/lob/lob-go/thumbnail"
	"github.com/lob/lob-go/tracking"

	"github.com/lob/lob-go/api"
)

const (
	basePath = "v1/checks"
)

// Check represents a check object as returned from the Lob API.
type Check struct {
	ID                           string                  `json:"id"`
	Description                  *string                 `json:"description"`
	CheckNumber                  int                     `json:"check_number"`
	Memo                         *string                 `json:"memo"`
	Amount                       float64                 `json:"amount"`
	Message                      *string                 `json:"message"`
	URL                          string                  `json:"url"`
	CheckBottomTemplateID        *string                 `json:"check_bottom_template_id"`
	AttachmentTemplateID         *string                 `json:"attachment_template_id"`
	CheckBottomTemplateVersionID *string                 `json:"check_bottom_template_version_id"`
	AttachmentTemplateVersionID  *string                 `json:"attachment_template_version_id"`
	To                           address.Address         `json:"to"`
	From                         address.Address         `json:"from"`
	BankAccount                  bankaccount.BankAccount `json:"bank_account"`
	Carrier                      string                  `json:"carrier"`
	TrackingNumber               *string                 `json:"tracking_number"`
	MailType                     string                  `json:"mail_type"`
	ExpectedDeliveryDate         string                  `json:"expected_delivery_date"`
	SendDate                     string                  `json:"send_date"`
	TrackingEvents               []tracking.Event        `json:"tracking_events"`
	Thumbnails                   []thumbnail.Thumbnail   `json:"thumbnails"`
	DateCreated                  string                  `json:"date_created"`
	DateModified                 string                  `json:"date_modified"`
	Metadata                     map[string]string       `json:"metadata"`
	Deleted                      bool                    `json:"deleted"`
}

func (p Check) String() string {
	psc, _ := json.Marshal(&p)
	return string(psc)
}

// Handler specifies methods for interacting with checks via the API.
type Handler interface {
	Create(context.Context, *scaffold.Check) (*Check, error)
	Retrieve(context.Context, string) (*Check, error)
	Cancel(context.Context, string) error
	List(context.Context, int, int) ([]Check, error)
}

// Client implements the Handler interface and manages client state.
type Client struct {
	client api.Client
}

// NewClient creates a Client
func NewClient(client api.Client) *Client {
	return &Client{client}
}

// Create attempts to create a new Check.
func (cc *Client) Create(ctx context.Context, scfld *scaffold.Check) (*Check, error) {
	var req *http.Request
	var err error

	if reflect.TypeOf(scfld.Logo) == reflect.TypeOf([]byte{}) {
		req, err = cc.newFileUploadRequest(ctx, scfld)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = cc.client.NewRequest(ctx, "POST", basePath, scfld)
		if err != nil {
			return nil, err
		}
	}

	resp, err := cc.client.Do(req)
	if err != nil {
		return nil, err
	}

	outPsc := &Check{}
	err = api.Deserialize(resp, outPsc)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	return outPsc, err
}

// Retrieve gets a Check corresponding to the given id.
func (cc *Client) Retrieve(ctx context.Context, id string) (*Check, error) {
	relPath := fmt.Sprintf("%s/%s", basePath, id)

	resp, err := cc.client.Get(ctx, relPath)
	if err != nil {
		return nil, err
	}

	psc := &Check{}
	err = api.Deserialize(resp, psc)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	return psc, err
}

// Cancel attempts to cancel the specified Check.
func (cc *Client) Cancel(ctx context.Context, id string) error {
	relPath := fmt.Sprintf("%s/%s", basePath, id)

	resp, err := cc.client.Delete(ctx, relPath)
	if err != nil {
		return err
	}

	err = resp.Body.Close()
	return err
}

// List returns all Checks subject to the given limit and offset.
func (cc *Client) List(ctx context.Context, limit int, offset int) ([]Check, error) {
	relPath := fmt.Sprintf("%s/?limit=%d&offset=%d", basePath, limit, offset)

	resp, err := cc.client.Get(ctx, relPath)
	if err != nil {
		return nil, err
	}

	msg := new(struct {
		Data  []Check `json:"data"`
		Count int     `json:"count"`
	})

	err = api.Deserialize(resp, msg)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	return msg.Data, err
}

func (cc *Client) newFileUploadRequest(ctx context.Context, scfld *scaffold.Check) (*http.Request, error) {

	formFields, err := scaffold.MarshalAsFormValues(scfld)
	if err != nil {
		return nil, err
	}

	buffer := &bytes.Buffer{}
	w := multipart.NewWriter(buffer)

	err = api.WriteFormFile(scfld.Logo, "logo", w)
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

	req, err := cc.client.NewMultiformRequest(ctx, basePath, w.FormDataContentType(), buffer)
	return req, err
}
