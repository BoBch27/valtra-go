package valtra

// Value holds a value to be validated along with and any
// validation errors that occur during validation.
type Value[T any] struct {
	value T
	errs  []error
}

// Val creates a new Value[T] that wraps a value.
//
// This is the entry point for validation. Use the returned
// Value's Validate method to apply validation rules.
//
// Example:
//
//	v := valtra.Val(25).Validate(valtra.Max(30))
func Val[T any](value T) Value[T] {
	return Value[T]{
		value: value,
		errs:  []error{},
	}
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
//	v := valtra.Val("test@example.com").Validate(valtra.Email())
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
//	v := valtra.Val(25).Validate(
//	    valtra.Required[int](),
//	    valtra.Min(20),
//	    valtra.Max(30),
//	)
func (v Value[T]) Validate(validations ...func(Value[T]) error) Value[T] {
	for _, fn := range validations {
		err := fn(v)
		if err != nil {
			v.errs = append(v.errs, err)
		}
	}

	return v
}
