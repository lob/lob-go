package testing

import (
	"net/http"
	"net/http/httptest"

	"github.com/lob/lob-go/api"
)

// NewClient creates a JSONClient for use in integration tests.
func NewClient() *api.JSONClient {
	opt := &api.Options{
		APIKey:  "test_cc4943f82e2883e1f94b66d2c8e9d95b2d2",
		MaxQPS:  5,
		BaseURL: "https://api.lob.com",
	}
	return api.NewJSONClient(opt)
}

// NewSrvAndClient creates an httptest.Server and an api.Client for use in unit tests.
func NewSrvAndClient(handler func(http.ResponseWriter, *http.Request)) (*httptest.Server, api.Client) {
	srv := httptest.NewServer(http.HandlerFunc(handler))
	client := api.NewJSONClient(&api.Options{
		BaseURL:    srv.URL,
		MaxQPS:     5,
		APIVersion: "2017-10-17",
	})
	return srv, client
}
