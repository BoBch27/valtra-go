package valtra

import "fmt"

func Validate[T any](value T, validations ...func(T) error) []error {
	errs := []error{}

	for _, fn := range validations {
		err := fn(value)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

func Required[T comparable]() func(T) error {
	return func(v T) error {
		var zero T
		if v == zero {
			return fmt.Errorf("value is required")
		}

		return nil
	}
}
