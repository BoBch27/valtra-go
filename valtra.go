package valtra

import (
	"fmt"
)

// Value holds a value to be validated along with and any
// validation errors that occur during validation.
type Value[T any] struct {
	value T
	errs  []error
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
