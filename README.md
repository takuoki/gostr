# gostr

Simple library for Go to convert any data to string.

```go
type hoge struct {
  foo stringmark
  bar int
}

v := struct {
  sl []*hoge
  m  map[string]*hoge
}{
  sl: []*hoge{{foo: "a", bar: 1}, {foo: "b", bar: 2}},
  m:  map[string]*hoge{"x": {foo: "y", bar: 9}},
}
```

**Too Bad!**

```go
fmt.Printf("%+v\n", v)
// {sl:[0xc04203ede0 0xc04203ee00] m:map[x:0xc04203ee20]}
```

**So Good!**

```go
fmt.Println(gostr.Stringify(v))
// {sl:[*hoge{foo:"a", bar:1}, *hoge{foo:"b", bar:2}], m:map["x":*hoge{foo:"y", bar:9}]}
```

## Usage

```go
str := gostr.Stringify(v)
```

**Caution!**

* This library doesn't support the following types.

  * func
  * chan
  * complex64, complex128
  * UnsafePointer

* As a map key, this library supports only the following types.

  * int, int8, int16, int32, int64
  * uint, uint8, uint16, uint32, uint64
  * string

## Installation

```txt
$ go get github.com/takuoki/gostr
```

## License

MIT