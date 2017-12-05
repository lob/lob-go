package scaffold

import (
	"testing"

	"github.com/lob/lob-go/nullable"
)

func TestStringCoercion(t *testing.T) {
	tests := []struct {
		label    string
		val      interface{}
		expected string
	}{
		{"int", 1, "1"},
		{"int64", int64(1), "1"},
		{"float64", float64(1.0), "1.00"},
		{"bool", true, "true"},
		{"string", "foo", "foo"},
		{"[]string", []string{"1", "2"}, "1 2"},
	}

	for _, test := range tests {
		s := coerceToString(test.val)
		if s != test.expected {
			t.Errorf("%s: expected string %s, actual: %s", test.label, test.expected, s)
		}
	}
}

func TestUnsupportedMarshalling(t *testing.T) {
	intVal := 1
	nestedPsc := &Postcard{
		To: &intVal,
	}
	psc := &Postcard{
		To: nestedPsc,
	}
	_, err := MarshalAsFormValues(psc)
	if err == nil {
		t.Errorf("Unsupported type: expected error coercing *int to string")
	}

}

func TestAddressMarshalling(t *testing.T) {
	expected := map[string]string{
		"name":          "Larry Lobster",
		"address_line1": "185 Berry St.",
		"address_line2": "#6100",
		"metadata[foo]": "bar",
	}
	addy := &Address{
		Name:     nullable.NewString("Larry Lobster"),
		Line1:    "185 Berry St.",
		Line2:    nullable.NewString("#6100"),
		Metadata: map[string]string{"foo": "bar"},
	}

	formVals, err := MarshalAsFormValues(addy)
	if err != nil {
		t.Errorf("Address: error marshalling scaffold to form values: %s", err)
	}
	for key, val := range formVals {
		if val != expected[key] {
			t.Errorf("Address: marshalling error for key %s. Expected %s, actual %s", key, expected[key], val)
		}
	}
	for key, val := range expected {
		if val != formVals[key] {
			t.Errorf("Address: marshalling error. Missing expected key %s and value %s", key, val)
		}
	}
}

func TestPostcardMarshalling(t *testing.T) {
	expected := map[string]string{
		"to":                  "addr_12345",
		"from[name]":          "Larry Lobster",
		"from[address_line1]": "185 Berry St.",
		"from[address_line2]": "#6100",
		"back":                "1.00",
	}
	addy := &Address{
		Name:  nullable.NewString("Larry Lobster"),
		Line1: "185 Berry St.",
		Line2: nullable.NewString("#6100"),
	}
	psc := &Postcard{
		To:    "addr_12345",
		From:  addy,
		Front: []byte("some local file"),
		Back:  nullable.NewFloat(1.0), // Note: this isn't a valid value for Back, we're just abusing the interface{} for testing purposes.
	}

	formVals, err := MarshalAsFormValues(psc)
	if err != nil {
		t.Errorf("Postcard: error marshalling scaffold to form values: %s", err)
	}
	for key, val := range formVals {
		if val != expected[key] {
			t.Errorf("Postcard: marshalling error for key %s. Expected %s, actual %s", key, expected[key], val)
		}
	}
	for key, val := range expected {
		if val != formVals[key] {
			t.Errorf("Postcard: marshalling error. Missing expected key %s and value %s", key, val)
		}
	}
}

func TestLetterMarshalling(t *testing.T) {
	expected := map[string]string{
		"to":              "addr_12345",
		"file":            "<html>an html string</html>",
		"color":           "true",
		"double_sided":    "true",
		"perforated_page": "3",
		"mail_type":       "standard",
	}
	ltr := &Letter{
		To:             "addr_12345",
		From:           nullable.Float{}, // Note: Same story as Back above, this is just meant to test nullable.Float marshalling.
		File:           "<html>an html string</html>",
		Color:          true,
		DoubleSided:    nullable.NewBool(true),
		ReturnEnvelope: nullable.Bool{},
		PerforatedPage: nullable.NewInt(3),
		MailType:       nullable.NewString("standard"),
		SendDate:       nullable.String{},
	}

	formVals, err := MarshalAsFormValues(ltr)
	if err != nil {
		t.Errorf("Letter: error marshalling scaffold to form values: %s", err)
	}
	for key, val := range formVals {
		if val != expected[key] {
			t.Errorf("Letter: marshalling error for key %s. Expected %s, actual %s", key, expected[key], val)
		}
	}
	for key, val := range expected {
		if val != formVals[key] {
			t.Errorf("Letter: marshalling error. Missing expected key %s and value %s", key, val)
		}
	}
}
