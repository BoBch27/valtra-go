package valtra

// Value holds a value to be validated/transformed, along
// with its name and any errors that occur during
// validation/transformation.
type Value[T any] struct {
	value T
	name  string
	errs  []error
}

// Val creates a new Value[T] that wraps a value.
//
// This is the entry point for validation and transformation.
//
// Use the returned Value's Validate method to apply
// validation rules.
//
// Use the returned Value's Transform method to apply
// transformations.
//
// The optional name parameter is used in error messages to
// identify which value failed validation/transformation.
// Default is "value".
//
// Example:
//
//	valtra.Val("bobby").Validate(valtra.Required[string]()).Transform(valtra.Uppercase())
func Val[T any](value T, name ...string) Value[T] {
	valName := "value"
	if len(name) > 0 && name[0] != "" {
		valName = name[0]
	}

	return Value[T]{
		value: value,
		name:  valName,
		errs:  []error{},
	}
}

// Value returns the value being validated/transformed.
func (v Value[T]) Value() T {
	return v.value
}

// Name returns the value's name, which is used to identify
// the value in validation/transformation errors.
func (v Value[T]) Name() string {
	return v.name
}

// Errors returns all errors that have occurred.
// Returns an empty slice if validation/transformation passed.
func (v Value[T]) Errors() []error {
	return v.errs
}

// IsValid returns true if there are no errors,
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

// Transform applies all provided transformation
// functions to the given value.
//
// Each transformation function that returns an
// error will add that error to the value's error list.
//
// Example:
//
//	v := valtra.Val("hello").Transform(valtra.Uppercase())
func (v Value[T]) Transform(transformations ...func(Value[T]) (T, error)) Value[T] {
	for _, fn := range transformations {
		newVal, err := fn(v)
		if err != nil {
			v.errs = append(v.errs, err)
		} else {
			v.value = newVal
		}
	}

	return v
}

// Collect appends all errors from the Value
// into the provided Collector and returns the underlying
// validated and transformed value.
//
// This allows multiple Value instances to contribute their
// results when validating/transforming the fields of a struct.
//
// Example:
//
//	c := valtra.NewCollector()
//	user := User{
//		name := valtra.Val(input.Name).Validate(valtra.Required[string]()).Collect(c)
//		age := valtra.Val(input.Age).Validate(valtra.Min(18)).Collect(c)
//	}
//	if !c.IsValid() {
//	    return c.Errors()
//	}
func (v Value[T]) Collect(c *Collector) T {
	c.errs = append(c.errs, v.errs...)
	return v.value
}
