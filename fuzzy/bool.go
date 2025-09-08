package fuzzy

import (
	"encoding/json/jsontext"
	"strconv"
	"strings"
)

// FuzzyBool is a custom JSON unmarshaler for the bool type that provides fuzzy decoding.
// It can decode a bool from JSON booleans, strings, numbers, and null.
func FuzzyBool(dec *jsontext.Decoder, t *bool) error {
	val, err := parseBool(dec)
	if err != nil {
		return err
	}
	*t = val
	return nil
}

func parseBool(dec *jsontext.Decoder) (bool, error) {
	kind := dec.PeekKind()
	switch kind {
	case 't': // true
		if err := dec.SkipValue(); err != nil {
			return false, err
		}
		return true, nil
	case 'f', 'n': // false, null
		if err := dec.SkipValue(); err != nil {
			return false, err
		}
		return false, nil
	case '0': // number
		result, err := parseUint64(kind, dec)
		if err != nil {
			return false, err
		}
		return result != 0, nil
	case '"': // string
		value, err := dec.ReadValue()
		if err != nil {
			return false, err
		}
		s, err := strconv.Unquote(string(value))
		if err != nil {
			return false, err
		}
		s = strings.TrimSpace(strings.ToLower(s))
		return s == "true" || s == "1", nil
	default:
		return false, dec.SkipValue()
	}
}
