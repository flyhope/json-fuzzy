package fuzzy

import (
	"encoding/json/v2"
)

func FuzzyUnmarshaler() *json.Unmarshalers {
	// int
	intUnmarshaler := json.UnmarshalFromFunc(FuzzyInt)

	return intUnmarshaler
}
