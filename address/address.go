package address

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lob/lob-go/api"
	"github.com/lob/lob-go/scaffold"
)

const (
	basePath = "v1/addresses"
)

// Address represents an address object as returned from the Lob API.
type Address struct {
	ID           string            `json:"id"`
	Name         *string           `json:"name"`
	Line1        string            `json:"address_line1"`
	Line2        *string           `json:"address_line2"`
	City         *string           `json:"address_city"`
	Zip          *string           `json:"address_zip"`
	State        *string           `json:"address_state"`
	Country      string            `json:"address_country"`
	Company      *string           `json:"company"`
	Email        *string           `json:"email"`
	Phone        *string           `json:"phone"`
	Description  *string           `json:"description"`
	DateCreated  string            `json:"date_created"`
	DateModified string            `json:"date_modified"`
	Metadata     map[string]string `json:"metadata"`
	Deleted      *bool             `json:"deleted"`
}

func (a Address) String() string {
	addy, _ := json.Marshal(&a)
	return string(addy)
}

// Handler specifies methods for interacting with addresses via the API.
type Handler interface {
	Create(context.Context, *scaffold.Address) (*Address, error)
	Retrieve(context.Context, string) (*Address, error)
	Delete(context.Context, string) error
	List(context.Context, int, int) ([]Address, error)
}

// Client implements the Handler interface and manages client state.
type Client struct {
	client api.Client
}

// NewClient creates a Client
func NewClient(client api.Client) *Client {
	return &Client{client}
}

// Create creates a new Address.
func (ac *Client) Create(ctx context.Context, scaffold *scaffold.Address) (*Address, error) {
	resp, err := ac.client.Post(ctx, basePath, scaffold)
	if err != nil {
		return nil, err
	}

	addr := &Address{}
	err = api.Deserialize(resp, addr)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	return addr, err
}

// Retrieve gets an Address corresponding to the given id.
func (ac *Client) Retrieve(ctx context.Context, id string) (*Address, error) {
	relPath := fmt.Sprintf("%s/%s", basePath, id)

	resp, err := ac.client.Get(ctx, relPath)
	if err != nil {
		return nil, err
	}

	addr := &Address{}
	err = api.Deserialize(resp, addr)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	return addr, err
}

// Delete attempts to mark the specified address as deleted.
func (ac *Client) Delete(ctx context.Context, id string) error {
	relPath := fmt.Sprintf("%s/%s", basePath, id)

	resp, err := ac.client.Delete(ctx, relPath)
	if err != nil {
		return err
	}

	err = resp.Body.Close()
	return err
}

// List returns all address subject to the given limit and offset.
func (ac *Client) List(ctx context.Context, limit int, offset int) ([]Address, error) {
	relPath := fmt.Sprintf("%s/?limit=%d&offset=%d", basePath, limit, offset)

	resp, err := ac.client.Get(ctx, relPath)
	if err != nil {
		return nil, err
	}

	msg := new(struct {
		Data       []Address `json:"data"`
		Count      int       `json:"count"`
		TotalCount *int      `json:"total_count"`
	})

	err = api.Deserialize(resp, msg)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	return msg.Data, err
}
