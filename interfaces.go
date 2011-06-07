package sexp

type Blitter interface {
	BlockCopy(destination, source, count int)
	BlockClear(start, count int)
}

type Equatable interface {
	Equal(interface{}) bool
}

type Indexable interface {
	At(index int) interface{}
	Set(index int, value interface{})
}

type Expandable interface {
	Expand(i, n int)
}

type Appendable interface {
	Append(interface{})
}

type Linear interface {
	Len() int
}

type FixedSize interface {
	Linear
	Cap()
}

type Resizeable interface {
	FixedSize
	Reallocate(n int)
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

type Linkable interface {
	Linear
	Start() ListNode
	End() ListNode
}