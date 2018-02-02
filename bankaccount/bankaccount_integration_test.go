package bankaccount_test

import (
	"context"
	"testing"

	"github.com/lob/lob-go/bankaccount"
	"github.com/lob/lob-go/scaffold"
	test "github.com/lob/lob-go/testing"
)

func TestBasicIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Create: Skipping integration test")
	}
	ctx := context.Background()

	scaffold := &scaffold.BankAccount{
		RoutingNumber: "322271627",
		AccountNumber: "123456789",
		AccountType:   "company",
		Signatory:     "Larry Lobster",
	}

	client := test.NewClient()
	c := bankaccount.NewClient(client)

	acct, err := c.Create(ctx, scaffold)
	if err != nil {
		t.Fatalf("Integration: error creating BankAccount: %s", err)
	}
	if acct.ID == "" || acct.Signatory == "" {
		t.Fatalf("Integration: BankAccount parsed incorrectly")
	}

	origID := acct.ID
	acct, err = c.Retrieve(ctx, acct.ID)
	if err != nil {
		t.Fatalf("Integration: error parsing response as BankAccount: %s", err)
	}
	if acct.ID != origID {
		t.Fatalf("Integration: BankAccount parsed incorrectly")
	}

	acct, err = c.Verify(ctx, acct.ID, []int{25, 63})
	if err != nil {
		t.Fatalf("Integration: error verifying BankAccount: %s", err)
	}

	_, err = c.Retrieve(ctx, "something fake")
	if err == nil {
		t.Errorf("Integration: expected error on invalid id")
	}

	limit, offset := 2, 1
	accts, err := c.List(ctx, limit, offset)
	if err != nil {
		t.Fatalf("List: error parsing BankAccounts: %s", err)
	}
	if len(accts) != 2 {
		t.Errorf("List: expected 2 BankAccounts, received %d", len(accts))
	}

	_, err = c.List(ctx, 5, -1000)
	if err == nil {
		t.Errorf("List: expected error on invalid offset")
	}

	err = c.Delete(ctx, acct.ID)
	if err != nil {
		t.Errorf("Integration: error canceling BankAccount: %s", err)
	}
	err = c.Delete(ctx, "something fake")
	if err == nil {
		t.Errorf("Integration: expected error on invalid id")
	}
}
