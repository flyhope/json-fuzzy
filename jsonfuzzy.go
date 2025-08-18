package jsonfuzzy

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
)

type JsonUnknow struct {
	UnknowMap map[string]any `json:",unknown"`
}

// Unmarshal unmarshals the JSON-encoded data in into out.
// 默认自带兼容选项：允许重复键名、允许无效UTF-8
func Unmarshal(in []byte, out any, opts ...json.Options) error {
	options := WithOptions(opts...)
	return json.Unmarshal(in, out, options)
}

// Marshal marshals v into JSON-encoded bytes.
func Marshal(in any, opts ...json.Options) ([]byte, error) {
	options := WithOptions(opts...)
	return json.Marshal(in, options)
}

// MarshalString marshals v into JSON-encoded string.
func MarshalString(in any, opts ...json.Options) (string, error) {
	options := WithOptions(opts...)
	data, err := json.Marshal(in, options)
	return string(data), err
}

// WithOptions 返回指定的选项（包含默认选项）。
func WithOptions(opts ...json.Options) json.Options {
	options := DefaultOptions()
	if len(opts) > 0 {
		for _, opt := range opts {
			options = json.JoinOptions(options, opt)
		}
	}
	return options
}

// DefaultOptions 默认自带JSON兼容选项：允许重复键名、允许无效UTF-8
func DefaultOptions() json.Options {
	fuzzy := FuzzyUnmarshaler()
	return json.JoinOptions(
		jsontext.AllowDuplicateNames(true),
		jsontext.AllowInvalidUTF8(true),
		json.WithUnmarshalers(fuzzy),
	)
}
