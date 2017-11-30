package nullable

import (
	"encoding/json"
)

// Int represents a nullable integer.
type Int struct {
	value int
	valid bool
}

// NewInt creates a new Bool with the given value.
func NewInt(i int) Int {
	return Int{
		value: i,
		valid: true,
	}
}

// Set sets the value of the Int.
func (i *Int) Set(val int) {
	i.value = val
	i.valid = true
}

// Clear clears the current Int and marks it as invalid.
func (i *Int) Clear() {
	i.value = 0
	i.valid = false
}

// Value returns the current value of the Int, and whether it is valid.
func (i *Int) Value() (int, bool) {
	return i.value, i.valid
}

// MarshalJSON implments the Marshaler interface for Int.
func (i Int) MarshalJSON() ([]byte, error) {
	return marshalOrNull(i.Value())
}

// String represents a nullable string.
type String struct {
	value string
	valid bool
}

// NewString creates a new String with the given value
func NewString(s string) String {
	return String{
		value: s,
		valid: true,
	}
}

// Set sets the value of the String.
func (s *String) Set(val string) {
	s.value = val
	s.valid = true
}

// Clear clears the current String and marks it as invalid.
func (s *String) Clear() {
	s.value = ""
	s.valid = false
}

// Value returns the current value of the String, and whether it is valid.
func (s *String) Value() (string, bool) {
	return s.value, s.valid
}

// MarshalJSON implments the Marshaler interface for String.
func (s String) MarshalJSON() ([]byte, error) {
	return marshalOrNull(s.Value())
}

// Float represents a nullable float64.
type Float struct {
	value float64
	valid bool
}

// NewFloat creates a new Float with the given value
func NewFloat(f float64) Float {
	return Float{
		value: f,
		valid: true,
	}
}

// Set sets the value of the Float.
func (f *Float) Set(val float64) {
	f.value = val
	f.valid = true
}

// Clear clears the current Float and marks it as invalid.
func (f *Float) Clear() {
	f.value = 0
	f.valid = false
}

// Value returns the current value of the Float, and whether it is valid.
func (f *Float) Value() (float64, bool) {
	return f.value, f.valid
}

// MarshalJSON implments the Marshaler interface for Float.
func (f Float) MarshalJSON() ([]byte, error) {
	return marshalOrNull(f.Value())
}

// Bool represents a nullable bool.
type Bool struct {
	value bool
	valid bool
}

// NewBool creates a new Bool with the given value
func NewBool(b bool) Bool {
	return Bool{
		value: b,
		valid: true,
	}
}

// Set sets the value of the Bool.
func (b *Bool) Set(val bool) {
	b.value = val
	b.valid = true
}

// Clear clears the current Bool and marks it as invalid.
func (b *Bool) Clear() {
	b.value = false
	b.valid = false
}

// Value returns the current value of the Bool, and whether it is valid.
func (b *Bool) Value() (bool, bool) {
	return b.value, b.valid
}

// MarshalJSON implments the Marshaler interface for Bool.
func (b Bool) MarshalJSON() ([]byte, error) {
	return marshalOrNull(b.Value())
}

func marshalOrNull(val interface{}, ok bool) ([]byte, error) {
	if !ok {
		return []byte("null"), nil
	}
	return json.Marshal(val)
}
