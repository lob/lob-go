package check_test

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/lob/lob-go/bankaccount"
	"github.com/lob/lob-go/check"
	"github.com/lob/lob-go/nullable"
	"github.com/lob/lob-go/scaffold"
	test "github.com/lob/lob-go/testing"
)

func createBankAccount(ctx context.Context) (*bankaccount.BankAccount, error) {
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
		return nil, err
	}

	return c.Verify(ctx, acct.ID, []int{25, 63})
}

func TestBasicIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	ctx := context.Background()

	acct, err := createBankAccount(ctx)
	if err != nil {
		t.Fatalf("Integration: error creating bank account for checks test")
	}

	client := test.NewClient()
	c := check.NewClient(client)

	addy := &scaffold.Address{
		Name:  nullable.NewString("Larry Lobster"),
		Line1: "185 Berry St.",
		Line2: nullable.NewString("Suite 6100"),
		City:  nullable.NewString("San Francisco"),
		Zip:   nullable.NewString("94107"),
		State: nullable.NewString("CA"),
	}

	file, err := os.Open("../testing/resources/lob_check_test_logo.png")
	if err != nil {
		log.Printf("Error: %s", err)
	}

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	defer file.Close()

	scaffold := &scaffold.Check{
		To:          addy,
		From:        addy,
		Logo:        contents,
		BankAccount: acct.ID,
		Amount:      1.00,
	}

	chk, err := c.Create(ctx, scaffold)
	if err != nil {
		t.Fatalf("Integration: error creating Check: %s", err)
	}
	if chk.ID == "" || chk.To.ID == "" {
		t.Fatalf("Integration: Check parsed incorrectly")
	}

	origID := chk.ID
	chk, err = c.Retrieve(ctx, chk.ID)
	if err != nil {
		t.Fatalf("Integration: error parsing response as Check: %s", err)
	}
	if chk.ID != origID {
		t.Fatalf("Integration: Check parsed incorrectly")
	}

	_, err = c.Retrieve(ctx, "something fake")
	if err == nil {
		t.Errorf("Integration: expected error on invalid id")
	}

	limit, offset := 2, 1
	chks, err := c.List(ctx, limit, offset)
	if err != nil {
		t.Fatalf("List: error parsing Checks: %s", err)
	}
	if len(chks) != 2 {
		t.Errorf("List: expected 2 Checks, received %d", len(chks))
	}

	_, err = c.List(ctx, 5, -1000)
	if err == nil {
		t.Errorf("List: expected error on invalid offset")
	}

	err = c.Cancel(ctx, chk.ID)
	if err != nil {
		t.Errorf("Integration: error canceling Check: %s", err)
	}
	err = c.Cancel(ctx, "something fake")
	if err == nil {
		t.Errorf("Integration: expected error on invalid id")
	}
}
