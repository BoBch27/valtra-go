package valtra_test

import (
	"testing"

	"github.com/bobch27/valtra-go"
)

func TestValueMethods(t *testing.T) {
	t.Run("Value() returns the wrapped value", func(t *testing.T) {
		v := valtra.Val(42)
		if v.Value() != 42 {
			t.Errorf("Expected 42, got %d", v.Value())
		}
	})

	t.Run("Name() returns custom name", func(t *testing.T) {
		v := valtra.Val("test", "username")
		if v.Name() != "username" {
			t.Errorf("Expected 'username', got %q", v.Name())
		}
	})

	t.Run("Name() returns default when not provided", func(t *testing.T) {
		v := valtra.Val("test")
		if v.Name() != "value" {
			t.Errorf("Expected 'value', got %q", v.Name())
		}
	})

	t.Run("Errors() returns empty slice when valid", func(t *testing.T) {
		v := valtra.Val(10).Validate(valtra.Min(5))
		if len(v.Errors()) != 0 {
			t.Errorf("Expected 0 errors, got %d", len(v.Errors()))
		}
	})

	t.Run("IsValid() returns true when valid", func(t *testing.T) {
		v := valtra.Val(10).Validate(valtra.Min(5))
		if !v.IsValid() {
			t.Errorf("Expected valid value, got %t", v.IsValid())
		}
	})

	t.Run("IsValid() returns false when invalid", func(t *testing.T) {
		v := valtra.Val(1).Validate(valtra.Min(5))
		if v.IsValid() {
			t.Errorf("Expected invalid value, got %t", v.IsValid())
		}
	})
}
