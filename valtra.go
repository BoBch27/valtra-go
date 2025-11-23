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

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Max[T Ordered](max T) func(T) error {
	return func(v T) error {
		if v > max {
			return fmt.Errorf("value cannot be larger than %v", max)
		}

		return nil
	}
}

func Min[T Ordered](min T) func(T) error {
	return func(v T) error {
		if v < min {
			return fmt.Errorf("value cannot be smaller than %v", min)
		}

		return nil
	}
}
