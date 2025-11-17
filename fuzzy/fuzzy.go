package fuzzy

import (
	"encoding/json/v2"
	"sync"
)

var FuzzyUnmarshaler = FuzzyUnmarshalerFull

var FuzzyUnmarshalerOrigin = sync.OnceValue(func() json.Options {
	return json.WithUnmarshalers(
		json.JoinUnmarshalers(FuzzyOriginUnmarshalers()...),
	)
})

var FuzzyUnmarshalerFull = sync.OnceValue(func() json.Options {
	unmarshalers := FuzzyOriginUnmarshalers()
	unmarshalers = append(unmarshalers, json.UnmarshalFromFunc(FuzzyAny))

	return json.WithUnmarshalers(
		json.JoinUnmarshalers(unmarshalers...),
	)
})

func FuzzyOriginUnmarshalers() []*json.Unmarshalers {
	return []*json.Unmarshalers{
		// string
		json.UnmarshalFromFunc(FuzzyString),
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
		// bool
		json.UnmarshalFromFunc(FuzzyBool),
	}
}
