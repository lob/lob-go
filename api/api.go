// Package api provides convenience methods and types for interacting with
// a JSON API over http.
package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultMaxQPS = 10
	mediaType     = "application/json"
	userAgent     = "LobGo/0.0.0" // TODO(damian): figure out how we want to handle versioning
)

// The Client interface defines methods for interacting with an HTTP API.
type Client interface {
	NewRequest(context.Context, string, string, interface{}) (*http.Request, error)
	NewMultiformRequest(context.Context, string, string, *bytes.Buffer) (*http.Request, error)
	Do(*http.Request) (*http.Response, error)
	Get(context.Context, string) (*http.Response, error)
	Post(context.Context, string, interface{}) (*http.Response, error)
	Delete(context.Context, string) (*http.Response, error)
}

// JSONClient implements the Client interface and provides convenience methods for constructing
// and executing JSON http requests.
type JSONClient struct {
	opts   *Options
	client *http.Client

	throttle <-chan time.Time
}

// Options contains necessary configuration for a Client.
type Options struct {
	// BaseURL specifies the base target URL for a Client.
	// Paths used in creating requests will interpreted as relative to this URL.
	BaseURL string

	// APIKey is the Lob API key that will be used to authenticate each request.
	APIKey string

	// MaxQPS is the maximum number of queries that a Client will attempt
	// per second. This is client-side rate limiting, server-side rate
	// limiting may enforce different limits.
	MaxQPS int

	// APIVersion is the version of the Lob API to be used. If no value is
	// provided, it defaults to the most recent version of the API.
	APIVersion string
}

// NewRequest creates a new http.Request for the specified method and path, with the given body encoded as JSON.
func (c *JSONClient) NewRequest(ctx context.Context, method string, relPath string, body interface{}) (*http.Request, error) {
	// Ensure that the given relative path is a valid URL.
	rel, err := url.Parse(relPath)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	if body != nil {
		if err = json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	path := fmt.Sprintf("%s/%s", c.opts.BaseURL, rel.String())
	req, err := http.NewRequest(method, path, buf)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Add("Accept", mediaType)
	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("User-Agent", userAgent)
	req.SetBasicAuth(c.opts.APIKey, "")

	if c.opts.APIVersion != "" {
		req.Header.Add("Lob-Version", c.opts.APIVersion)
	}

	return req, nil
}

// NewMultiformRequest creates a new multipart form http.Request for the specified path (relative to the Client's BaseURL).
// The contentType needs to include the form boundary, and the payload buffer needs to be properly formatted.
func (c *JSONClient) NewMultiformRequest(ctx context.Context, relPath string, contentType string, formPayload *bytes.Buffer) (*http.Request, error) {
	// Ensure that the given relative path is a valid URL.
	rel, err := url.Parse(relPath)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/%s", c.opts.BaseURL, rel.String())
	req, err := http.NewRequest("POST", path, formPayload)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Add("Accept", mediaType)
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("User-Agent", userAgent)
	req.SetBasicAuth(c.opts.APIKey, "")

	if c.opts.APIVersion != "" {
		req.Header.Add("Lob-Version", c.opts.APIVersion)
	}

	return req, nil
}

// Do executes the given http.Request and returns an http.Response or error.
func (c *JSONClient) Do(req *http.Request) (*http.Response, error) {
	<-c.throttle

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err = CheckError(resp); err != nil {
		return nil, err
	}

	return resp, err
}

// Get provides a convenience wrapper for issuing http GET requests.
func (c *JSONClient) Get(ctx context.Context, url string) (*http.Response, error) {
	req, err := c.NewRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Post provides a convenience wrapper for issuing http POST requests with a JSON body.
func (c *JSONClient) Post(ctx context.Context, url string, body interface{}) (*http.Response, error) {
	req, err := c.NewRequest(ctx, "POST", url, body)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Delete provides a convenience wrapper for issuing http Delete requests.
func (c *JSONClient) Delete(ctx context.Context, url string) (*http.Response, error) {
	req, err := c.NewRequest(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// NewJSONClient creates a new JSONCLient object with the given options.
func NewJSONClient(opt *Options) *JSONClient {
	if opt.MaxQPS <= 0 {
		opt.MaxQPS = defaultMaxQPS
	}

	rate := time.Second / time.Duration(opt.MaxQPS)
	throttle := time.Tick(rate)

	return &JSONClient{
		opts:     opt,
		client:   &http.Client{},
		throttle: throttle,
	}
}

// Deserialize deserializes from.Body as JSON into the specified interface.
// Note: clients are responsible for closing from.Body; Deserialize does not do so.
func Deserialize(from *http.Response, to interface{}) error {
	return json.NewDecoder(from.Body).Decode(to)
}

// CheckError deserializes any errors present in the given http.Response.
func CheckError(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	errMsg := new(struct {
		Err struct {
			Message string `json:"message"`
		} `json:"error"`
	})

	err := json.NewDecoder(resp.Body).Decode(errMsg)
	if err != nil {
		return err
	}

	return errors.New(errMsg.Err.Message)
}
