package fuzzy

import (
	"encoding/json/jsontext"
	"errors"
	"fmt"
)

// FuzzyString is a custom JSON unmarshaler for the string type that provides fuzzy decoding.
// It can decode a string from JSON numbers, booleans, null, and strings.
//
// Behavior:
//   - JSON strings (e.g., "hello") are decoded as is.
//   - JSON numbers (e.g., 123, 45.67) are converted to their string representation.
//   - JSON booleans are decoded as "true" or "false".
//   - JSON null is decoded as an empty string "".
func FuzzyString(dec *jsontext.Decoder, t *string) error {
	result, err := showStringValue(dec)
	if err != nil {
		return err
	}
	*t = result
	return nil
}

func showStringValue(dec *jsontext.Decoder) (string, error) {
	kind := dec.PeekKind()
	switch kind {
	case 'n': // null
		err := dec.SkipValue()
		return "", err
	case 't': // true
		err := dec.SkipValue()
		return "true", err
	case 'f': // false
		err := dec.SkipValue()
		return "false", err
	case '"': // string
		return "", errors.ErrUnsupported
	case '0': // number
		val, err := dec.ReadValue()
		return string(val), err
	default:
		return "", fmt.Errorf("fuzzy string must be a JSON string, number, boolean or null, got %v", kind)
	}
}
