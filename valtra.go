package valtra

import (
	"fmt"
	"regexp"
)

// Value holds a value to be validated along with and any
// validation errors that occur during validation.
type Value[T any] struct {
	value T
	errs  []error
}

// Value returns the value being validated.
func (v Value[T]) Value() T {
	return v.value
}

// Errors returns all validation errors that have occurred.
// Returns an empty slice if validation passed.
func (v Value[T]) Errors() []error {
	return v.errs
}

// IsValid returns true if there are no validation errors,
// false otherwise.
//
// This is a convenience method equivalent to checking
// len(v.Errors()) == 0.
//
// Example:
//
//	v := valtra.Validate("test@example.com", valtra.Email())
//	if v.IsValid() {
//	    email := v.Value()
//	    // proceed with valid email
//	}
func (v Value[T]) IsValid() bool {
	return len(v.errs) == 0
}

// Validate applies all provided validation functions for
// the given value.
//
// Each validation function that returns an
// error will add that error to the value's error list.
//
// Example:
//
//	v := valtra.Validate(25,
//	    valtra.Required[int](),
//	    valtra.Min(20),
//	    valtra.Max(30),
//	)
func Validate[T any](value T, validations ...func(Value[T]) error) Value[T] {
	v := Value[T]{
		value: value,
		errs:  []error{},
	}

	for _, fn := range validations {
		err := fn(v)
		if err != nil {
			v.errs = append(v.errs, err)
		}
	}

	return v
}

// Required returns a validation that ensures the value is
// not the zero value for its type.
//
// For strings, this means non-empty. For numbers, this means
// non-zero. For pointers, this means non-nil.
//
// Example:
//
//	valtra.Validate("", valtra.Required[string]())  // fails
//	valtra.Validate("John", valtra.Required[string]())  // passes
func Required[T comparable]() func(Value[T]) error {
	return func(v Value[T]) error {
		var zero T
		if v.value == zero {
			return fmt.Errorf("value is required")
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
// Example:
//
//	valtra.Validate(100, valtra.Max(100))
func Max[T Ordered](max T) func(Value[T]) error {
	return func(v Value[T]) error {
		if v.value > max {
			return fmt.Errorf("value cannot be larger than %v", max)
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
// Example:
//
//	valtra.Validate(5, valtra.Min(1))
func Min[T Ordered](min T) func(Value[T]) error {
	return func(v Value[T]) error {
		if v.value < min {
			return fmt.Errorf("value cannot be smaller than %v", min)
		}

		return nil
	}
}

// MaxLengthString returns a validation that ensures the
// length of a string does not exceed the given maximum.
//
// Example:
//
//	valtra.Validate("username", valtra.MaxLengthString(20))
func MaxLengthString[T ~string](max int) func(Value[T]) error {
	return func(v Value[T]) error {
		if len(v.value) > max {
			return fmt.Errorf("value's length cannot be larger than %v", max)
		}

		return nil
	}
}

// MaxLengthSlice returns a validation that ensures the
// length of a slice does not exceed the given maximum.
//
// Example:
//
//	valtra.Validate([]int{1}, valtra.MaxLengthSlice(2))
func MaxLengthSlice[T any](max int) func(Value[[]T]) error {
	return func(v Value[[]T]) error {
		if len(v.value) > max {
			return fmt.Errorf("value's length cannot be larger than %v", max)
		}

		return nil
	}
}

// MaxLengthMap returns a validation that ensures the
// length of a map does not exceed the given maximum.
//
// Example:
//
//	valtra.Validate(map[string]int{"no": 1}, valtra.MaxLengthMap(2))
func MaxLengthMap[K comparable, V any](max int) func(Value[map[K]V]) error {
	return func(v Value[map[K]V]) error {
		if len(v.value) > max {
			return fmt.Errorf("value's length cannot be larger than %v", max)
		}

		return nil
	}
}

// MinLengthString returns a validation that ensures the
// length of a string is at least the given minimum.
//
// Example:
//
//	valtra.Validate("username", valtra.MinLengthString(5))
func MinLengthString[T ~string](min int) func(Value[T]) error {
	return func(v Value[T]) error {
		if len(v.value) < min {
			return fmt.Errorf("value's length cannot be smaller than %v", min)
		}

		return nil
	}
}

// MinLengthSlice returns a validation that ensures the
// length of a slice is at least the given minimum.
//
// Example:
//
//	valtra.Validate([]int{1}, valtra.MinLengthSlice(1))
func MinLengthSlice[T any](min int) func(Value[[]T]) error {
	return func(v Value[[]T]) error {
		if len(v.value) < min {
			return fmt.Errorf("value's length cannot be smaller than %v", min)
		}

		return nil
	}
}

// MinLengthMap returns a validation that ensures the
// length of a map is at least the given minimum.
//
// Example:
//
//	valtra.Validate(map[string]int{"no": 1}, valtra.MinLengthMap(1))
func MinLengthMap[K comparable, V any](min int) func(Value[map[K]V]) error {
	return func(v Value[map[K]V]) error {
		if len(v.value) < min {
			return fmt.Errorf("value's length cannot be smaller than %v", min)
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
// Example:
//
//	valtra.Validate("user@example.com",, valtra.Email())
func Email() func(Value[string]) error {
	return func(v Value[string]) error {
		if !emailRegex.MatchString(v.value) {
			return fmt.Errorf("value must be in correct email format")
		}

		return nil
	}
}
