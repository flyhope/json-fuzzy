package fuzzy

import (
	"encoding/json/jsontext"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

// FuzzyInt is a custom JSON unmarshaler for the int type that provides fuzzy decoding.
// It can decode an int from JSON numbers, strings, booleans, and null.
//
// Behavior:
//   - JSON numbers (e.g., 123, 45.67) are decoded as integers (floats are truncated).
//   - JSON strings (e.g., "123", " 45.67 ") are parsed as integers after trimming whitespace (floats are truncated).
//   - JSON booleans are decoded as 1 for true and 0 for false.
//   - JSON null is decoded as 0.
//   - Empty strings or strings with only whitespace are decoded as 0.
func FuzzyInt(dec *jsontext.Decoder, t *int) error {
	return fuzzyInteger(dec, t, strconv.ParseInt, 64)
}

// FuzzyInt8 is a custom JSON unmarshaler for the int8 type that provides fuzzy decoding.
// See FuzzyInt for more details on the decoding behavior.
func FuzzyInt8(dec *jsontext.Decoder, t *int8) error {
	return fuzzyInteger(dec, t, strconv.ParseInt, 8)
}

// FuzzyInt16 is a custom JSON unmarshaler for the int16 type that provides fuzzy decoding.
// See FuzzyInt for more details on the decoding behavior.
func FuzzyInt16(dec *jsontext.Decoder, t *int16) error {
	return fuzzyInteger(dec, t, strconv.ParseInt, 16)
}

// FuzzyInt32 is a custom JSON unmarshaler for the int32 type that provides fuzzy decoding.
// See FuzzyInt for more details on the decoding behavior.
func FuzzyInt32(dec *jsontext.Decoder, t *int32) error {
	return fuzzyInteger(dec, t, strconv.ParseInt, 32)
}

// FuzzyInt64 is a custom JSON unmarshaler for the int64 type that provides fuzzy decoding.
// See FuzzyInt for more details on the decoding behavior.
func FuzzyInt64(dec *jsontext.Decoder, t *int64) error {
	return fuzzyInteger(dec, t, strconv.ParseInt, 64)
}

// FuzzyUint is a custom JSON unmarshaler for the uint type that provides fuzzy decoding.
// See FuzzyInt for more details on the decoding behavior.
func FuzzyUint(dec *jsontext.Decoder, t *uint) error {
	return fuzzyInteger(dec, t, strconv.ParseUint, 64)
}

// FuzzyUint8 is a custom JSON unmarshaler for the uint8 type that provides fuzzy decoding.
// See FuzzyInt for more details on the decoding behavior.
func FuzzyUint8(dec *jsontext.Decoder, t *uint8) error {
	return fuzzyInteger(dec, t, strconv.ParseUint, 8)
}

// FuzzyUint16 is a custom JSON unmarshaler for the uint16 type that provides fuzzy decoding.
// See FuzzyInt for more details on the decoding behavior.
func FuzzyUint16(dec *jsontext.Decoder, t *uint16) error {
	return fuzzyInteger(dec, t, strconv.ParseUint, 16)
}

// FuzzyUint32 is a custom JSON unmarshaler for the uint32 type that provides fuzzy decoding.
// See FuzzyInt for more details on the decoding behavior.
func FuzzyUint32(dec *jsontext.Decoder, t *uint32) error {
	return fuzzyInteger(dec, t, strconv.ParseUint, 32)
}

// FuzzyUint64 is a custom JSON unmarshaler for the uint64 type that provides fuzzy decoding.
// See FuzzyInt for more details on the decoding behavior.
func FuzzyUint64(dec *jsontext.Decoder, t *uint64) error {
	return fuzzyInteger(dec, t, strconv.ParseUint, 64)
}

// fuzzyInteger is a generic helper function that implements the core fuzzy decoding logic
// for various integer types (int, int8, int32, int64, uint, etc.). It's designed to
// be called by type-specific fuzzy decoders, such as FuzzyInt.
//
// This function handles the decoding of an integer from different JSON types:
//   - JSON numbers (e.g., 123, 45.67) are decoded as integers (floats are truncated).
//   - JSON strings (e.g., "123", " 45.67 ") are parsed as integers after trimming whitespace (floats are truncated).
//   - JSON booleans are decoded as 1 for true and 0 for false.
//   - JSON null is decoded as 0.
//   - Empty strings or strings with only whitespace are decoded as 0.
//
// Parameters:
//   - dec: The jsontext.Decoder to read from.
//   - t: A pointer to the target integer variable to store the decoded value.
//   - method: The function used to parse the string representation of the number (e.g., strconv.ParseInt or strconv.ParseUint).
//   - bit: The bit size to be used by the parsing method.
func fuzzyInteger[T constraints.Integer](dec *jsontext.Decoder, t *T, method any, bit int) error {
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
				return fmt.Errorf("failed to unquote string for fuzzy integer: %w", err)
			}
			valueString = s
		} else {
			valueString = string(value)
		}

		// Trim whitespace for string values.
		valueString = strings.TrimSpace(valueString)

		// If there is a decimal point, truncate the fractional part.
		if pointIndex := strings.IndexByte(valueString, '.'); pointIndex >= 0 {
			valueString = valueString[:pointIndex]
		}

		if valueString == "" {
			*t = 0
			return nil
		}

		var result any
		var errAtoi error
		switch m := method.(type) {
		case func(string, int, int) (int64, error):
			result, errAtoi = m(valueString, 10, bit)
		case func(string, int, int) (uint64, error):
			result, errAtoi = m(valueString, 10, bit)
		default:
			return fmt.Errorf("unsupported parse function type")
		}

		if errAtoi != nil {
			return errAtoi
		}

		switch r := result.(type) {
		case int64:
			*t = T(r)
		case uint64:
			*t = T(r)
		default:
			return fmt.Errorf("unsupported parse result type")
		}

		return nil

	default:
		return fmt.Errorf("fuzzy integer must be a JSON string, number, boolean or null, got %v", kind)
	}
}