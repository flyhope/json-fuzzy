
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
			name:     "Valid float number",
			jsonData: `{"value": 789.99}`,
			expected: 789,
		},
		{
			name:     "Valid string float number",
			jsonData: `{"value": "123.45"}`,
			expected: 123,
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
			name:        "Invalid string value",
			jsonData:    `{"value": "abc"}`,
			expectError: true,
		},
		{
			name:        "Null value",
			jsonData:    `{"value": null}`,
			expectError: true,
		},
		{
			name:        "Boolean value",
			jsonData:    `{"value": true}`,
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
