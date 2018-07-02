package scaffold

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/lob/lob-go/nullable"
)

// Address scaffold wraps fields used to create an Address object in Lob API.
type Address struct {
	Name        nullable.String   `json:"name"`
	Line1       string            `json:"address_line1"`
	Line2       nullable.String   `json:"address_line2"`
	City        nullable.String   `json:"address_city"`
	Zip         nullable.String   `json:"address_zip"`
	State       nullable.String   `json:"address_state"`
	Country     nullable.String   `json:"address_country"`
	Company     nullable.String   `json:"company"`
	Email       nullable.String   `json:"email"`
	Phone       nullable.String   `json:"phone"`
	Description nullable.String   `json:"description"`
	Metadata    map[string]string `json:"metadata"`
}

// BankAccount scaffold wraps fields used to create a BankAccount object in Lob API.
type BankAccount struct {
	AccountNumber string            `json:"account_number"`
	AccountType   string            `json:"account_type"`
	RoutingNumber string            `json:"routing_number"`
	Signatory     string            `json:"signatory"`
	Description   nullable.String   `json:"description"`
	Metadata      map[string]string `json:"metadata"`
}

// Check scaffold wraps fields used to create a Check object in Lob API.
type Check struct {
	To             interface{}       `json:"to"`
	From           interface{}       `json:"from"`
	Description    nullable.String   `json:"description"`
	BankAccount    string            `json:"bank_account"`
	Logo           interface{}       `json:"logo"`
	Amount         float64           `json:"amount"`
	Message        nullable.String   `json:"message"`
	CheckBottom    interface{}       `json:"check_bottom"`
	Attachment     interface{}       `json:"attachment"`
	MergeVariables map[string]string `json:"merge_variables"`
	MailType       nullable.String   `json:"mail_type"`
	SendDate       nullable.String   `json:"send_date"`
	Metadata       map[string]string `json:"metadata"`
	IdempotencyKey string            `json:"-"`
}

// Letter scaffold wraps fields used to create a Letter object in Lob API.
type Letter struct {
	To               interface{}       `json:"to"`
	From             interface{}       `json:"from"`
	File             interface{}       `json:"file"`
	Color            bool              `json:"color"`
	DoubleSided      nullable.Bool     `json:"double_sided"`
	AddressPlacement nullable.String   `json:"address_placement"`
	ReturnEnvelope   nullable.Bool     `json:"return_envelope"`
	PerforatedPage   nullable.Int      `json:"perforated_page"`
	ExtraService     nullable.String   `json:"extra_service"`
	MergeVariables   map[string]string `json:"merge_variables"`
	MailType         nullable.String   `json:"mail_type"`
	SendDate         nullable.String   `json:"send_date"`
	Description      string            `json:"description"`
	Metadata         map[string]string `json:"metadata"`
	IdempotencyKey   string            `json:"-"`
}

// Postcard scaffold wraps fields used to create a Postcard object in Lob API.
type Postcard struct {
	To             interface{}       `json:"to"`
	From           interface{}       `json:"from"`
	Description    nullable.String   `json:"description"`
	Front          interface{}       `json:"front"`
	Back           interface{}       `json:"back"`
	MergeVariables map[string]string `json:"merge_variables"`
	Size           nullable.String   `json:"size"`
	MailType       nullable.String   `json:"mail_type"`
	SendDate       nullable.String   `json:"send_date"`
	Metadata       map[string]string `json:"metadata"`
	IdempotencyKey string            `json:"-"`
}

// MarshalAsFormValues attempts to coerce the fields of the given scaffold into strings,
// and places them in a map formatted for use in a multipart form request.
func MarshalAsFormValues(scaffold interface{}) (map[string]string, error) {
	v := reflect.ValueOf(scaffold).Elem()
	t := v.Type()
	fields := make(map[string]string)
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")

		val := v.FieldByName(field.Name).Interface()
		if val == nil {
			continue // Note: this branch is definitely hit by tests (if you remove it, you hit a seg fault), but doesn't show up in coverage.
		}

		switch fieldVal := val.(type) {
		case int, int64, float64, bool, string, []string:
			str := coerceToString(fieldVal)
			if len(str) > 0 {
				fields[tag] = str
			}
		case nullable.String:
			s, ok := fieldVal.Value()
			if ok {
				str := coerceToString(s)
				if len(str) > 0 {
					fields[tag] = str
				}
			}
		case nullable.Int:
			i, ok := fieldVal.Value()
			if ok {
				str := coerceToString(i)
				if len(str) > 0 {
					fields[tag] = str
				}
			}
		case nullable.Float:
			f, ok := fieldVal.Value()
			if ok {
				str := coerceToString(f)
				if len(str) > 0 {
					fields[tag] = str
				}
			}
		case nullable.Bool:
			b, ok := fieldVal.Value()
			if ok {
				str := coerceToString(b)
				if len(str) > 0 {
					fields[tag] = str
				}
			}
		case map[string]string:
			for key, value := range fieldVal {
				nestedTag := fmt.Sprintf("%s[%s]", tag, key)
				fields[nestedTag] = value
			}
		case *Address, *BankAccount, *Check, *Letter, *Postcard:
			nestedFieldMap, err := MarshalAsFormValues(fieldVal)
			if err != nil {
				return nil, err
			}
			for key, value := range nestedFieldMap {
				nestedTag := fmt.Sprintf("%s[%s]", tag, key)
				if len(value) > 0 {
					fields[nestedTag] = value
				}
			}
		case []byte:
			continue
		default:
			return nil, fmt.Errorf("Unable to parse field: %s", field.Name)
		}
	}
	return fields, nil
}

func coerceToString(val interface{}) string {
	var output string
	switch fieldVal := val.(type) {
	case int:
		output = strconv.Itoa(fieldVal)
	case bool:
		output = fmt.Sprintf("%v", fieldVal)
	case int64:
		output = strconv.FormatInt(fieldVal, 10)
	case float64:
		output = fmt.Sprintf("%.2f", fieldVal)
	case string:
		output = fieldVal
	case []string:
		output = strings.Join(fieldVal, " ")
	}
	return output
}
