package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewJSONClient(&Options{"foo.com", "api_key", 0, ""})

	if c.opts == nil {
		t.Fatalf("NewClient: unable to set client Options")
	}

	if c.opts.MaxQPS != defaultMaxQPS {
		t.Errorf("NewClient: unable to set default MaxQPS")
	}

	if c.client == nil {
		t.Errorf("NewClient: unable to create http.Client")
	}
}

func TestNewRequest(t *testing.T) {
	c := NewJSONClient(&Options{"foo.com", "api_key", 10, "2017-10-17"})

	var tests = []struct {
		label     string
		method    string
		relPath   string
		body      interface{}
		shouldErr bool
	}{
		{"ValidRequest", "GET", "v1/foo", nil, false},
		{"ValidRequestWithPayload", "POST", "v1/foo", map[string]string{"foo": "bar"}, false},
		{"BadURL", "GET", "@wtf@http://@\\", nil, true},
		{"BadJSON", "GET", "v1/foo", make(chan int), true},
		{"BadMethod", "SOMETHING FAKE", "v1/foo", nil, true},
	}

	for _, test := range tests {
		_, err := c.NewRequest(context.Background(), test.method, test.relPath, test.body)
		if err == nil && test.shouldErr {
			t.Errorf("%s failed: error was expected but not received", test.label)
		}
		if err != nil && !test.shouldErr {
			t.Errorf("%s failed: unexpected error encountered: %s", test.label, err)
		}
	}
}

func TestNewMultiformRequest(t *testing.T) {
	c := NewJSONClient(&Options{"foo.com", "api_key", 10, "2017-10-17"})

	var tests = []struct {
		label     string
		relPath   string
		body      *bytes.Buffer
		shouldErr bool
	}{
		{"ValidRequest", "v1/foo", &bytes.Buffer{}, false},
		{"BadURL", "@wtf@http://@\\", &bytes.Buffer{}, true},
	}

	for _, test := range tests {
		_, err := c.NewMultiformRequest(context.Background(), test.relPath, "some fake content type", test.body)
		if err == nil && test.shouldErr {
			t.Errorf("%s failed: error was expected but not received", test.label)
		}
	}
}

func TestDo(t *testing.T) {
	mux := http.NewServeMux()
	srv := httptest.NewServer(mux)
	client := NewJSONClient(&Options{srv.URL, "api_key", 10, ""})
	defer srv.Close()

	mux.HandleFunc("/pass", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/fail", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	goodReq, _ := client.NewRequest(context.Background(), "GET", "pass", nil)
	badReq, _ := client.NewRequest(context.Background(), "GET", "fail", nil)

	tests := []struct {
		label     string
		req       *http.Request
		shouldErr bool
	}{
		{"ValidRequest", goodReq, false},
		{"BadResponse", badReq, true},
	}

	for _, test := range tests {
		_, err := client.Do(test.req)

		if test.shouldErr {
			if err == nil {
				t.Errorf("%s: expected error", test.label)
			}
		} else {
			if err != nil {
				t.Errorf("%s: unexpected error %#v: %s", test.label, err, err)
			}
		}
	}
}

func TestGet(t *testing.T) {
	mux := http.NewServeMux()
	srv := httptest.NewServer(mux)
	client := NewJSONClient(&Options{srv.URL, "api_key", 10, ""})
	defer srv.Close()

	mux.HandleFunc("/pass", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/fail", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	tests := []struct {
		label     string
		url       string
		shouldErr bool
	}{
		{"ValidRequest", "/pass", false},
		{"BadResponse", "/fail", true},
	}

	for _, test := range tests {
		_, err := client.Get(context.Background(), test.url)

		if test.shouldErr {
			if err == nil {
				t.Errorf("%s: expected error", test.label)
			}
		} else {
			if err != nil {
				t.Errorf("%s: unexpected error %#v: %s", test.label, err, err)
			}
		}
	}
}

func TestCheckError(t *testing.T) {
	msg := `{"error":{"message":"fail","status_code":500}}`

	tests := []struct {
		label     string
		resp      *http.Response
		shouldErr bool
		errMsg    string
	}{
		{"GoodResponse",
			&http.Response{
				Body:       ioutil.NopCloser(strings.NewReader("ok")),
				StatusCode: http.StatusOK,
			},
			false,
			"",
		},
		{"StatusCodeError",
			&http.Response{
				Body:       ioutil.NopCloser(strings.NewReader(msg)),
				StatusCode: http.StatusInternalServerError,
			},
			true,
			msg,
		},
		{"BadResponseBodyError",
			&http.Response{
				Body:       ioutil.NopCloser(strings.NewReader("fail")),
				StatusCode: http.StatusInternalServerError,
			},
			true,
			"fail",
		},
	}

	for _, test := range tests {
		err := CheckError(test.resp)
		if test.shouldErr {
			if err == nil {
				t.Errorf("%s: expected error", test.label)
			}
			if err.Error() == test.errMsg {
				t.Errorf("%s: bad error message, expected %s, actual %s", test.label, test.errMsg, err.Error())
			}
		} else {
			if err != nil {
				t.Errorf("%s: unexpected error %s", test.label, err)
			}
		}
	}
}

func TestDeserialize(t *testing.T) {
	type testType struct {
		OK bool
	}

	tests := []struct {
		label     string
		from      *http.Response
		to        interface{}
		shouldErr bool
	}{
		{"DeserializationSucceeds",
			&http.Response{
				Body: ioutil.NopCloser(strings.NewReader(`{"ok":true}`)),
			},
			&testType{},
			false,
		},
		{"TargetMismatch",
			&http.Response{
				Body: ioutil.NopCloser(strings.NewReader(`{"notBool":what?}`)),
			},
			&testType{},
			true,
		},
	}

	for _, test := range tests {
		err := Deserialize(test.from, test.to)
		if err == nil && test.shouldErr {
			t.Errorf("%s: expected error", test.label)
		}
		if !test.shouldErr {
			if err != nil {
				t.Errorf("%s: unexpected error %s", test.label, err)
			}
			if !test.to.(*testType).OK {
				t.Errorf("%s: json deserialization failed", test.label)
			}
		}
	}
}
