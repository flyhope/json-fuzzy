# json-fuzzy

[中文文档](README_zh.md)

[![Go Reference](https://pkg.go.dev/badge/github.com/flyhope/json-fuzzy.svg)](https://pkg.go.dev/github.com/flyhope/json-fuzzy)

A Go library for fuzzy JSON unmarshaling that handles type mismatches gracefully. This package extends the standard `encoding/json/v2` functionality with lenient parsing options and custom unmarshalers to automatically convert between different JSON types.

## Features

- **Fuzzy Type Conversion**: Automatically converts between JSON types during unmarshaling:
  - Numbers, booleans, and null values to strings
  - Strings, booleans, and null values to integers and other numeric types
  - Various string representations to boolean values
  - Strings and numbers to float values
- **Lenient JSON Parsing**: Includes default options for more permissive JSON handling:
  - `AllowDuplicateNames`: Allows duplicate field names in JSON objects
  - `AllowInvalidUTF8`: Handles invalid UTF-8 sequences in strings
- **Seamless Integration**: Works as a drop-in replacement for standard JSON operations with familiar API
- **Extensible Design**: Custom unmarshalers can be easily added or modified

## Fuzzy Unmarshaler Modes

The library provides two configurable unmarshaler modes:

- **FuzzyUnmarshalerOrigin**: Handles basic type conversions (strings, numbers, booleans) with minimal overhead. Ideal for performance-critical applications where `any` type isn't used.
  
- **FuzzyUnmarshalerFull** (default): Extends basic mode with `any` type support through `FuzzyAny` unmarshaler. Provides maximum flexibility for handling heterogeneous data at a small performance cost (~10-15% overhead in benchmarks).

### Changing Default Behavior

You can change the global default behavior by modifying the `fuzzy.FuzzyUnmarshaler` variable **at program initialization** (before any JSON operations):

```go
package main

import (
	"github.com/flyhope/json-fuzzy/fuzzy"
)

func init() {
	// Switch to origin mode for better performance globally
	fuzzy.FuzzyUnmarshaler = fuzzy.FuzzyUnmarshalerOrigin
}

// Now all Unmarshal calls will use the origin mode by default
func main() {
	var result struct{ Value int }
	err := jsonfuzzy.Unmarshal([]byte(`{"Value":"42"}`), &result)
}
```

> **Note**: This should only be done during program initialization as it's not safe for concurrent modification.

### Code Examples

```go
package main

import (
	"github.com/flyhope/json-fuzzy"
	"github.com/flyhope/json-fuzzy/fuzzy"
)

func main() {
	// Using basic mode (FuzzyUnmarshalerOrigin) for better performance
	// when you don't need 'any' type support
	basicOpts := jsonfuzzy.WithOptions(fuzzy.FuzzyUnmarshalerOrigin())
	var basicResult struct{ Value int }
	err := jsonfuzzy.Unmarshal([]byte(`{"Value":"42"}`), &basicResult, basicOpts)
	
	// Using full mode (default) which supports 'any' type
	var fullResult struct{ Value any }
	err = jsonfuzzy.Unmarshal([]byte(`{"Value":true}`), &fullResult)
	// Equivalent to:
	// err = jsonfuzzy.Unmarshal([]byte(`{"Value":true}`), &fullResult, fuzzy.FuzzyUnmarshalerFull())
}
```

## Installation

```bash
go get github.com/flyhope/json-fuzzy
```

## Usage

The library provides functions that mirror the standard library's JSON API but with fuzzy parsing capabilities:

```go
package main

import (
	"github.com/flyhope/json-fuzzy"
)

func main() {
	type testStruct struct {
		StringField string  `json:"string_field"`
		IntField    int     `json:"int_field"`
		FloatField  float64 `json:"float_field"`
		BoolField   bool    `json:"bool_field"`
	}

	// JSON input with mismatched types
	jsonData := []byte(`{
		"string_field": 123,
		"int_field":    "456",
		"float_field":  "78.9",
		"bool_field":   "true"
	}`)

	var result testStruct
	err := jsonfuzzy.Unmarshal(jsonData, &result)
	// This will succeed even though types don't match the struct fields
}
```

## API

- `Unmarshal(data []byte, v any, opts ...json.Options) error`: Unmarshals JSON data with fuzzy type conversion
- `Marshal(v any, opts ...json.Options) ([]byte, error)`: Marshals Go values to JSON (standard behavior)
- `MarshalString(v any, opts ...json.Options) (string, error)`: Marshals Go values to JSON string
- `WithOptions(opts ...json.Options) json.Options`: Combines provided options with default fuzzy options
- `DefaultOptions() json.Options`: Returns the default options including fuzzy unmarshalers and lenient parsing settings

## License

MIT License