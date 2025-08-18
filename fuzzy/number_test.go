package fuzzy

import (
	"encoding/json/v2"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testFuzzyIntStruct struct {
	Value int `json:"value"`
}

type fuzzyIntTestCase struct {
	name        string
	jsonData    string
	expected    int
	expectError bool
}

func TestFuzzyInt(t *testing.T) {
	unmarshaler := json.WithUnmarshalers(FuzzyUnmarshaler())

	tests := []fuzzyIntTestCase{
		// --- Valid Cases ---
		{
			name:     "Valid integer number",
			jsonData: `{"value": 123}`,
			expected: 123,
		},
		{
			name:     "Valid string number",
			jsonData: `{"value": "456"}`,
			expected: 456,
		},
		{
			name:     "String with whitespace",
			jsonData: `{"value": " 123 "}`,
			expected: 123,
		},
		{
			name:     "Valid float number, truncated",
			jsonData: `{"value": 789.99}`,
			expected: 789,
		},
		{
			name:     "Valid string float number, truncated",
			jsonData: `{"value": "123.45"}`,
			expected: 123,
		},
		{
			name:     "String float with whitespace, truncated",
			jsonData: `{"value": " 456.78 "}`,
			expected: 456,
		},
		{
			name:     "Zero value",
			jsonData: `{"value": 0}`,
			expected: 0,
		},
		{
			name:     "Empty string value",
			jsonData: `{"value": ""}`,
			expected: 0,
		},
		{
			name:     "String with only whitespace",
			jsonData: `{"value": "   "}`,
			expected: 0,
		},
		{
			name:     "Null value",
			jsonData: `{"value": null}`,
			expected: 0, // Was expectError: true
		},
		{
			name:     "Boolean true value",
			jsonData: `{"value": true}`,
			expected: 1, // Was expectError: true
		},
		{
			name:     "Boolean false value",
			jsonData: `{"value": false}`,
			expected: 0,
		},

		// --- Error Cases ---
		{
			name:        "Invalid non-numeric string",
			jsonData:    `{"value": "abc"}`,
			expectError: true,
		},
		{
			name:        "Object value",
			jsonData:    `{"value": {}}`,
			expectError: true,
		},
		{
			name:        "Array value",
			jsonData:    `{"value": []}`,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result testFuzzyIntStruct
			err := json.Unmarshal([]byte(tt.jsonData), &result, unmarshaler)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result.Value)
			}
		})
	}
}