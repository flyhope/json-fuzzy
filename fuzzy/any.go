package fuzzy

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"reflect"
)

func FuzzyAny(dec *jsontext.Decoder, t any) error {
	// only support ptr
	typ := reflect.TypeOf(t)
	if typ.Kind() != reflect.Ptr {
		return json.SkipFunc
	}

	elem := typ.Elem()
	switch elem.Kind() {

	// parse int
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		elemV := reflect.ValueOf(t).Elem()
		if !elemV.CanSet() {
			return json.SkipFunc
		}

		v, kind, ok, err := showIntegerByPeek[int64](dec)
		if ok || err != nil {
			if ok {
				elemV.SetInt(v)
			}
			return err
		}

		result, err := parseInt64(kind, dec)
		if err != nil {
			return err
		}
		elemV.SetInt(result)

	// parse uint
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		elemV := reflect.ValueOf(t).Elem()
		if !elemV.CanSet() {
			return json.SkipFunc
		}

		v, kind, ok, err := showIntegerByPeek[uint64](dec)
		if ok || err != nil {
			if ok {
				elemV.SetUint(v)
			}
			return err
		}

		result, err := parseUint64(kind, dec)
		if err != nil {
			return err
		}
		elemV.SetUint(result)

	// parse string
	case reflect.String:
		result, err := showStringValue(dec)
		if err != nil {
			return err
		}

		elemV := reflect.ValueOf(t).Elem()
		if !elemV.CanSet() {
			return json.SkipFunc
		}
		elemV.SetString(result)

	// parse bool
	case reflect.Bool:
		result, err := parseBool(dec)
		if err != nil {
			return err
		}

		elemV := reflect.ValueOf(t).Elem()
		if !elemV.CanSet() {
			return json.SkipFunc
		}
		elemV.SetBool(result)

	default:
		return json.SkipFunc
	}

	return nil
}
