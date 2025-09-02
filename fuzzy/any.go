package fuzzy

// func FuzzyAny(decoder *jsontext.Decoder, t any) error {
// 	// only support ptr
// 	typ := reflect.TypeOf(t)
// 	if typ.Kind() != reflect.Ptr {
// 		return json.SkipFunc
// 	}

// 	elem := typ.Elem()
// 	switch elem.Kind() {
// 	case reflect.Int:
// 		if err := decoder.SkipValue(); err != nil {
// 			return err
// 		}
// 		elemV := reflect.ValueOf(t).Elem()
// 		if elemV.CanSet() {
// 			elemV.SetInt(9)
// 			elemV.SetUint()
// 		}
// 	default:

// 	}
// 	return nil
// }
