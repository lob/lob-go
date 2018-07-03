package address_test

import (
	"context"
	"testing"

	"github.com/lob/lob-go/address"
	"github.com/lob/lob-go/nullable"
	"github.com/lob/lob-go/scaffold"
	test "github.com/lob/lob-go/testing"
)

func TestBasicIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Create: Skipping integration test")
	}
	ctx := context.Background()

	scfld := &scaffold.Address{
		Name:  nullable.NewString("Larry Lobster"),
		Line1: "185 Berry St.",
		Line2: nullable.NewString("Suite 6100"),
		City:  nullable.NewString("San Francisco"),
		Zip:   nullable.NewString("94107"),
		State: nullable.NewString("CA"),
	}

	client := test.NewClient()
	c := address.NewClient(client)

	addr, err := c.Create(ctx, scfld)
	if err != nil {
		t.Fatalf("Error creating Address: %s", err)
	}
	if addr.ID == "" {
		t.Fatalf("Address parsed incorrectly")
	}

	_, err = c.Create(ctx, &scaffold.Address{Line1: "185 Berry St."})
	if err == nil {
		t.Errorf("Expected error creating invalid Address")
	}

	origID := addr.ID
	addr, err = c.Retrieve(ctx, addr.ID)
	if err != nil {
		t.Fatalf("Error parsing response as Address: %s", err)
	}
	if addr.ID != origID {
		t.Fatalf("Address parsed incorrectly")
	}

	_, err = c.Retrieve(ctx, "something fake")
	if err == nil {
		t.Errorf("Expected error on invalid id")
	}

	limit, offset := 2, 1
	addrs, err := c.List(ctx, limit, offset)
	if err != nil {
		t.Fatalf("Error parsing address: %s", err)
	}
	if len(addrs) != 2 {
		t.Errorf("Expected 2 address, received %d", len(addrs))
	}

	_, err = c.List(ctx, 5, -1000)
	if err == nil {
		t.Errorf("Expected error on invalid offset")
	}

	err = c.Delete(ctx, addr.ID)
	if err != nil {
		t.Errorf("Error deleting Address: %s", err)
	}
	err = c.Delete(ctx, "something fake")
	if err == nil {
		t.Errorf("Expected error on invalid id")
	}
}
