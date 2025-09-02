package fuzzy

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuzzyFloat64(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  float64
	}{
		{"number", "123.45", 123.45},
		{"string", `"67.89"`, 67.89},
		{"true", "true", 1.0},
		{"false", "false", 0.0},
		{"null", "null", 0.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var v struct {
				F float64 `json:",decoder=FuzzyFloat64"`
			}
			err := json.Unmarshal([]byte(`{"F":`+tc.input+"}"), &v)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, v.F)
		})
	}
}

func TestFuzzyFloat32(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  float32
	}{
		{"number", "123.45", 123.45},
		{"string", `"67.89"`, 67.89},
		{"true", "true", 1.0},
		{"false", "false", 0.0},
		{"null", "null", 0.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var v struct {
				F float32 `json:",decoder=FuzzyFloat32"`
			}
			err := json.Unmarshal([]byte(`{"F":`+tc.input+"}"), &v)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, v.F)
		})
	}
}
