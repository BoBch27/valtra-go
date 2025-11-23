# Valtra

A type-safe, performant validation library for Go that uses generics and functional composition. No reflection, no struct tags - just clean, declarative validation! ⚡

## Why Yet Another Validation Package?

Ever spent the night chasing a bug caused by a typo in a struct tag? 😴 I know I have... 🤦

[go-playground/validator](https://github.com/go-playground/validator) is great. Really. In fact, I even built a [Firestore ODM](https://github.com/bobch27/firevault_go) with a validation engine that works similarly (though differently under the hood), and it matches validator's performance almost exactly. 

But reflection and string-based struct tags still lead to **runtime errors** when they should be **compile-time errors**. 🤷 That’s just not the Go way.

Enter **Valtra**. It’s hardly even a package - just clever use of generics that let you validate data **declaratively and safely**. No reflection, no string parsing, just functions, types, and the compiler doing its job. ⚡

## Features

- **🔒 Type-safe**: Your IDE will yell at you before the compiler does. The compiler will yell at you before your users do.
- **⚡ Zero reflection**: Because it's 2025 and we have generics now.
- **🎯 Declarative**: Reads like English, works like magic (except it's not magic, it's just good design).
- **🔧 Composable**: Build complex validators from simple functions, like LEGO but for paranoid backend developers.
- **📦 Minimal**: Small enough to read in one sitting, powerful enough to actually use.
- **🚀 Fast**: Faster than reflection-based validators. Your users won't notice, but your benchmarks will look great.

## Installation

```bash
go get github.com/bobch27/valtra-go
```

## Usage

Here's a complete example showing how to validate a struct using Valtra:

```go
package main

import (
    "fmt"

    "github.com/bobch27/valtra-go"
)

type User struct {
    Name  string
    Email string
    Age   int
}

func (u User) Validate() []error {
    errs := []error{}

    nameRes := valtra.Val(u.Name).Validate(valtra.Required[string](), valtra.MinLengthString[string](3))
    errs = append(errs, nameRes.Errors()...)
    
    emailRes := valtra.Val(u.Email).Validate(valtra.Required[string](), valtra.Email())
    errs = append(errs, emailRes.Errors()...)

    ageRes := valtra.Val(u.Age).Validate(valtra.Min[int](18))
    errs = append(errs, ageRes.Errors()...)

    return errs
}

func main() {
    user := User{
        Name: "Bobby",
        Email: "hello@bobbydonev.com",
        Age: 28,
    }

    errs := user.Validate()
    if len(errs) > 0 {
        log.Fatalln(errs[0])
    }
}
```

## Performance

Valtra is designed for performance:

- **No reflection**: All type checking happens at compile time
- **Zero allocations** (except for error messages when validation fails)
- **Direct comparisons**: No indirection or type assertions in hot paths

## Design Philosophy

1. **Type safety over convenience**: Catch errors at compile time, not runtime
2. **Composition over configuration**: Build complex validators from simple functions
3. **Explicitness over magic**: No reflection, no struct tags, no hidden behavior
4. **Performance matters**: Zero-cost abstractions where possible

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Inspiration

Valtra was built to provide a type-safe validation experience in Go, drawing inspiration from:

- Rust's type system and traits  
- OCaml's functional approach to validation  
- Go's philosophy of simplicity and explicitness

---

**Note**: Valtra requires Go 1.18 or later for generics support.