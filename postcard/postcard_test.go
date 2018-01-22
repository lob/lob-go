package postcard

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/lob/lob-go/address"
	"github.com/lob/lob-go/api"
	"github.com/lob/lob-go/scaffold"
	test "github.com/lob/lob-go/testing"
)

func TestNewClient(t *testing.T) {
	client := new(struct {
		api.Client
	})

	c := NewClient(client)

	if c.client == nil {
		t.Errorf("TestNewClient: unable to set api.Client")
	}
}

func TestCreateInline(t *testing.T) {
	id := "psc_1"
	name := "Larry Lobster"
	to := address.Address{
		ID:    "addr_1",
		Name:  &name,
		Line1: "185 Berry St.",
	}

	scaffold := &scaffold.Postcard{
		To: to,
	}

	handler := func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/"+basePath {
			http.Error(rw, "internal server error", http.StatusInternalServerError)
		}
		if req.Header.Get("Content-Type") != "application/json" {
			http.Error(rw, "internal server error", http.StatusInternalServerError)
		}
		io.WriteString(rw, fmt.Sprintf(`{"id":"%s", "to": %s}`, id, to.String()))
	}

	srv, client := test.NewSrvAndClient(handler)
	defer srv.Close()

	ac := NewClient(client)

	psc, err := ac.Create(context.Background(), scaffold)
	if err != nil {
		t.Errorf("Create: error parsing response as Postcard: %s", err)
	}
	if psc.ID == "" || psc.To.ID != to.ID {
		t.Errorf("Create: address parsed incorrectly")
	}
}

func TestCreateFileUpload(t *testing.T) {
	id := "psc_1"
	front := []byte("some local file contents")
	back := []byte("some other local file contents")
	to := "addr_1"

	scaffold := &scaffold.Postcard{
		To:    to,
		Front: front,
		Back:  back,
	}

	handler := func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/"+basePath {
			http.Error(rw, "internal server error", http.StatusInternalServerError)
		}
		if strings.Contains(req.Header.Get("Content-Type"), "mulitpart") {
			http.Error(rw, "internal server error", http.StatusInternalServerError)
		}
		io.WriteString(rw, fmt.Sprintf(`{"id":"%s", "to": {"id": "%s"}}`, id, to))
	}

	srv, client := test.NewSrvAndClient(handler)
	defer srv.Close()

	ac := NewClient(client)

	psc, err := ac.Create(context.Background(), scaffold)
	if err != nil {
		t.Errorf("Create: error parsing response as Postcard: %s", err)
	}
	if psc.ID == "" || psc.To.ID != to {
		t.Errorf("Create: address parsed incorrectly")
	}
}

func TestRetrieve(t *testing.T) {
	id := "addr_1"
	handler := func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != fmt.Sprintf("/%s/%s", basePath, id) {
			http.Error(rw, "resource not found", http.StatusNotFound)
		}
		io.WriteString(rw, fmt.Sprintf(`{"id":"%s"}`, id))
	}

	srv, client := test.NewSrvAndClient(handler)
	defer srv.Close()

	pc := NewClient(client)

	psc, err := pc.Retrieve(context.Background(), id)
	if err != nil {
		t.Errorf("Retrieve: error parsing response as Postcard: %s", err)
	}
	if psc.ID != id {
		t.Errorf("Retrieve: Postcard parsed incorrectly")
	}

	_, err = pc.Retrieve(context.Background(), "something fake")
	if err == nil {
		t.Errorf("Retrieve: expected error on invalid id")
	}
}

func TestCancel(t *testing.T) {
	id := "addr_1"
	handler := func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != fmt.Sprintf("/%s/%s", basePath, id) {
			http.Error(rw, "resource not found", http.StatusNotFound)
		}
		io.WriteString(rw, fmt.Sprintf(`{"id":"%s"}`, id))
	}

	srv, client := test.NewSrvAndClient(handler)
	defer srv.Close()

	pc := NewClient(client)

	err := pc.Cancel(context.Background(), id)
	if err != nil {
		t.Errorf("Cancel: error canceling Postcard: %s", err)
	}
	err = pc.Cancel(context.Background(), "something fake")
	if err == nil {
		t.Errorf("Cancel: expected error on invalid id")
	}
}

func TestList(t *testing.T) {
	limit, offset := 2, 1
	handler := func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.RawQuery != fmt.Sprintf("limit=%d&offset=%d", limit, offset) {
			http.Error(rw, "resource not found", http.StatusNotFound)
		}
		io.WriteString(rw, `{"data":[{"id": "psc_1"}, {"id": "psc_2"}], "count": 2}`)
	}

	srv, client := test.NewSrvAndClient(handler)
	defer srv.Close()

	pc := NewClient(client)

	pscs, err := pc.List(context.Background(), limit, offset)
	if err != nil {
		t.Errorf("List: error parsing Postcards: %s", err)
	}
	if len(pscs) != 2 {
		t.Errorf("List: expected 2 Postcards, received %d", len(pscs))
	}

	_, err = pc.List(context.Background(), 5, -1000)
	if err == nil {
		t.Errorf("List: expected error on invalid offset")
	}
}

func TestString(t *testing.T) {
	id := "psc_1"
	name := "Larry Lobster"
	to := address.Address{
		ID:    "addr_1",
		Name:  &name,
		Line1: "185 Berry St.",
	}

	psc := Postcard{
		ID:   id,
		To:   to,
		From: nil,
	}

	s := psc.String()
	valid := strings.Contains(s, fmt.Sprintf(`"id":"%s"`, id))
	valid = valid && strings.Contains(s, to.String())
	valid = valid && strings.Contains(s, `"from":null`)
	if !valid {
		t.Errorf("String: unable to print expected address")
	}
}
