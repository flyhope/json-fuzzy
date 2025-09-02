package fuzzy

import (
	"encoding/json/jsontext"
	"fmt"
	"strconv"
)

// FuzzyString is a custom JSON unmarshaler for the string type that provides fuzzy decoding.
// It can decode a string from JSON numbers, booleans, null, and strings.
//
// Behavior:
//   - JSON strings (e.g., "hello") are decoded as is.
//   - JSON numbers (e.g., 123, 45.67) are converted to their string representation.
//   - JSON booleans are decoded as "true" or "false".
//   - JSON null is decoded as an empty string "".
func FuzzyString[T ~string](dec *jsontext.Decoder, t *T) error {
	kind := dec.PeekKind()
	switch kind {
	case 'n': // null
		if err := dec.SkipValue(); err != nil {
			return err
		}
		*t = ""
		return nil
	case 't': // true
		if err := dec.SkipValue(); err != nil {
			return err
		}
		*t = "true"
		return nil
	case 'f': // false
		if err := dec.SkipValue(); err != nil {
			return err
		}
		*t = "false"
		return nil
	case '"': // string
		val, err := dec.ReadValue()
		if err != nil {
			return err
		}
		s, err := strconv.Unquote(string(val))
		if err != nil {
			return err
		}
		*t = T(s)
		return nil
	case '0': // number
		val, err := dec.ReadValue()
		if err != nil {
			return err
		}
		*t = T(val)
		return nil
	default:
		return fmt.Errorf("fuzzy string must be a JSON string, number, boolean or null, got %v", kind)
	}
}
