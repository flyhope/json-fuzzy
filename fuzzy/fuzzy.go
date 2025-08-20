package fuzzy

import (
	"encoding/json/v2"
)

func FuzzyUnmarshaler() json.Options {
	return json.WithUnmarshalers(
		// FuzzyInt and its variants for different integer types
		json.JoinUnmarshalers(
			json.UnmarshalFromFunc(FuzzyUint),
			json.UnmarshalFromFunc(FuzzyUint64),
			json.UnmarshalFromFunc(FuzzyUint32),
			json.UnmarshalFromFunc(FuzzyUint16),
			json.UnmarshalFromFunc(FuzzyUint8),
			json.UnmarshalFromFunc(FuzzyInt),
			json.UnmarshalFromFunc(FuzzyInt64),
			json.UnmarshalFromFunc(FuzzyInt32),
			json.UnmarshalFromFunc(FuzzyInt16),
			json.UnmarshalFromFunc(FuzzyInt8),
		),
	)
}
