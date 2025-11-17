package fuzzy

import (
	"encoding/json/v2"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testFuzzyBoolStruct struct {
	Value bool `json:"value"`
}

type fuzzyBoolTestCase struct {
	name        string
	jsonData    string
	expected    bool
	expectError bool
}

func runFuzzyBoolTestCases(t *testing.T, tests []fuzzyBoolTestCase, opts ...json.Options) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result testFuzzyBoolStruct
			err := json.Unmarshal([]byte(tt.jsonData), &result, opts...)

			if tt.expectError {
				assert.Error(t, err, fmt.Sprintf("Expected an error for test case: %s", tt.name))
			} else {
				assert.NoError(t, err, fmt.Sprintf("Did not expect an error for test case: %s", tt.name))
				assert.Equal(t, tt.expected, result.Value, fmt.Sprintf("Mismatched expectation for test case: %s", tt.name))
			}
		})
	}
}

func TestFuzzyBool(t *testing.T) {
	tests := []fuzzyBoolTestCase{
		{
			name:     "Boolean true value",
			jsonData: `{"value": true}`,
			expected: true,
		},
		{
			name:     "Boolean false value",
			jsonData: `{"value": false}`,
			expected: false,
		},
		{
			name:     "Null value",
			jsonData: `{"value": null}`,
			expected: false,
		},
		{
			name:     "Number 0",
			jsonData: `{"value": 0}`,
			expected: false,
		},
		{
			name:     "Number 1",
			jsonData: `{"value": 1}`,
			expected: true,
		},
		{
			name:     "Number non-zero",
			jsonData: `{"value": 123}`,
			expected: true,
		},
		{
			name:     "String true",
			jsonData: `{"value": "true"}`,
			expected: true,
		},
		{
			name:     "String false",
			jsonData: `{"value": "false"}`,
			expected: false,
		},
		{
			name:     "String 1",
			jsonData: `{"value": "1"}`,
			expected: true,
		},
		{
			name:     "String 0",
			jsonData: `{"value": "0"}`,
			expected: false,
		},
		{
			name:     "Empty string",
			jsonData: `{"value": ""}`,
			expected: false,
		},
		{
			name:     "String with spaces",
			jsonData: `{"value": "  true  "}`,
			expected: true,
		},
		{
			name:        "Invalid string",
			jsonData:    `{"value": "abc"}`,
			expected:    false,
			expectError: false, // should be false
		},
	}

	unmarshaler := FuzzyUnmarshalerOrigin()
	runFuzzyBoolTestCases(t, tests, unmarshaler)
}
