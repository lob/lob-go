package bankaccount

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/lob/lob-go/nullable"

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

func TestCreate(t *testing.T) {
	name := "Crustacean Credit Union"
	acctNum := "12345-67890"
	scaffold := &scaffold.BankAccount{
		Signatory:     "Larry Lobster",
		Description:   nullable.String{},
		AccountNumber: acctNum,
	}

	handler := func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/"+basePath {
			http.Error(rw, "internal server error", http.StatusInternalServerError)
		}
		io.WriteString(rw, fmt.Sprintf(`{"id":"acct_1", "bank_name": "%s", "description": null, "account_number":"%s"}`, name, acctNum))
	}

	srv, client := test.NewSrvAndClient(handler)
	defer srv.Close()

	c := NewClient(client)

	acct, err := c.Create(context.Background(), scaffold)
	if err != nil {
		t.Errorf("Create: error parsing response as BankAccount: %s", err)
	}
	if acct.ID == "" || acct.BankName != name || acct.Description != nil {
		t.Errorf("Create: account parsed incorrectly")
	}
}

func TestRetrieve(t *testing.T) {
	id := "bank_1"
	handler := func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != fmt.Sprintf("/%s/%s", basePath, id) {
			http.Error(rw, "resource not found", http.StatusNotFound)
		}
		io.WriteString(rw, fmt.Sprintf(`{"id":"%s"}`, id))
	}

	srv, client := test.NewSrvAndClient(handler)
	defer srv.Close()

	c := NewClient(client)

	acct, err := c.Retrieve(context.Background(), id)
	if err != nil {
		t.Errorf("Retrieve: error parsing response as BankAccount: %s", err)
	}
	if acct.ID != id {
		t.Errorf("Retrieve: account parsed incorrectly")
	}

	_, err = c.Retrieve(context.Background(), "something fake")
	if err == nil {
		t.Errorf("Retrieve: expected error on invalid id")
	}
}

func TestDelete(t *testing.T) {
	id := "bank_1"
	handler := func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != fmt.Sprintf("/%s/%s", basePath, id) {
			http.Error(rw, "resource not found", http.StatusNotFound)
		}
		io.WriteString(rw, fmt.Sprintf(`{"id":"%s"}`, id))
	}

	srv, client := test.NewSrvAndClient(handler)
	defer srv.Close()

	c := NewClient(client)

	err := c.Delete(context.Background(), id)
	if err != nil {
		t.Errorf("Delete: error deleting BankAccount: %s", err)
	}
	err = c.Delete(context.Background(), "something fake")
	if err == nil {
		t.Errorf("Delete: expected error on invalid id")
	}
}

func TestVerify(t *testing.T) {
	id := "bank_1"
	handler := func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != fmt.Sprintf("/%s/%s/verify", basePath, id) {
			http.Error(rw, "resource not found", http.StatusNotFound)
		}
		io.WriteString(rw, fmt.Sprintf(`{"id":"%s"}`, id))
	}

	srv, client := test.NewSrvAndClient(handler)
	defer srv.Close()

	c := NewClient(client)

	_, err := c.Verify(context.Background(), id, []int{1, 35})
	if err != nil {
		t.Errorf("Verify: error deleting BankAccount: %s", err)
	}
	_, err = c.Verify(context.Background(), "some fake id", []int{0, 0})
	if err == nil {
		t.Errorf("Verify: expected error on invalid id")
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

	c := NewClient(client)

	accts, err := c.List(context.Background(), limit, offset)
	if err != nil {
		t.Errorf("List: error parsing BankAccount: %s", err)
	}
	if len(accts) != 2 {
		t.Errorf("List: expected 2 BankAccounts, received %d", len(accts))
	}

	_, err = c.List(context.Background(), 5, 1000)
	if err == nil {
		t.Errorf("List: expected error on invalid id")
	}
}

func TestString(t *testing.T) {
	id := "acct_1"
	sig := "Larry Lobster"
	acct := BankAccount{
		ID:        id,
		Signatory: sig,
		Verified:  true,
	}

	s := acct.String()
	valid := strings.Contains(s, fmt.Sprintf(`"id":"%s"`, id))
	valid = valid && strings.Contains(s, fmt.Sprintf(`"signatory":"%s"`, sig))
	valid = valid && strings.Contains(s, fmt.Sprintf(`"verified":%v`, true))
	if !valid {
		t.Errorf("String: unable to print expected bank account")
	}
}
