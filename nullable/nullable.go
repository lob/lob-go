package nullable

import (
	"fmt"
)

type Int struct {
	value int
	valid bool
}

func NewInt(i int) Int {
	return Int{
		value: i,
		valid: true,
	}
}

func (i *Int) Set(val int) {
	i.value = val
	i.valid = true
}

func (i *Int) Clear() {
	i.value = 0
	i.valid = false
}

func (i *Int) Value() (int, bool) {
	return i.value, i.valid
}

func (i Int) MarshalJSON() ([]byte, error) {
	val, ok := i.Value()
	if !ok {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%d"`, val)), nil
}

type String struct {
	value string
	valid bool
}

func NewString(s string) String {
	return String{
		value: s,
		valid: true,
	}
}

func (i *String) Set(val string) {
	i.value = val
	i.valid = true
}

func (i *String) Clear() {
	i.value = ""
	i.valid = false
}

func (i *String) Value() (string, bool) {
	return i.value, i.valid
}

func (i String) MarshalJSON() ([]byte, error) {
	val, ok := i.Value()
	if !ok {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, val)), nil
}

type Float struct {
	value float64
	valid bool
}

func NewFloat(f float64) Float {
	return Float{
		value: f,
		valid: true,
	}
}

func (i *Float) Set(val float64) {
	i.value = val
	i.valid = true
}

func (i *Float) Clear() {
	i.value = 0
	i.valid = false
}

func (i *Float) Value() (float64, bool) {
	return i.value, i.valid
}

func (i Float) MarshalJSON() ([]byte, error) {
	val, ok := i.Value()
	if !ok {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%.2f"`, val)), nil
}

type Bool struct {
	value bool
	valid bool
}

func NewBool(b bool) Bool {
	return Bool{
		value: b,
		valid: true,
	}
}

func (i *Bool) Set(val bool) {
	i.value = val
	i.valid = true
}

func (i *Bool) Clear() {
	i.value = false
	i.valid = false
}

func (i *Bool) Value() (bool, bool) {
	return i.value, i.valid
}

func (i Bool) MarshalJSON() ([]byte, error) {
	val, ok := i.Value()
	if !ok {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%v"`, val)), nil
}
