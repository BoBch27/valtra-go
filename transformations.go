package valtra

import "strings"

// Uppercase returns a transformation that converts the
// value to upper case.
//
// Example:
//
//	valtra.Val("").Transform(valtra.Uppercase())
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
//	valtra.Val("").Transform(valtra.Lowercase())
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
//	valtra.Val("").Transform(valtra.TrimSpace())
func TrimSpace() func(Value[string]) (string, error) {
	return func(v Value[string]) (string, error) {
		return strings.TrimSpace(v.value), nil
	}
}
