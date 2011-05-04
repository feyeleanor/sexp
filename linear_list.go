package sexp

import "fmt"
import "strings"

func List(items... interface{}) (l LinearList) {
	n := &Node{}
	l.node = n
	tails := len(items) - 1
	for i, v := range items {
		if v == nil {
			n.Head = &Node{}
		} else {
			n.Head = v
		}
		if i < tails {
			n.Tail = &Node{}
		}
		n = n.Tail
	}
	l.length = len(items)
	return
}

type LinearList struct {
	node	*Node
	length	int
}

func (l LinearList) IsNil() bool {
	return l.node == nil || l.length == 0
}

func (l LinearList) Len() int {
	return l.length
}

func (l LinearList) Each(f func(interface{})) {
	for n := l.node; n != nil; n = n.Tail {
		f(n.Head)
	}
}

func (l LinearList) equal(o LinearList) (r bool) {
	switch {
	case l.IsNil():
		r = o.IsNil()
	case l.Len() == o.Len():
		r = true
		n := l.node
		x := o.node
		for i := 0; r && i < l.Len(); i++ {
			if r = n.Equal(x); r {
				n = n.Tail
				x = x.Tail
			}
		}
	}
	return
}

func (l LinearList) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *LinearList:	r = l.equal(*o)
	case LinearList:	r = l.equal(o)
	default:			r = l.node.Equal(o)
	}
	return 
}

func (l LinearList) String() (t string) {
	if l.length > 0 {
		terms := []string{}
		for n := l.node; n != nil; n = n.Tail {
			terms = append(terms, fmt.Sprintf("%v", n.Head))
		}
		t = strings.Join(terms, " ")
		t = strings.Replace(t, "()", "nil", -1)
		t = strings.Replace(t, "<nil>", "nil", -1)
	}
	return "(" + t + ")"
}

func (l LinearList) Depth() (d int) {
	for n := l.node; n != nil; n = n.Tail {
		if v, ok := n.Head.(Nested); ok {
			if r := v.Depth() + 1; r > d {
				d = r
			}
		}
	}
	return
}

func (l LinearList) Reverse() {
	if !l.IsNil() {
		var n, next		*Node
		current := &Node{ Head: l.node.Head, Tail: l.node.Tail }
		for ; current != nil; {
			next = current.Tail
			current.Tail = n
			n = current
			current = next				
		}
		(*l.node) = *n
	}
}

func (l *LinearList) Flatten() {
	if l.length > 0 {
		for n := l.node; n != nil; n = n.Tail {
			switch h := n.Head.(type) {
			case LinearList:	h.Flatten()
								n.Head = h.node.Head
								t := h.node.End()
								t.Tail = n.Tail
								n.Tail = h.node.Tail
								l.length += h.length - 1
			case Flattenable:	h.Flatten()
			}
		}
	}
}

func (l LinearList) At(i int) (n interface{}) {
	if i < l.length {
		var c	*Node
		for c = l.node; i > 0 && c.Tail != nil; i-- {
			c = c.Tail
		}
		if i == 0 {
			n = c.Head
		}
	}
	return
}

func (l LinearList) Set(i int, v interface{}) {
	if i < l.length {
		var c	*Node
		for c = l.node; i > 0 && c.Tail != nil; i-- {
			c = c.Tail
		}
		if c != nil {
			c.Head = v
		}
	}
}

func (l *LinearList) Append(v interface{}) {
	if l.IsNil() {
		l.node = &Node{ Head: v }
	} else {
		e := l.node.End()
		e.Append(v)
	}
	l.length++
}

func (l *LinearList) Delete(from, to int) {
	if from >= 0 && from <= to && to < l.length {
		var r				LinearList
		var tail_l, tail_r	*Node

		r.node = l.node
		for r.length = 0; r.length < from; r.length++ {
			tail_l = r.node
			r.node = r.node.Tail
		}

		for tail_r = r.node; r.length < to; r.length++ {
			tail_r = tail_r.Tail
		}
		if from == 0 {
			l.node = tail_r.Tail
		} else {
			tail_l.Tail = tail_r.Tail
		}
		l.length -= to - from + 1
		tail_r.Tail = nil
	}
}

func (l *LinearList) Cut(from, to int) (r LinearList) {
	defer func() {
		if recover() != nil {
			r.node = nil
			r.length = 0
		}
	}()

	if to < from {
		panic(to)
	}

	var tail_l, tail_r	*Node

	r.node = l.node
	for r.length = 0; r.length < from; r.length++ {
		tail_l = r.node
		r.node = r.node.Tail
	}

	for tail_r = r.node; r.length < to; r.length++ {
		tail_r = tail_r.Tail
	}
	if from == 0 {
		l.node = tail_r.Tail
	} else {
		tail_l.Tail = tail_r.Tail
	}
	r.length = to - from + 1
	l.length -= r.length
	tail_r.Tail = nil
	return
}

func (l LinearList) Insert(i int, n LinearList) {
	
}

func (l LinearList) Car() (r interface{}) {
	if l.node != nil {
		r = l.node.Head
	}
	return
}

func (l LinearList) Cdr() (r LinearList) {
	if l.length > 0 {
		r.node = l.node.Tail
		r.length = l.length - 1
	}
	return
}

func (l *LinearList) Rplaca(i interface{}) {
	if l.length == 0 {
		l.node = &Node{ Head: i }
		l.length = 1
	} else {
		l.node.Head = i
	}
}

func (l *LinearList) Rplacd(tail LinearList) {
	if l.length > 0 {
		l.node.Tail = tail.node
		l.length = tail.length + 1
	} else {
		l.node = tail.node
		l.length = tail.length
	}
}