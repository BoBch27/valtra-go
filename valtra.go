package valtra

import (
	"fmt"
)

type Value[T any] struct {
	value T
	errs  []error
}

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

func Required[T comparable]() func(Value[T]) error {
	return func(v Value[T]) error {
		var zero T
		if v.value == zero {
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

func Max[T Ordered](max T) func(Value[T]) error {
	return func(v Value[T]) error {
		if v.value > max {
			return fmt.Errorf("value cannot be larger than %v", max)
		}

		return nil
	}
}

func Min[T Ordered](min T) func(Value[T]) error {
	return func(v Value[T]) error {
		if v.value < min {
			return fmt.Errorf("value cannot be smaller than %v", min)
		}

		return nil
	}
}
