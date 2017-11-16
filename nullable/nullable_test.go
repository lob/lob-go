package nullable

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	exp := "foo"
	var zero string
	s := NewString(exp)

	val, ok := s.Value()
	if !ok {
		t.Errorf("TestString: expected non-null value")
	}
	if val != exp {
		t.Errorf("TestString: expected: %s, actual: %s", exp, val)
	}

	json, _ := s.MarshalJSON()
	if string(json) != fmt.Sprintf(`"%s"`, exp) {
		t.Errorf("TestString: error marshaling JSON")
	}

	s.Clear()
	val, ok = s.Value()
	if ok {
		t.Errorf("TestString: expected null value")
	}
	if val != zero {
		t.Errorf("TestString: expected empty string")
	}

	json, _ = s.MarshalJSON()
	if string(json) != "null" {
		t.Errorf("TestString: error marshaling JSON")

	}

	s.Set(exp)
	val, ok = s.Value()
	if !ok {
		t.Errorf("TestString: expected non-null value")
	}
	if val != exp {
		t.Errorf("TestString: expected: %s, actual: %s", exp, val)
	}
}

func TestInt(t *testing.T) {
	var zero int
	exp := 3
	s := NewInt(exp)

	val, ok := s.Value()
	if !ok {
		t.Errorf("TestInt: expected non-null value")
	}
	if val != exp {
		t.Errorf("TestInt: expected: %d, actual: %d", exp, val)
	}

	json, _ := s.MarshalJSON()
	if string(json) != fmt.Sprintf("%d", exp) {
		t.Errorf("TestInt: error marshaling JSON")
	}

	s.Clear()
	val, ok = s.Value()
	if ok {
		t.Errorf("TestInt: expected null value")
	}
	if val != zero {
		t.Errorf("TestInt: expected %d", zero)
	}

	json, _ = s.MarshalJSON()
	if string(json) != "null" {
		t.Errorf("TestInt: error marshaling JSON")
	}

	s.Set(exp)
	val, ok = s.Value()
	if !ok {
		t.Errorf("TestInt: expected non-null value")
	}
	if val != exp {
		t.Errorf("TestInt: expected: %d, actual: %d", exp, val)
	}
}

func TestFloat(t *testing.T) {
	var zero float64
	exp := 1.35
	s := NewFloat(exp)

	val, ok := s.Value()
	if !ok {
		t.Errorf("TestFloat: expected non-null value")
	}
	if val != exp {
		t.Errorf("TestFloat: expected: %f, actual: %f", exp, val)
	}

	json, _ := s.MarshalJSON()
	if string(json) != fmt.Sprintf("%.2f", exp) {
		t.Errorf("TestFloat: error marshaling JSON")
	}

	s.Clear()
	val, ok = s.Value()
	if ok {
		t.Errorf("TestFloat: expected null value")
	}
	if val != zero {
		t.Errorf("TestFloat: expected expected %f", zero)
	}

	json, _ = s.MarshalJSON()
	if string(json) != "null" {
		t.Errorf("TestFloat: error marshaling JSON")
	}

	s.Set(exp)
	val, ok = s.Value()
	if !ok {
		t.Errorf("TestFloat: expected non-null value")
	}
	if val != exp {
		t.Errorf("TestFloat: expected: %f, actual: %f", exp, val)
	}
}

func TestBool(t *testing.T) {
	var zero bool
	exp := true
	s := NewBool(exp)

	val, ok := s.Value()
	if !ok {
		t.Errorf("TestBool: expected non-null value")
	}
	if val != exp {
		t.Errorf("TestBool: expected: %v, actual: %v", exp, val)
	}

	json, _ := s.MarshalJSON()
	if string(json) != fmt.Sprintf("%v", exp) {
		t.Errorf("TestBool: error marshaling JSON")

	}

	s.Clear()
	val, ok = s.Value()
	if ok {
		t.Errorf("TestBool: expected null value")
	}
	if val != zero {
		t.Errorf("TestBool: expected %v", zero)
	}

	json, _ = s.MarshalJSON()
	if string(json) != "null" {
		t.Errorf("TestBool: error marshaling JSON")
	}

	s.Set(exp)
	val, ok = s.Value()
	if !ok {
		t.Errorf("TestBool: expected non-null value")
	}
	if val != exp {
		t.Errorf("TestBool: expected: %v, actual: %v", exp, val)
	}
}
