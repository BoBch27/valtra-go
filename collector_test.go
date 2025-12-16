package valtra_test

import (
	"testing"

	"github.com/bobch27/valtra-go"
)

func TestCollector(t *testing.T) {
	t.Run("collects errors from multiple values", func(t *testing.T) {
		c := valtra.NewCollector()

		name := valtra.Val("").Validate(valtra.Required[string]()).Collect(c)
		age := valtra.Val(15).Validate(valtra.Min(18)).Collect(c)

		if c.IsValid() {
			t.Error("Collector should have errors")
		}

		if len(c.Errors()) != 2 {
			t.Errorf("Expected 2 errors, got %d: %v", len(c.Errors()), c.Errors())
		}

		// Verify collected values are returned correctly
		if name != "" {
			t.Errorf("Expected empty string, got %q", name)
		}
		if age != 15 {
			t.Errorf("Expected 15, got %d", age)
		}
	})

	t.Run("collector with no errors", func(t *testing.T) {
		c := valtra.NewCollector()

		name := valtra.Val(" John ").Validate(valtra.Required[string]()).Transform(valtra.TrimSpace()).Collect(c)
		age := valtra.Val(25).Validate(valtra.Min(18)).Collect(c)

		if !c.IsValid() {
			t.Errorf("Collector should be valid, got errors: %v", c.Errors())
		}

		if len(c.Errors()) != 0 {
			t.Errorf("Expected 0 errors, got %d", len(c.Errors()))
		}

		// Verify collected values
		if name != "John" {
			t.Errorf("Expected 'John', got %q", name)
		}
		if age != 25 {
			t.Errorf("Expected 25, got %d", age)
		}
	})
}
