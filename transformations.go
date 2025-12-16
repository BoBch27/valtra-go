package valtra

import "strings"

func Uppercase() func(Value[string]) (string, error) {
	return func(v Value[string]) (string, error) {
		return strings.ToUpper(v.value), nil
	}
}

func Lowercase() func(Value[string]) (string, error) {
	return func(v Value[string]) (string, error) {
		return strings.ToLower(v.value), nil
	}
}

func TrimSpace() func(Value[string]) (string, error) {
	return func(v Value[string]) (string, error) {
		return strings.TrimSpace(v.value), nil
	}
}
