package bankaccount

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lob/lob-go/api"
	"github.com/lob/lob-go/scaffold"
)

const (
	basePath = "v1/bank_accounts"
)

// BankAccount represents a bank account object as returned from the Lob API.
type BankAccount struct {
	ID            string            `json:"id"`
	AccountNumber string            `json:"account_number"`
	AccountType   string            `json:"account_type"`
	BankName      string            `json:"bank_name"`
	RoutingNumber string            `json:"routing_number"`
	Signatory     string            `json:"signatory"`
	SignatoryURL  *string           `json:"signatory_url"`
	Verified      bool              `json:"verified"`
	DateCreated   string            `json:"date_created"`
	DateModified  string            `json:"date_modified"`
	Description   *string           `json:"description"`
	Metadata      map[string]string `json:"metadata"`
	Deleted       *bool             `json:"deleted"`
}

func (ba BankAccount) String() string {
	acct, _ := json.Marshal(&ba)
	return string(acct)
}

// Handler specifies methods for interacting with bank accounts via the API.
type Handler interface {
	Create(context.Context, *scaffold.BankAccount) (*BankAccount, error)
	Retrieve(context.Context, string) (*BankAccount, error)
	Delete(context.Context, string) error
	Verify(context.Context, string, []int) (*BankAccount, error)
	List(context.Context, int, int) ([]BankAccount, error)
}

// Client implements the BankAccountHandler interface and manages client state.
type Client struct {
	client api.Client
}

// NewClient creates a BankAccountClient
func NewClient(client api.Client) *Client {
	return &Client{client}
}

// Create attempts to create a new BankAccount.
func (bac *Client) Create(ctx context.Context, scaffold *scaffold.BankAccount) (*BankAccount, error) {
	resp, err := bac.client.Post(ctx, basePath, scaffold)
	if err != nil {
		return nil, err
	}

	acct := &BankAccount{}
	err = api.Deserialize(resp, acct)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	return acct, err
}

// Retrieve gets a BankAccount corresponding to the given id.
func (bac *Client) Retrieve(ctx context.Context, id string) (*BankAccount, error) {
	relPath := fmt.Sprintf("%s/%s", basePath, id)

	resp, err := bac.client.Get(ctx, relPath)
	if err != nil {
		return nil, err
	}

	acct := &BankAccount{}
	err = api.Deserialize(resp, acct)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	return acct, err
}

// Delete attempts to mark the specified BankAccount as deleted.
func (bac *Client) Delete(ctx context.Context, id string) error {
	relPath := fmt.Sprintf("%s/%s", basePath, id)

	resp, err := bac.client.Delete(ctx, relPath)
	if err != nil {
		return err
	}

	err = resp.Body.Close()
	return err
}

// Verify attempts to verify a bank account.
func (bac *Client) Verify(ctx context.Context, id string, amounts []int) (*BankAccount, error) {
	relPath := fmt.Sprintf("%s/%s/verify", basePath, id)

	payload := new(struct {
		Amounts []int `json:"amounts"`
	})
	payload.Amounts = amounts

	resp, err := bac.client.Post(ctx, relPath, payload)
	if err != nil {
		return nil, err
	}

	acct := &BankAccount{}
	err = api.Deserialize(resp, acct)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	return acct, err
}

// List returns all BankAccounts subject to the given limit and offset.
func (bac *Client) List(ctx context.Context, limit int, offset int) ([]BankAccount, error) {
	relPath := fmt.Sprintf("%s/?limit=%d&offset=%d", basePath, limit, offset)

	resp, err := bac.client.Get(ctx, relPath)
	if err != nil {
		return nil, err
	}

	msg := new(struct {
		Data  []BankAccount `json:"data"`
		Count int           `json:"count"`
	})

	err = api.Deserialize(resp, msg)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	return msg.Data, err
}
