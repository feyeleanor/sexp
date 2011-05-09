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