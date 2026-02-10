package valtra_test

import (
	"fmt"
	"testing"

	"github.com/bobch27/valtra-go"
)

func TestUppercase(t *testing.T) {
	t.Run("string to upper case", func(t *testing.T) {
		v := valtra.Val("hello").Transform(valtra.Uppercase())
		if v.Value() != "HELLO" {
			t.Errorf("Expected transformation to pass, got errors: %v", v.Errors())
		}
	})
}

func TestLowercase(t *testing.T) {
	t.Run("string to lower case", func(t *testing.T) {
		v := valtra.Val("HELLO").Transform(valtra.Lowercase())
		if v.Value() != "hello" {
			t.Errorf("Expected transformation to pass, got errors: %v", v.Errors())
		}
	})
}

func TestTrimSpace(t *testing.T) {
	t.Run("string with trimmed space", func(t *testing.T) {
		v := valtra.Val(" hello ").Transform(valtra.TrimSpace())
		if v.Value() != "hello" {
			t.Errorf("Expected transformation to pass, got errors: %v", v.Errors())
		}
	})
}

func TestCapitalise(t *testing.T) {
	t.Run("string with capital first letter", func(t *testing.T) {
		v := valtra.Val("bobby").Transform(valtra.Capitalise())
		if v.Value() != "Bobby" {
			t.Errorf("Expected transformation to pass, got errors: %v", v.Errors())
		}
	})
}

func TestMultipleTransformations(t *testing.T) {
	t.Run("all transformations pass", func(t *testing.T) {
		v := valtra.Val(" hello ").Transform(
			valtra.TrimSpace(),
			valtra.Uppercase(),
		)

		if !v.IsValid() {
			t.Errorf("Expected all transformations to pass, got errors: %v", v.Errors())
		}
		if v.Value() != "HELLO" {
			t.Errorf("Expected transformed value, got: %s", v.Value())
		}
	})
}

func TestTransformWithError(t *testing.T) {
	t.Run("transformation error is collected", func(t *testing.T) {
		// Custom transformation that returns an error
		v := valtra.Val("hello").Transform(func(v valtra.Value[string]) (string, error) {
			return "", fmt.Errorf("transformation failed")
		})

		if v.IsValid() {
			t.Error("Expected transformation to fail")
		}

		if len(v.Errors()) != 1 {
			t.Errorf("Expected 1 error, got %d", len(v.Errors()))
		}

		// Verify value remains unchanged when transformation fails
		if v.Value() != "hello" {
			t.Errorf("Expected value to remain 'hello', got %q", v.Value())
		}
	})
}
