package valtra

type Collector struct {
	errs []error
}

func NewCollector() *Collector {
	return &Collector{errs: []error{}}
}

func (c *Collector) Errors() []error {
	return c.errs
}

func (c *Collector) IsValid() bool {
	return len(c.errs) == 0
}
