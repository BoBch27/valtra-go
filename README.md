# Valtra

A type-safe, performant validation library for Go that uses generics and functional composition. No reflection, no struct tags - just clean, declarative validation! ⚡

## Why Yet Another Validation Package?

Ever spent the night chasing a bug caused by a typo in a struct tag? 😴 I know I have... 🤦

[go-playground/validator](https://github.com/go-playground/validator) is great. Really. In fact, I even built a [Firestore ODM](https://github.com/bobch27/firevault_go) with a validation engine that works similarly (though differently under the hood), and it matches validator's performance almost exactly. 

But reflection and string-based struct tags still lead to **runtime errors** when they should be **compile-time errors**. 🤷 That’s just not the Go way.

Enter **Valtra**. It’s hardly even a package - just clever use of generics that let you validate data **declaratively and safely**. No reflection, no string parsing, just functions, types, and the compiler doing its job. ⚡