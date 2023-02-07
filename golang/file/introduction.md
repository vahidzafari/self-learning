# The io.Reader and io.Writer interfaces

io.Reader and io.Writer interfaces are the basis of file I/O in Go—the former allows you to read from a file whereas the latter allows you to write to a file.

This definition, which should be revisited when we want one of our data types to satisfy the io.Reader interface, tells us the following:

- The Reader interface requires the implementation of a single method
- The parameter of Read() is a byte slice
- The return values of Read() are an integer and an error

The Read() method takes a byte slice as input, which is going to be filled with data
up to its length, and returns the number of bytes read as well as an error variable.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

The previous definition, which should be revisited when we want one of our data types to satisfy the io.Writer interface and to write to a file, reveals the following information:

- The interface requires the implementation of a single method
- The parameter of Write() is a byte slice
- The return values of Write() are an integer and an error value

The Write() method takes a byte slice, which contains the data that you want to
write, as input and returns the number of bytes written and an error variable.

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```
