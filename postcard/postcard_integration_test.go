package postcard_test

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/lob/lob-go/nullable"
	"github.com/lob/lob-go/postcard"
	"github.com/lob/lob-go/scaffold"
	test "github.com/lob/lob-go/testing"
)

func TestBasicIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Create: Skipping integration test")
	}
	ctx := context.Background()

	to := &scaffold.Address{
		Name:  nullable.NewString("Larry Lobster"),
		Line1: "185 Berry St.",
		Line2: nullable.NewString("Suite 6100"),
		City:  nullable.NewString("San Francisco"),
		Zip:   nullable.NewString("94107"),
		State: nullable.NewString("CA"),
	}

	file, err := os.Open("../testing/resources/6_25x11_25.pdf")
	if err != nil {
		log.Printf("Error: %s", err)
	}

	front, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	defer file.Close()

	scaffold := &scaffold.Postcard{
		To:             to,
		From:           to,
		Front:          front,
		Back:           "<html style='padding: 1in; font-size: 50;'>Back HTML for {{name}}</html>",
		MergeVariables: map[string]string{"name": "Larry Lobster"},
		Size:           nullable.NewString("6x11"),
		Metadata:       map[string]string{"origin": "integration test"},
	}

	client := test.NewClient()
	c := postcard.NewClient(client)

	psc, err := c.Create(ctx, scaffold)
	if err != nil {
		t.Fatalf("Integration: error creating Postcard: %s", err)
	}
	if psc.ID == "" || psc.To.ID == "" {
		t.Fatalf("Integration: Postcard parsed incorrectly")
	}

	origID := psc.ID
	psc, err = c.Retrieve(ctx, psc.ID)
	if err != nil {
		t.Fatalf("Integration: error parsing response as Postcard: %s", err)
	}
	if psc.ID != origID {
		t.Fatalf("Integration: Postcard parsed incorrectly")
	}

	_, err = c.Retrieve(ctx, "something fake")
	if err == nil {
		t.Errorf("Integration: expected error on invalid id")
	}

	err = c.Cancel(ctx, psc.ID)
	if err != nil {
		t.Errorf("Integration: error canceling Postcard: %s", err)
	}
	err = c.Cancel(ctx, "something fake")
	if err == nil {
		t.Errorf("Integration: expected error on invalid id")
	}

	limit, offset := 2, 1
	pscs, err := c.List(ctx, limit, offset)
	if err != nil {
		t.Fatalf("List: error parsing Postcards: %s", err)
	}
	if len(pscs) != 2 {
		t.Errorf("List: expected 2 Postcards, received %d", len(pscs))
	}

	_, err = c.List(ctx, 5, -1000)
	if err == nil {
		t.Errorf("List: expected error on invalid offset")
	}
}
