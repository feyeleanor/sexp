package sexp

type ListHeader struct {
	start 	*Node
	end		*Node
	length	int
}

func (l *ListHeader) Clear() {
	l.start = nil
	l.end = nil
	l.length = 0
}

func (l ListHeader) IsNil() bool {
	return l.start == nil && l.end == nil && l.length == 0
}

func (l ListHeader) NotNil() bool {
	return l.start != nil || l.end != nil || l.length > 0
}

func (l ListHeader) Len() (c int) {
	if l.NotNil() {
		c = l.length
	}
	return
}

func (l ListHeader) Each(f func(interface{})) {
	if l.NotNil() {
		for n := l.start; n != nil; n = n.Tail {
			f(n.Head)
		}
	}
}

func (l ListHeader) At(i int) (r interface{}) {
	if l.NotNil() {
		if n := l.start.MoveTo(i); n != nil {
			r = n.Head
		}
	}
	return
}

func (l ListHeader) Set(i int, v interface{}) {
	if l.NotNil() {
		if n := l.start.MoveTo(i); n != nil {
			n.Head = v
		}
	}
}