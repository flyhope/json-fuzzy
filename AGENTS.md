# AGENTS.md

## Project

Go library `github.com/flyhope/json-fuzzy` for lenient JSON unmarshaling, built on the stable `encoding/json/v2` (requires **Go 1.27+**; no `GOEXPERIMENT` needed). Two packages only:

- root (`jsonfuzzy.go`): public API mirroring stdlib json (`Unmarshal`, `Marshal`, `WithOptions`, ...)
- `fuzzy/`: custom unmarshalers registered via `json.WithUnmarshalers`; default mode is set by the package-level var `FuzzyUnmarshaler = FuzzyUnmarshalerFull`

## Commands

```
go build ./...
go test ./...
go vet ./...        # CI runs all three plus golangci-lint (v2.x)
golangci-lint run   # not installed locally; enforced in .github/workflows/test.yml
```

No codegen, no external services. Tests are plain `go test`.

## Gotchas

- Requires Go >= 1.27. The experimental `json.SkipFunc` sentinel no longer exists in stable `encoding/json/v2`; unmarshalers must return `errors.ErrUnsupported` to fall back to default decoding.
- Use `reflect.Pointer`, not deprecated `reflect.Ptr` (govet/golangci-lint fails CI on it).
- Unmarshalers that skip must not consume from the `jsontext.Decoder` before returning `errors.ErrUnsupported`.
- `FuzzyUnmarshaler*` are `sync.OnceValue`-wrapped options; changing global default mode is init-time only.

## Git workflow

- Default branch: `main`; work happens on `develop`, PRs go `develop` -> `main`.
- README.md and README_zh.md should be kept in sync when behavior changes.
