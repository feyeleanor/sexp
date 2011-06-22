package sexp

import "reflect"

type Typed interface {
	Type()	reflect.Type
}

type Blitter interface {
	BlockCopy(destination, source, count int)
	BlockClear(start, count int)
}

type Equatable interface {
	Equal(interface{}) bool
}

type Linear interface {
	Len() int
}

type IndexedReader interface {
	At(index int) interface{}
}

type IndexedWriter interface {
	Set(index int, value interface{})
	Clear(index int)
}

type Indexable interface {
	Linear
	IndexedReader
	IndexedWriter
}

type Reversible interface {
	Reverse()
}

type Repeatable interface {
	Repeat(count int)
}

type Expandable interface {
	Indexable
	Expand(i, n int)
}

type Appendable interface {
	Append(interface{})
}

type FixedSize interface {
	Linear
	Cap() int
}

type Resizeable interface {
	FixedSize
	Reallocate(l, c int)
}

type Nested interface {
	Depth() int
}

type Flattenable interface {
	Flatten()
}

type Iterable interface {
	Each(func(interface{}))
}

type Transformable interface {
	Transform(func(interface{}) interface{})
}

type Collectable interface {
	Collect(func(interface{}) interface{}) interface{}
}

type Combinable interface {
	Combine(interface{}, func(interface{}, interface{}) interface{}) interface{}
}

type Linkable interface {
	Linear
	Start() ListNode
	End() ListNode
}

type Feeder interface {
	Feed(chan interface{}, func(interface{}) interface{})
}

type Piper interface {
	Pipe(func(interface{}) interface{}) chan interface{}
}