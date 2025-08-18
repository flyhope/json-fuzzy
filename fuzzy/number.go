package fuzzy

import (
	"encoding/json/jsontext"
	"fmt"
	"strconv"
	"strings"
)

func FuzzyInt(dec *jsontext.Decoder, t *int) error {
	kind := dec.PeekKind()

	// Only JSON strings or numbers are allowed here.
	if kind != '0' && kind != '"' {
		return fmt.Errorf("fuzzy int must be a JSON string or number, got %v", kind)
	}
	value, err := dec.ReadValue()
	if err != nil {
		return err
	}
	if len(value) == 0 {
		return nil
	}

	// Convert the JSON value to a string for parsing.
	var valueString string
	switch kind {
	// string
	case '"':
		s, err := strconv.Unquote(string(value))
		if err != nil {
			return fmt.Errorf("failed to unquote string for fuzzy int: %w", err)
		}
		valueString = s
	// number or float
	case '0':
		valueString = string(value)
	}

	// If there is a decimal point, truncate the fractional part.
	if pointIndex := strings.IndexByte(valueString, '.'); pointIndex > 0 {
		valueString = string(valueString[:pointIndex])
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
}
