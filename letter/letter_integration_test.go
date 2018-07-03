package letter_test

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/lob/lob-go/letter"
	"github.com/lob/lob-go/nullable"
	"github.com/lob/lob-go/scaffold"
	test "github.com/lob/lob-go/testing"
)

func TestBasicIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	ctx := context.Background()

	addy := &scaffold.Address{
		Name:  nullable.NewString("Larry Lobster"),
		Line1: "185 Berry St.",
		Line2: nullable.NewString("Suite 6100"),
		City:  nullable.NewString("San Francisco"),
		Zip:   nullable.NewString("94107"),
		State: nullable.NewString("CA"),
	}

	file, err := os.Open("../testing/resources/8_5x11_2.pdf")
	if err != nil {
		log.Printf("Error: %s", err)
	}

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	defer file.Close()

	scaffold := &scaffold.Letter{
		To:       addy,
		From:     addy,
		File:     contents,
		Metadata: map[string]string{"origin": "integration test"},
	}

	client := test.NewClient()
	c := letter.NewClient(client)

	ltr, err := c.Create(ctx, scaffold)
	if err != nil {
		t.Fatalf("Integration: error creating Letter: %s", err)
	}
	if ltr.ID == "" || ltr.To.ID == "" {
		t.Fatalf("Integration: Letter parsed incorrectly")
	}

	origID := ltr.ID
	ltr, err = c.Retrieve(ctx, ltr.ID)
	if err != nil {
		t.Fatalf("Integration: error parsing response as Letter: %s", err)
	}
	if ltr.ID != origID {
		t.Fatalf("Integration: Letter parsed incorrectly")
	}

	_, err = c.Retrieve(ctx, "something fake")
	if err == nil {
		t.Errorf("Integration: expected error on invalid id")
	}

	err = c.Cancel(ctx, ltr.ID)
	if err != nil {
		t.Errorf("Integration: error canceling Letter: %s", err)
	}

	limit, offset := 2, 1
	ltrs, err := c.List(ctx, limit, offset)
	if err != nil {
		t.Fatalf("List: error parsing Letters: %s", err)
	}
	if len(ltrs) != 2 {
		t.Errorf("List: expected 2 Letters, received %d", len(ltrs))
	}

	_, err = c.List(ctx, 5, -1000)
	if err == nil {
		t.Errorf("List: expected error on invalid offset")
	}
}
