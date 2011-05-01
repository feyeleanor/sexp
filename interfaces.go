package sexp

import "fmt"

type CyclicNested interface {
	depth(memo) int
}

type Nested interface {
	Depth() int
}

type Addressable interface {
	Addr() uintptr
}

func printAddress(a Addressable) string {
	return fmt.Sprintf("[%v]", a.Addr())
}
