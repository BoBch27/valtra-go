package valtra

import "strings"

// Uppercase returns a transformation that converts the
// value to upper case.
//
// Example:
//
//	valtra.Val("ok").Transform(valtra.Uppercase())
func Uppercase() func(Value[string]) (string, error) {
	return func(v Value[string]) (string, error) {
		return strings.ToUpper(v.value), nil
	}
}

// Lowercase returns a transformation that converts the
// value to lower case.
//
// Example:
//
//	valtra.Val("JOHN@EXAMPLE.COM").Transform(valtra.Lowercase())
func Lowercase() func(Value[string]) (string, error) {
	return func(v Value[string]) (string, error) {
		return strings.ToLower(v.value), nil
	}
}

// TrimSpace returns a transformation that removes all
// leading and trailing white space from the value.
//
// Example:
//
//	valtra.Val(" john ").Transform(valtra.TrimSpace())
func TrimSpace() func(Value[string]) (string, error) {
	return func(v Value[string]) (string, error) {
		return strings.TrimSpace(v.value), nil
	}
}

// Capitalise returns a transformation that converts the
// first character of the value to upper case.
//
// Example:
//
//	valtra.Val("john").Transform(valtra.Capitalise())
func Capitalise() func(Value[string]) (string, error) {
	return func(v Value[string]) (string, error) {
		return strings.ToUpper(v.value[:1]) + strings.ToLower(v.value[1:]), nil
	}
}
