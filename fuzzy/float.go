package fuzzy

import (
	"encoding/json/jsontext"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

// FuzzyFloat32 is a custom JSON unmarshaler for the float32 type that provides fuzzy decoding.
// See fuzzyFloat for more details on the decoding behavior.
func FuzzyFloat32(dec *jsontext.Decoder, t *float32) error {
	return fuzzyFloat(dec, t, 32)
}

// FuzzyFloat64 is a custom JSON unmarshaler for the float64 type that provides fuzzy decoding.
// See fuzzyFloat for more details on the decoding behavior.
func FuzzyFloat64(dec *jsontext.Decoder, t *float64) error {
	return fuzzyFloat(dec, t, 64)
}

// fuzzyFloat is a generic helper function that implements the core fuzzy decoding logic
// for various float types (float32, float64). It's designed to
// be called by type-specific fuzzy decoders, such as FuzzyFloat64.
//
// This function handles the decoding of a float from different JSON types:
//   - JSON numbers (e.g., 123, 45.67) are decoded as floats.
//   - JSON strings (e.g., "123", "45.67") are parsed as floats after trimming whitespace.
//   - JSON booleans are decoded as 1.0 for true and 0.0 for false.
//   - JSON null is decoded as 0.0.
//   - Empty strings or strings with only whitespace are decoded as 0.0.
//
// Parameters:
//   - dec: The jsontext.Decoder to read from.
//   - t: A pointer to the target float variable to store the decoded value.
//   - bit: The bit size to be used by the parsing method (32 or 64).
func fuzzyFloat[T constraints.Float](dec *jsontext.Decoder, t *T, bit int) error {
	kind := dec.PeekKind()

	switch kind {
	case 'n': // null
		if err := dec.SkipValue(); err != nil {
			return err
		}
		*t = 0
		return nil
	case 't': // true
		if err := dec.SkipValue(); err != nil {
			return err
		}
		*t = 1
		return nil
	case 'f': // false
		if err := dec.SkipValue(); err != nil {
			return err
		}
		*t = 0
		return nil
	case '0', '"': // number or string
		value, err := dec.ReadValue()
		if err != nil {
			return err
		}
		if len(value) == 0 {
			*t = 0
			return nil
		}

		var valueString string
		if kind == '"' {
			s, err := strconv.Unquote(string(value))
			if err != nil {
				return fmt.Errorf("failed to unquote string for fuzzy float: %w", err)
			}
			valueString = s
		} else {
			valueString = string(value)
		}

		// Trim whitespace for string values.
		valueString = strings.TrimSpace(valueString)

		if valueString == "" {
			*t = 0
			return nil
		}

		f, err := strconv.ParseFloat(valueString, bit)
		if err != nil {
			return err
		}
		*t = T(f)
		return nil

	default:
		return fmt.Errorf("fuzzy float must be a JSON string, number, boolean or null, got %v", kind)
	}
}

// func showFloatByPeek[T constraints.Float](dec *jsontext.Decoder) (T, jsontext.Kind, bool, error) {
// 	// @todo
// }
