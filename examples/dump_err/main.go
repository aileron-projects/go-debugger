package main

import (
	"io"

	"github.com/aileron-projects/go-debugger"
)

type MyError struct {
	inner error
	typ   string
	msg   string
}

func (e *MyError) Error() string {
	return e.typ + ": " + e.msg
}

func main() {
	e1 := &MyError{
		inner: io.ErrUnexpectedEOF,
		typ:   "HTTP_ERROR",
		msg:   "request failed",
	}
	e2 := &MyError{
		inner: io.ErrUnexpectedEOF,
		typ:   "VALIDATION_ERROR",
		msg:   "parameter must be number",
	}

	// Run with the tag.
	// go run -tags dumperr ./main.go
	debugger.DumpErr("dump error", e1)

	// Tag is not required.
	debugger.DumpErrAlways("dump error always", e2)
}
