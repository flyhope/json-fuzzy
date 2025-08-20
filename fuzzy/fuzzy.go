package fuzzy

import (
	"encoding/json/v2"
)

func FuzzyUnmarshaler() json.Options {
	return json.JoinOptions(
		// integer
		json.WithUnmarshalers(json.UnmarshalFromFunc(FuzzyInt)),
		// json.WithUnmarshalers(json.UnmarshalFromFunc(FuzzyInteger[int64])),
		// json.WithUnmarshalers(json.UnmarshalFromFunc(FuzzyInteger[int32])),
		// json.WithUnmarshalers(json.UnmarshalFromFunc(FuzzyInteger[int16])),
		// json.WithUnmarshalers(json.UnmarshalFromFunc(FuzzyInteger[int8])),
		// json.WithUnmarshalers(json.UnmarshalFromFunc(FuzzyInteger[uint])),
		// json.WithUnmarshalers(json.UnmarshalFromFunc(FuzzyInteger[uint64])),
		// json.WithUnmarshalers(json.UnmarshalFromFunc(FuzzyInteger[uint32])),
		// json.WithUnmarshalers(json.UnmarshalFromFunc(FuzzyInteger[uint16])),
		// json.WithUnmarshalers(json.UnmarshalFromFunc(FuzzyInteger[uint8])),
	)
}
