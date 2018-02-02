package check

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
	id := "chk_1"
	name := "Larry Lobster"
	to := address.Address{
		ID:    "addr_1",
		Name:  &name,
		Line1: "185 Berry St.",
	}

	scaffold := &scaffold.Check{
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

	chk, err := ac.Create(context.Background(), scaffold)
	if err != nil {
		t.Errorf("Create: error parsing response as Check: %s", err)
	}
	if chk.ID == "" || chk.To.ID != to.ID {
		t.Errorf("Create: address parsed incorrectly")
	}
}

func TestCreateFileUpload(t *testing.T) {
	id := "chk_1"
	logo := []byte("some local file contents")
	to := "addr_1"

	scaffold := &scaffold.Check{
		To:   to,
		Logo: logo,
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

	chk, err := ac.Create(context.Background(), scaffold)
	if err != nil {
		t.Errorf("Create: error parsing response as Check: %s", err)
	}
	if chk.ID == "" || chk.To.ID != to {
		t.Errorf("Create: Check parsed incorrectly")
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

	chk, err := pc.Retrieve(context.Background(), id)
	if err != nil {
		t.Errorf("Retrieve: error parsing response as Check: %s", err)
	}
	if chk.ID != id {
		t.Errorf("Retrieve: Check parsed incorrectly")
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
		t.Errorf("Cancel: error canceling Check: %s", err)
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
		io.WriteString(rw, `{"data":[{"id": "chk_1"}, {"id": "chk_2"}], "count": 2}`)
	}

	srv, client := test.NewSrvAndClient(handler)
	defer srv.Close()

	pc := NewClient(client)

	chks, err := pc.List(context.Background(), limit, offset)
	if err != nil {
		t.Errorf("List: error parsing Checks: %s", err)
	}
	if len(chks) != 2 {
		t.Errorf("List: expected 2 Checks, received %d", len(chks))
	}

	_, err = pc.List(context.Background(), 5, -1000)
	if err == nil {
		t.Errorf("List: expected error on invalid offset")
	}
}

func TestString(t *testing.T) {
	id := "chk_1"
	name := "Larry Lobster"
	to := address.Address{
		ID:    "addr_1",
		Name:  &name,
		Line1: "185 Berry St.",
	}

	chk := Check{
		ID: id,
		To: to,
	}

	s := chk.String()
	valid := strings.Contains(s, fmt.Sprintf(`"id":"%s"`, id))
	valid = valid && strings.Contains(s, to.String())
	if !valid {
		t.Errorf("String: unable to print expected Check")
	}
}
