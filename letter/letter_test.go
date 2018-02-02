package letter

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
	id := "ltr_1"
	name := "Larry Lobster"
	to := address.Address{
		ID:    "addr_1",
		Name:  &name,
		Line1: "185 Berry St.",
	}

	scaffold := &scaffold.Letter{
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

	c := NewClient(client)

	ltr, err := c.Create(context.Background(), scaffold)
	if err != nil {
		t.Errorf("Create: error parsing response as Letter: %s", err)
	}
	if ltr.ID == "" || ltr.To.ID != to.ID {
		t.Errorf("Create: address parsed incorrectly")
	}
}

func TestCreateFileUpload(t *testing.T) {
	id := "ltr_1"
	file := []byte("some local file contents")
	to := "addr_1"

	scaffold := &scaffold.Letter{
		To:   to,
		File: file,
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

	c := NewClient(client)

	ltr, err := c.Create(context.Background(), scaffold)
	if err != nil {
		t.Errorf("Create: error parsing response as Letter: %s", err)
	}
	if ltr.ID == "" || ltr.To.ID != to {
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

	c := NewClient(client)

	ltr, err := c.Retrieve(context.Background(), id)
	if err != nil {
		t.Errorf("Retrieve: error parsing response as Letter: %s", err)
	}
	if ltr.ID != id {
		t.Errorf("Retrieve: Letter parsed incorrectly")
	}

	_, err = c.Retrieve(context.Background(), "something fake")
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

	c := NewClient(client)

	err := c.Cancel(context.Background(), id)
	if err != nil {
		t.Errorf("Cancel: error canceling Letter: %s", err)
	}
	err = c.Cancel(context.Background(), "something fake")
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
		io.WriteString(rw, `{"data":[{"id": "ltr_1"}, {"id": "ltr_2"}], "count": 2}`)
	}

	srv, client := test.NewSrvAndClient(handler)
	defer srv.Close()

	c := NewClient(client)

	ltrs, err := c.List(context.Background(), limit, offset)
	if err != nil {
		t.Errorf("List: error parsing Letter: %s", err)
	}
	if len(ltrs) != 2 {
		t.Errorf("List: expected 2 Letters, received %d", len(ltrs))
	}

	_, err = c.List(context.Background(), 5, -1000)
	if err == nil {
		t.Errorf("List: expected error on invalid offset")
	}
}

func TestString(t *testing.T) {
	id := "ltr_1"
	name := "Larry Lobster"
	to := address.Address{
		ID:    "addr_1",
		Name:  &name,
		Line1: "185 Berry St.",
	}

	ltr := Letter{
		ID:   id,
		To:   to,
		From: to,
	}

	s := ltr.String()
	valid := strings.Contains(s, fmt.Sprintf(`"id":"%s"`, id))
	valid = valid && strings.Contains(s, fmt.Sprintf(`"to":%s`, to.String()))
	valid = valid && strings.Contains(s, fmt.Sprintf(`"from":%s`, to.String()))
	if !valid {
		t.Errorf("String: unable to print expected address")
	}
}
