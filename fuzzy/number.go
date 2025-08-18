package fuzzy

import (
	"encoding/json/jsontext"
	"fmt"
	"strconv"
	"strings"
)

func FuzzyInt(dec *jsontext.Decoder, t *int) error {
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
				return fmt.Errorf("failed to unquote string for fuzzy int: %w", err)
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

		if result, errAtoi := strconv.Atoi(valueString); errAtoi != nil {
			return errAtoi
		} else {
			*t = result
		}

		return nil

	default:
		return fmt.Errorf("fuzzy int must be a JSON string, number, boolean or null, got %v", kind)
	}
}