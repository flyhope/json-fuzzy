package fuzzy

import (
	"encoding/json/v2"
	"sync"
)

var FuzzyUnmarshaler = sync.OnceValue(func() json.Options {
	return json.WithUnmarshalers(
		// FuzzyInt and its variants for different integer types
		json.JoinUnmarshalers(
			// string
			json.UnmarshalFromFunc(FuzzyString[string]),
			// int
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
			// float
			json.UnmarshalFromFunc(FuzzyFloat32),
			json.UnmarshalFromFunc(FuzzyFloat64),
			// any
		),
	)
})
