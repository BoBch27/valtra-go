package valtra_test

import (
	"testing"

	"github.com/bobch27/valtra-go"
)

func TestRequired(t *testing.T) {
	t.Run("empty string fails", func(t *testing.T) {
		v := valtra.Val("").Validate(valtra.Required[string]())
		if v.IsValid() {
			t.Error("Expected validation to fail for empty string")
		}
	})

	t.Run("non-empty string passes", func(t *testing.T) {
		v := valtra.Val("hello").Validate(valtra.Required[string]())
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("zero int fails", func(t *testing.T) {
		v := valtra.Val(0).Validate(valtra.Required[int]())
		if v.IsValid() {
			t.Error("Expected validation to fail for zero int")
		}
	})
}

func TestMin(t *testing.T) {
	t.Run("below min fails", func(t *testing.T) {
		v := valtra.Val(5).Validate(valtra.Min(10))
		if v.IsValid() {
			t.Error("Expected validation to fail for value below min")
		}
	})

	t.Run("at min passes", func(t *testing.T) {
		v := valtra.Val(10).Validate(valtra.Min(10))
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("above min passes", func(t *testing.T) {
		v := valtra.Val(15).Validate(valtra.Min(10))
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})
}

func TestMax(t *testing.T) {
	t.Run("above max fails", func(t *testing.T) {
		v := valtra.Val(15).Validate(valtra.Max(10))
		if v.IsValid() {
			t.Error("Expected validation to fail for value above max")
		}
	})

	t.Run("at max passes", func(t *testing.T) {
		v := valtra.Val(10).Validate(valtra.Max(10))
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})
}

func TestMinLengthString(t *testing.T) {
	t.Run("below min length fails", func(t *testing.T) {
		v := valtra.Val("ab").Validate(valtra.MinLengthString(5))
		if v.IsValid() {
			t.Error("Expected validation to fail for string below min length")
		}
	})

	t.Run("at min length passes", func(t *testing.T) {
		v := valtra.Val("hello").Validate(valtra.MinLengthString(5))
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})
}

func TestMaxLengthString(t *testing.T) {
	t.Run("above max length fails", func(t *testing.T) {
		v := valtra.Val("hello").Validate(valtra.MaxLengthString(3))
		if v.IsValid() {
			t.Error("Expected validation to fail for string above max length")
		}
	})

	t.Run("at max length passes", func(t *testing.T) {
		v := valtra.Val("abc").Validate(valtra.MaxLengthString(3))
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})
}

func TestMinLengthSlice(t *testing.T) {
	t.Run("below min length fails", func(t *testing.T) {
		v := valtra.Val([]int{1}).Validate(valtra.MinLengthSlice[int](2))
		if v.IsValid() {
			t.Error("Expected validation to fail for slice below min length")
		}
	})

	t.Run("at min length passes", func(t *testing.T) {
		v := valtra.Val([]int{1, 2}).Validate(valtra.MinLengthSlice[int](2))
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})
}

func TestMaxLengthSlice(t *testing.T) {
	t.Run("above max length fails", func(t *testing.T) {
		v := valtra.Val([]int{1, 2, 3}).Validate(valtra.MaxLengthSlice[int](2))
		if v.IsValid() {
			t.Error("Expected validation to fail for slice above max length")
		}
	})

	t.Run("at max length passes", func(t *testing.T) {
		v := valtra.Val([]int{1, 2}).Validate(valtra.MaxLengthSlice[int](2))
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})
}

func TestMinLengthMap(t *testing.T) {
	t.Run("below min length fails", func(t *testing.T) {
		v := valtra.Val(map[string]int{"a": 1}).Validate(valtra.MinLengthMap[string, int](2))
		if v.IsValid() {
			t.Error("Expected validation to fail for map below min length")
		}
	})

	t.Run("at min length passes", func(t *testing.T) {
		v := valtra.Val(map[string]int{"a": 1, "b": 2}).Validate(valtra.MinLengthMap[string, int](2))
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})
}

func TestMaxLengthMap(t *testing.T) {
	t.Run("above max length fails", func(t *testing.T) {
		v := valtra.Val(map[string]int{"a": 1, "b": 2, "c": 3}).Validate(valtra.MaxLengthMap[string, int](2))
		if v.IsValid() {
			t.Error("Expected validation to fail for map above max length")
		}
	})

	t.Run("at max length passes", func(t *testing.T) {
		v := valtra.Val(map[string]int{"a": 1, "b": 2}).Validate(valtra.MaxLengthMap[string, int](2))
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})
}

func TestEmail(t *testing.T) {
	t.Run("valid email passes", func(t *testing.T) {
		v := valtra.Val("test@example.com").Validate(valtra.Email())
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("invalid email fails", func(t *testing.T) {
		v := valtra.Val("not-an-email").Validate(valtra.Email())
		if v.IsValid() {
			t.Error("Expected validation to fail for invalid email")
		}
	})

	t.Run("email with unicode passes", func(t *testing.T) {
		v := valtra.Val("tëst@example.com").Validate(valtra.Email())
		if !v.IsValid() {
			t.Errorf("Expected validation to pass for unicode email, got errors: %v", v.Errors())
		}
	})
}

func TestMultipleValidations(t *testing.T) {
	t.Run("accumulates multiple errors", func(t *testing.T) {
		v := valtra.Val("ab").Validate(
			valtra.Required[string](),
			valtra.MinLengthString(5),
			valtra.MaxLengthString(1),
		)

		if len(v.Errors()) != 2 {
			t.Errorf("Expected 2 errors (min and max), got %d: %v", len(v.Errors()), v.Errors())
		}
	})

	t.Run("all validations pass", func(t *testing.T) {
		v := valtra.Val("hello").Validate(
			valtra.Required[string](),
			valtra.MinLengthString(3),
			valtra.MaxLengthString(10),
		)

		if !v.IsValid() {
			t.Errorf("Expected all validations to pass, got errors: %v", v.Errors())
		}
	})
}

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
}
