package lob

import (
	"github.com/lob/lob-go/address"
	"github.com/lob/lob-go/api"
	"github.com/lob/lob-go/bankaccount"
	"github.com/lob/lob-go/check"
	"github.com/lob/lob-go/letter"
	"github.com/lob/lob-go/postcard"
)

// Client wraps resource-specific interfaces.
type Client struct {
	Address     address.Handler
	BankAccount bankaccount.Handler
	Check       check.Handler
	Letter      letter.Handler
	Postcard    postcard.Handler
}

// New creates a new Lob Client using the specified options.
func New(options *api.Options) *Client {
	client := api.NewJSONClient(options)
	return &Client{
		Address:     address.NewClient(client),
		BankAccount: bankaccount.NewClient(client),
		Check:       check.NewClient(client),
		Letter:      letter.NewClient(client),
		Postcard:    postcard.NewClient(client),
	}
}
