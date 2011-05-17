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