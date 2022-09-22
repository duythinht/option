# Proof of Concept generic Option type in go

`disclaimer`: this is just a PoC, don't use it on production, I do not recommend to use none idomatic package.


### How to

```
go get github.com/duythinht/option
```

```go

five := option.Some(5)
another := option.None[float64]()

// check with if/else
if five.IsSome() {
    fmt.Printf("Some is `%d`\n", five.Unwrap())
}

if another.IsNone() {
    fmt.Printf("another variable is a None!\n")
}

// switch type

switch s := five.(type) {
case SomeOf[int]:
    fmt.Printf("Some is `%d`\n", five.Unwrap())
case NoneOf[int]:
    fmt.Printf("None of int\n")
default:
    fmt.Printf("unknown type")
}

```
