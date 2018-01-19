package address

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/lob/lob-go/api"
	"github.com/lob/lob-go/nullable"
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

func TestCreate(t *testing.T) {
	name := "Larry Lobster"
	line1 := "185 Berry St"
	scaffold := &scaffold.Address{
		Name:  nullable.NewString(name),
		Line1: line1,
	}

	handler := func(rw http.ResponseWriter, req *http.Request) {
		log.Printf(req.URL.Path)
		if req.URL.Path != "/"+basePath {
			http.Error(rw, "internal server error", http.StatusInternalServerError)
		}
		io.WriteString(rw, fmt.Sprintf(`{"id":"addr_1", "name": "%s", "address_line1": "%s"}`, name, line1))
	}

	srv, client := test.NewSrvAndClient(handler)
	defer srv.Close()

	ac := NewClient(client)

	addy, err := ac.Create(context.Background(), scaffold)
	if err != nil {
		t.Errorf("Create: error parsing response as Address: %s", err)
	}
	if addy.ID == "" || addy.Name == nil || *addy.Name != name || addy.Line1 != line1 {
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

	ac := NewClient(client)

	addy, err := ac.Retrieve(context.Background(), id)
	if err != nil {
		t.Errorf("Retrieve: error parsing response as Address: %s", err)
	}
	if addy.ID != id {
		t.Errorf("Retrieve: address parsed incorrectly")
	}

	_, err = ac.Retrieve(context.Background(), "something fake")
	if err == nil {
		t.Errorf("Retrieve: expected error on invalid id")
	}
}

func TestDelete(t *testing.T) {
	id := "addr_1"
	handler := func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != fmt.Sprintf("/%s/%s", basePath, id) {
			http.Error(rw, "resource not found", http.StatusNotFound)
		}
		io.WriteString(rw, fmt.Sprintf(`{"id":"%s"}`, id))
	}

	srv, client := test.NewSrvAndClient(handler)
	defer srv.Close()

	ac := NewClient(client)

	err := ac.Delete(context.Background(), id)
	if err != nil {
		t.Errorf("Delete: error deleting Address: %s", err)
	}
	err = ac.Delete(context.Background(), "something fake")
	if err == nil {
		t.Errorf("Delete: expected error on invalid id")
	}
}

func TestList(t *testing.T) {
	limit, offset := 2, 1
	handler := func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.RawQuery != fmt.Sprintf("limit=%d&offset=%d", limit, offset) {
			http.Error(rw, "resource not found", http.StatusNotFound)
		}
		io.WriteString(rw, `{"data":[{"id": "adr_1"}, {"id": "adr_2"}], "count": 2}`)
	}

	srv, client := test.NewSrvAndClient(handler)
	defer srv.Close()

	ac := NewClient(client)

	addies, err := ac.List(context.Background(), limit, offset)
	if err != nil {
		t.Errorf("List: error parsing address: %s", err)
	}
	if len(addies) != 2 {
		t.Errorf("List: expected 2 address, received %d", len(addies))
	}

	_, err = ac.List(context.Background(), 5, 1000)
	if err == nil {
		t.Errorf("List: expected error on invalid id")
	}
}

func TestString(t *testing.T) {
	id := "addr_1"
	name := "Larry Lobster"
	line1 := "185 Berry St."
	addy := Address{
		ID:    id,
		Line1: line1,
		Name:  &name,
		Line2: nil,
	}

	s := addy.String()
	valid := strings.Contains(s, fmt.Sprintf(`"id":"%s"`, id))
	valid = valid && strings.Contains(s, fmt.Sprintf(`"name":"%s"`, name))
	valid = valid && strings.Contains(s, fmt.Sprintf(`"address_line1":"%s"`, line1))
	valid = valid && strings.Contains(s, `"address_line2":null`)
	if !valid {
		t.Errorf("String: unable to print expected address")
	}
}
