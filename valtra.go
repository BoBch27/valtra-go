package valtra

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
