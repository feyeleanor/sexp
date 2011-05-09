package sexp

type Nested interface {
	Depth() int
}

type Flattenable interface {
	Flatten()
}

type Equatable interface {
	Equal(interface{}) bool
}

type Indexable interface {
	At(int) interface{}
	Set(int, interface{})
}

type Value interface {
	IsNil() bool
	NotNil() bool
	String() string
}

type Iterable interface {
	Each(func(interface{}))
}

type SEXP interface {
	Nested
	Flattenable
	Equatable
	Indexable
	Value
}