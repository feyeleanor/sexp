package sexp

import "fmt"

type Nested interface {
	Depth() int
}

type Flattenable interface {
	Flatten()
}

type Equatable interface {
	Equal(interface{}) bool
}

type Addressable interface {
	Addr() uintptr
}

func printAddress(a Addressable) string {
	return fmt.Sprintf("[%v]", a.Addr())
}
