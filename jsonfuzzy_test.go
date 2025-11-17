package jsonfuzzy

import (
	"testing"
)

func TestUnmarshal(t *testing.T) {
	type testStruct struct {
		StringField string  `json:"string_field"`
		IntField    int     `json:"int_field"`
		FloatField  float64 `json:"float_field"`
		BoolField   bool    `json:"bool_field"`
	}

	// JSON input with mismatched types
	jsonData := []byte(`{
		"string_field": 123,
		"int_field":    "456",
		"float_field":  "78.9",
		"bool_field":   "true"
	}`)

	var result testStruct
	err := Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if result.StringField != "123" {
		t.Errorf("expected StringField to be '123', got '%s'", result.StringField)
	}
	if result.IntField != 456 {
		t.Errorf("expected IntField to be 456, got %d", result.IntField)
	}
	if result.FloatField != 78.9 {
		t.Errorf("expected FloatField to be 78.9, got %f", result.FloatField)
	}
	if !result.BoolField {
		t.Errorf("expected BoolField to be true, got %v", result.BoolField)
	}
}
