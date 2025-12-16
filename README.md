# Valtra

A type-safe, performant **val**idation and **tra**nsformation library for Go that uses generics and functional composition. No reflection, no struct tags - just clean, declarative logic! ⚡

[![Go Report Card](https://goreportcard.com/badge/github.com/bobch27/valtra-go)](https://goreportcard.com/report/github.com/bobch27/valtra-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/bobch27/valtra-go.svg)](https://pkg.go.dev/github.com/bobch27/valtra-go)
[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Why Yet Another Validation Package?

Ever spent the night chasing a bug caused by a typo in a struct tag? 😴 I know I have... 🤦

[go-playground/validator](https://github.com/go-playground/validator) is great. Really. In fact, I even built a [Firestore ODM](https://github.com/bobch27/firevault_go) with a validation engine that works similarly (though differently under the hood), and it matches validator's performance almost exactly. 

But reflection and string-based struct tags still lead to **runtime errors** when they should be **compile-time errors**. 🤷 That’s just not the Go way.

Enter **Valtra**. It’s barely even a package - just clever use of generics that let you validate (and transform) data **declaratively and safely**. No reflection, no string parsing, just functions, types, and the compiler doing its job. 🔥

Oh, and it's **~3x faster** than validator (or about **~40x** on cold starts). ⚡ Plus, it lets you shape your data, not just check it... 🧩

## Features

- **🔒 Type-safe**: Your IDE will yell at you before the compiler does. The compiler will yell at you before your users do.
- **⚡ Zero reflection**: Because it's 2025 and we have generics now.
- **🎯 Declarative**: Reads like English, works like magic (except it's not magic, it's just good design).
- **🔧 Composable**: Build complex validations and transformations from simple functions, like LEGO but for paranoid backend developers.
- **📦 Minimal**: Small enough to read in one sitting, powerful enough to actually use.
- **🚀 Fast**: 3x faster than the industry standard. Your users won't notice, but your [benchmarks](#performance) will look great.

## Installation

```bash
go get github.com/bobch27/valtra-go
```

## Usage

Here's a complete example showing how to validate and transform a struct using Valtra:

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

func NewUser(name string, email string, age int) (User, error) {
    c := valtra.NewCollector()

    // Value names (e.g. "email") and custom error messages (e.g. "Name is required") are optional.
    // Names are only used in default error messages.
    user := User{
        Name: valtra.Val(name).
            Transform(valtra.TrimSpace()).
            Validate(valtra.Required[string]("Name is required"), valtra.MinLengthString(3)).
            Collect(c),
        Email: valtra.Val(email, "email").
            Transform(valtra.TrimSpace(), valtra.Lowercase()).
            Validate(valtra.Required[string](), valtra.Email()).
            Collect(c),
        Age: valtra.Val(age).
            Validate(valtra.Min(18, "Age must be 18 or over")).
            Collect(c),
    }

    // check if there are any errors
    if !c.IsValid() {
        // return only first error
        return User{}, c.Errors()[0]
    }

    return user, nil
}

func main() {
    user, err := NewUser("Bobby", "hello@bobbydonev.com", 28)
    if err != nil {
        log.Fatalln("failed to initiate user: %w", err)
    }

    fmt.Println("Success!")
}
```

## Performance

Valtra is designed for compile-time safety, but as a side effect, it’s incredibly fast. Here’s how it compares to popular validation libraries:

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8200Y CPU @ 1.30GHz

BenchmarkValidator-4              335829              4008 ns/op             227 B/op          6 allocs/op
BenchmarkValidatorNoCache-4        18385             54536 ns/op           18878 B/op        288 allocs/op
BenchmarkOzzoValidation-4          41101             32812 ns/op            6678 B/op         81 allocs/op
BenchmarkGoValidator-4            465760              3270 ns/op             390 B/op         22 allocs/op
BenchmarkValtra-4                 922815              1283 ns/op               0 B/op          0 allocs/op
```

**Valtra is ~3x faster than the next fastest competitor with zero allocations.**

[→ View full benchmark code and run it yourself](https://gist.github.com/bobch27/9fee6efd472b1d364f07bdbf975c719e)

### Cold Start Performance

In serverless environments (AWS Lambda, Cloud Functions) or short-lived processes (CLI tools, scripts), the difference is even more dramatic:

```
BenchmarkValidator (warm cache)                  ~3,800 ns/op
BenchmarkValidatorNoCache (cold cache)           ~54,000 ns/op  ⚠️ 14x slower
BenchmarkValtra (no cache needed)                ~1,300 ns/op
```

**Valtra is ~38x faster on cold starts** because there's no reflection cache to build. The "cache" is the compiled binary itself.

This makes Valtra particularly well-suited for:
- 🚀 Serverless functions (Lambda, Cloud Run, etc.)
- 🛠️ CLI tools and scripts
- 🔄 Microservices with frequent restarts
- 📦 Short-lived containers

### How Is It So Fast?

- **No reflection**: All type checking happens at compile time
- **Zero allocations** (except for error messages when validation/transformation fails)
- **Direct comparisons**: No indirection or type assertions in hot paths

## Testing

Valtra has **100% test coverage** with focused unit tests for each validation, transformation and the Collector. 

Run tests locally:
```bash
go test -v
go test -cover  # Shows 100% coverage
```


## Design Philosophy

1. **Type safety over convenience**: Catch errors at compile time, not runtime
2. **Composition over configuration**: Build complex validations and transformations from simple functions
3. **Explicitness over magic**: No reflection, no struct tags, no hidden behavior
4. **Performance matters**: Zero-cost abstractions where possible

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Inspiration

Valtra was built to provide a type-safe validation and transformation experience in Go, drawing inspiration from:

- Rust's type system and traits  
- OCaml's functional approach to validation  
- Go's philosophy of simplicity and explicitness

---

**Note**: Valtra requires Go 1.18 or later for generics support.