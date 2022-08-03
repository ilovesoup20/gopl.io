package main

import (
	"bytes"
	"io"
	"os"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type ReadWriter interface {
	Reader
	Writer
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

type ReadWriter2 interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

type ReadWriter3 interface {
	Read(p []byte) (n int, err error)
	Writer
}

func main() {
	var w io.Writer
	w = os.Stdout         // OK: *os.File has Write method
	w = new(bytes.Buffer) // OK: *bytes.Buffer has Write method
	// w = time.Second			// compile error

	var rwc io.ReadWriteCloser
	rwc = os.Stdout // OK: *os.File has Read, Write, Close methods
	// rwc = new(bytes.Buffer) // compile error

	w = rwc // OK: io.ReadWriteCloser has Write method
	// rwc = w				// compile error

}
