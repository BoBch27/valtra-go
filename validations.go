package valtra

import (
	"fmt"
	"regexp"
)

// Required returns a validation that ensures the value is
// not the zero value for its type.
//
// For strings, this means non-empty. For numbers, this means
// non-zero. For pointers, this means non-nil.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	valtra.Val("").Validate(valtra.Required[string]())  // fails
//	valtra.Val("John").Validate(valtra.Required[string]())  // passes
func Required[T comparable](errMssg ...string) func(Value[T]) error {
	return func(v Value[T]) error {
		var zero T
		if v.value == zero {
			// Return custom error message, if provided
			if len(errMssg) > 0 && errMssg[0] != "" {
				return fmt.Errorf("%s", errMssg[0])
			}

			return fmt.Errorf("%s is required", v.name)
		}

		return nil
	}
}

// Ordered is a constraint that permits all numeric types
// that support comparison operations (<, >, <=, >=).
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Max returns a validation that ensures the value does
// not exceed the given maximum.
//
// Works with all numeric types defined by the Ordered
// constraint.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	valtra.Val(100).Validate(valtra.Max(100))
func Max[T Ordered](max T, errMssg ...string) func(Value[T]) error {
	return func(v Value[T]) error {
		if v.value > max {
			// Return custom error message, if provided
			if len(errMssg) > 0 && errMssg[0] != "" {
				return fmt.Errorf("%s", errMssg[0])
			}

			return fmt.Errorf("%s cannot be larger than %v", v.name, max)
		}

		return nil
	}
}

// Min returns a validation that ensures the value is
// at least the given minimum.
//
// Works with all numeric types defined by the Ordered
// constraint.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	valtra.Val(5).Validate(valtra.Min(1))
func Min[T Ordered](min T, errMssg ...string) func(Value[T]) error {
	return func(v Value[T]) error {
		if v.value < min {
			// Return custom error message, if provided
			if len(errMssg) > 0 && errMssg[0] != "" {
				return fmt.Errorf("%s", errMssg[0])
			}

			return fmt.Errorf("%s cannot be smaller than %v", v.name, min)
		}

		return nil
	}
}

// MaxLengthString returns a validation that ensures the
// length of a string does not exceed the given maximum.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	valtra.Val("username").Validate(valtra.MaxLengthString(20))
func MaxLengthString(max int, errMssg ...string) func(Value[string]) error {
	return func(v Value[string]) error {
		if len(v.value) > max {
			// Return custom error message, if provided
			if len(errMssg) > 0 && errMssg[0] != "" {
				return fmt.Errorf("%s", errMssg[0])
			}

			return fmt.Errorf("%s's length cannot be larger than %v", v.name, max)
		}

		return nil
	}
}

// MaxLengthSlice returns a validation that ensures the
// length of a slice does not exceed the given maximum.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	valtra.Val([]int{1}).Validate(valtra.MaxLengthSlice(2))
func MaxLengthSlice[T any](max int, errMssg ...string) func(Value[[]T]) error {
	return func(v Value[[]T]) error {
		if len(v.value) > max {
			// Return custom error message, if provided
			if len(errMssg) > 0 && errMssg[0] != "" {
				return fmt.Errorf("%s", errMssg[0])
			}

			return fmt.Errorf("%s's length cannot be larger than %v", v.name, max)
		}

		return nil
	}
}

// MaxLengthMap returns a validation that ensures the
// length of a map does not exceed the given maximum.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	valtra.Val(map[string]int{"no": 1}).Validate(valtra.MaxLengthMap(2))
func MaxLengthMap[K comparable, V any](max int, errMssg ...string) func(Value[map[K]V]) error {
	return func(v Value[map[K]V]) error {
		if len(v.value) > max {
			// Return custom error message, if provided
			if len(errMssg) > 0 && errMssg[0] != "" {
				return fmt.Errorf("%s", errMssg[0])
			}

			return fmt.Errorf("%s's length cannot be larger than %v", v.name, max)
		}

		return nil
	}
}

// MinLengthString returns a validation that ensures the
// length of a string is at least the given minimum.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	valtra.Val("username").Validate(valtra.MinLengthString(5))
func MinLengthString(min int, errMssg ...string) func(Value[string]) error {
	return func(v Value[string]) error {
		if len(v.value) < min {
			// Return custom error message, if provided
			if len(errMssg) > 0 && errMssg[0] != "" {
				return fmt.Errorf("%s", errMssg[0])
			}

			return fmt.Errorf("%s's length cannot be smaller than %v", v.name, min)
		}

		return nil
	}
}

// MinLengthSlice returns a validation that ensures the
// length of a slice is at least the given minimum.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	valtra.Val([]int{1}).Validate(valtra.MinLengthSlice(1))
func MinLengthSlice[T any](min int, errMssg ...string) func(Value[[]T]) error {
	return func(v Value[[]T]) error {
		if len(v.value) < min {
			// Return custom error message, if provided
			if len(errMssg) > 0 && errMssg[0] != "" {
				return fmt.Errorf("%s", errMssg[0])
			}

			return fmt.Errorf("%s's length cannot be smaller than %v", v.name, min)
		}

		return nil
	}
}

// MinLengthMap returns a validation that ensures the
// length of a map is at least the given minimum.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	valtra.Val(map[string]int{"no": 1}).Validate(valtra.MinLengthMap(1))
func MinLengthMap[K comparable, V any](min int, errMssg ...string) func(Value[map[K]V]) error {
	return func(v Value[map[K]V]) error {
		if len(v.value) < min {
			// Return custom error message, if provided
			if len(errMssg) > 0 && errMssg[0] != "" {
				return fmt.Errorf("%s", errMssg[0])
			}

			return fmt.Errorf("%s's length cannot be smaller than %v", v.name, min)
		}

		return nil
	}
}

// emailRegex is a practical, internationally-aware email format.
// Supports Unicode characters (accents, non-Latin scripts)
// in email addresses.
var emailRegex = regexp.MustCompile(`^(?:"(?:[^"]|\\")*"|[\p{L}\p{N}\p{M}._%+-]+)@[\p{L}\p{N}\p{M}.-]+\.[\p{L}\p{M}]{2,}$`)

// Email returns a validation that ensures the value
// is a valid email address.
//
// It uses a practical, internationally-aware pattern
// that catches common errors, while remaining permissive.
//
// For true validation, send a confirmation email.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	valtra.Val("user@example.com").Validate(valtra.Email())
func Email(errMssg ...string) func(Value[string]) error {
	return func(v Value[string]) error {
		if !emailRegex.MatchString(v.value) {
			// Return custom error message, if provided
			if len(errMssg) > 0 && errMssg[0] != "" {
				return fmt.Errorf("%s", errMssg[0])
			}

			return fmt.Errorf("%s must be in correct email format", v.name)
		}

		return nil
	}
}
