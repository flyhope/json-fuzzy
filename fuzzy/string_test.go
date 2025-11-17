package fuzzy

import (
	"encoding/json/v2"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fuzzyTestCaseString[T ~string] struct {
	name  string
	input T
	want  T
}

func TestFuzzyString(t *testing.T) {
	testCases := []fuzzyTestCaseString[string]{
		{"string", `"hello"`, "hello"},
		{"true", "true", "true"},
		{"false", "false", "false"},
		{"null", "null", ""},
		{`string with newline`, `"hello\nworld"`, "hello\nworld"},
		{`string with unicode escape`, `"hello\u0026world"`, "hello&world"},
	}

	unmarshaler := FuzzyUnmarshaler()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var v struct {
				S string
			}
			err := json.Unmarshal([]byte(`{"S":`+tc.input+"}"), &v, unmarshaler)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, v.S)
		})
	}
}

type testTypeString string

func TestFuzzyTypeString(t *testing.T) {
	testCases := []fuzzyTestCaseString[testTypeString]{
		{"string", `"hello"`, "hello"},
		{"true", "true", "true"},
		{"false", "false", "false"},
		{"null", "null", ""},
		{`string with newline`, `"hello\nworld"`, "hello\nworld"},
		{`string with unicode escape`, `"hello\u0026world"`, "hello&world"},
	}

	unmarshaler := FuzzyUnmarshaler()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var v struct {
				S testTypeString
			}
			err := json.Unmarshal([]byte(`{"S":`+tc.input+"}"), &v, unmarshaler)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, v.S)
		})
	}
}
