package valtra

// Collector accumulates validation errors from multiple
// Value instances.
//
// It is useful when validating the fields of a struct,
// allowing all validation errors to be gathered and returned
// at once, instead of manually checking each value result.
//
// Collectors are created with NewCollector and are updated
// via the Collect method on a Value.
type Collector struct {
	errs []error
}

// NewCollector creates and returns a new Collector with
// an empty error list.
//
// Collector accumulates validation errors from multiple
// Value instances.
//
// It is useful when validating the fields of a struct,
// allowing all validation errors to be gathered and returned
// at once, instead of manually checking each value result.
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
func NewCollector() *Collector {
	return &Collector{errs: []error{}}
}

// Errors returns all accumulated validation errors.
// Returns an empty slice if no errors were collected.
func (c *Collector) Errors() []error {
	return c.errs
}

// IsValid returns true if no validation errors have been
// collected, or false otherwise.
//
// This is a convenience method equivalent to checking
// len(v.Errors()) == 0.
func (c *Collector) IsValid() bool {
	return len(c.errs) == 0
}
