package sexp

import "fmt"
import "strings"

/*
	A LinearList is a finitely-terminated list structure.
	Each node in the list may point to exactly one other node in the list.
	The terminating node does not point to any other node.
	No node may be pointed to by more than one other node in the list.
	There are no nil links between nodes in the list.
*/

func List(items... interface{}) (l *LinearList) {
	l = new(LinearList)
	l.AppendSlice(items)
	return
}

type LinearList struct {
	start	*Node
	end		*Node
	length	int
}

func (l LinearList) IsNil() bool {
	return l.start == nil || l.end == nil || l.length == 0
}

func (l LinearList) NotNil() bool {
	return l.start != nil && l.end != nil && l.length != 0
}

func (l LinearList) Len() (c int) {
	if l.NotNil() {
		c = l.length
	}
	return
}

func (l LinearList) Each(f func(interface{})) {
	if l.NotNil() {
		for n := l.start; n != nil; n = n.Tail {
			f(n.Head)
		}
	}
}

func (l LinearList) At(i int) (r interface{}) {
	if l.NotNil() {
		var n	*Node
		for n = l.start; i > 0; i-- {
			n = n.Tail
		}
		r = n.Head
	}
	return
}

func (l LinearList) Set(i int, v interface{}) {
	if l.NotNil() {
		var n	*Node
		for n = l.start; i > 0; i-- {
			n = n.Tail
		}
		n.Head = v
	}
}

func (l *LinearList) Append(v interface{}) {
	if l.IsNil() {
		l.start = &Node{ Head: v }
		l.end = l.start
		l.length = 1
	} else {
		l.end.Tail = &Node{ Head: v }
		l.end = l.end.Tail
		l.length++
	}
}

func (l *LinearList) AppendSlice(s []interface{}) {
	if len(s) > 0 {
		if l.IsNil() {
			l.Append(s[0])
			s = s[1:]
		}
		for _, v := range s {
			l.end.Tail = &Node{ Head: v }
			l.end = l.end.Tail
		}
		l.length += len(s)
	}
}

func (l LinearList) equal(o LinearList) (r bool) {
	switch {
	case l.IsNil():
		r = o.IsNil()
	case l.Len() == o.Len():
		r = true
		n := l.start
		x := o.start
		for i := 0; r && i < l.Len(); i++ {
			if r = n.Equal(x); r {
				n = n.Tail
				x = x.Tail
			}
		}
	}
	return
}

//	Determines if another object is equivalent to the LinearList
//	Two CycLists are identical if they both have the same number of nodes, and the head of each node is the same
func (l LinearList) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *LinearList:	r = l.equal(*o)
	case LinearList:	r = l.equal(o)
	default:			r = l.start.Equal(o)
	}
	return 
}

func (l LinearList) String() (t string) {
	if l.length > 0 {
		terms := []string{}
		for n := l.start; n != nil; n = n.Tail {
			terms = append(terms, fmt.Sprintf("%v", n.Head))
		}
		t = strings.Join(terms, " ")
		t = strings.Replace(t, "()", "nil", -1)
		t = strings.Replace(t, "<nil>", "nil", -1)
	}
	return "(" + t + ")"
}

func (l LinearList) Depth() (d int) {
	for n := l.start; n != nil; n = n.Tail {
		if v, ok := n.Head.(Nested); ok {
			if r := v.Depth() + 1; r > d {
				d = r
			}
		}
	}
	return
}

//	Reverses the order in which elements of a CycList are traversed
func (l *LinearList) Reverse() {
	if l.NotNil() {
		var next, result		*Node
//		current := &Node{ Head: l.start.Head, Tail: l.start.Tail }
		current := l.start
		l.end = current

		for ; current != nil; {
			next = current.Tail
			current.Tail = result
			result = current
			current = next				
		}
		l.start = result
	}
}

func (l *LinearList) Flatten() {
	if l.NotNil() {
		for n := l.start; n != nil; n = n.Tail {
			switch h := n.Head.(type) {
			case *LinearList:		switch {
									case h.IsNil():			n.Head = nil
									case h.length == 1:		n.Head = h.start.Head
									case n == l.end:		h.Flatten()
															l.end = h.end
															n.Head = h.start.Head
															n.Tail = h.start.Tail
															l.length += h.length - 1
									default:				h.Flatten()
															h.end.Tail = n.Tail
															n.Head = h.start.Head
															n.Tail = h.start.Tail
															l.length += h.length - 1
									}
			case Flattenable:		h.Flatten()
			}
		}
	}
}

func (l *LinearList) Delete(from, to int) {
	if l.NotNil() && from >= 0 && to < l.length && from <= to {
		last_element_index := l.length - 1
		switch {
		case from == 0:
			switch {
			case to == 0:
				l.start = l.start.Tail
				l.length -= 1
			case to == last_element_index:
				l.start = nil
				l.end = nil
				l.length = 0
			default:
				l.start = l.start.Traverse(to + 1)
				l.length -= to + 1
			}

		case from == to:
			s := l.start.Traverse(from - 1)
			e := s.Traverse(1)
			s.Tail = e.Tail
			l.length -= 1

		case from == last_element_index:
			l.end = l.start.Traverse(from - 1)
			l.end.Tail = nil
			l.length -= 1

		case to == last_element_index:
			l.end = l.start.Traverse(from - 1)
			l.end.Tail = nil
			l.length = from

		case from < 0:					fallthrough
		case to > last_element_index:	fallthrough
		case from > to:					//	do nothing

		default:
			e := l.start.Traverse(from - 1)
			e.Tail = e.Traverse(to - from + 2)
			l.length -= to - from + 1
		}
	}
}

func (l *LinearList) Cut(from, to int) (r LinearList) {
	defer func() {
		if recover() != nil {
			r = LinearList{}
		}
	}()

	if to < from {
		panic(to)
	}

	var tail_l, tail_r	*Node

	r.start = l.start
	for r.length = 0; r.length < from; r.length++ {
		tail_l = r.start
		r.start = r.start.Tail
	}

	for tail_r = r.start; r.length < to; r.length++ {
		tail_r = tail_r.Tail
	}
	if from == 0 {
		l.start = tail_r.Tail
	} else {
		tail_l.Tail = tail_r.Tail
	}
	r.length = to - from + 1
	l.length -= r.length
	tail_r.Tail = nil
	r.end = tail_r
	return
}

func (l *LinearList) Insert(i int, o *LinearList) {
	if l.NotNil() {
		var n	*Node
		for n = l.start; i > 0; i-- {
			n = n.Tail
		}
		switch {
		case n == nil:
			l.start = o.start
			l.end = o.end
			l.length = o.length
		case n.Tail == nil:
			n.Head = o.start.Head
			n.Tail = o.start.Tail
			l.end = o.end
			l.length += o.length
		default:
			n.Head = o.start.Head
			n.Tail = o.start.Tail
			l.length += o.length
		}
	}
}

func (l LinearList) Car() (r interface{}) {
	if l.NotNil() {
		r = l.start.Head
	}
	return
}

func (l LinearList) Cdr() (r LinearList) {
	if l.NotNil() {
		r.start = l.start.Tail
		r.end = l.end
		r.length = l.length - 1
	}
	return
}

func (l *LinearList) Rplaca(i interface{}) {
	if l.IsNil() {
		*l = *(List(i))
	} else {
		l.start.Head = i
	}
}

func (l *LinearList) Rplacd(tail *LinearList) {
	if l.IsNil() {
		l.start = tail.start
		l.end = tail.end
		l.length = tail.length
	} else {
		l.start.Tail = tail.start
		l.end = tail.end
		l.length = tail.length + 1
	}
}