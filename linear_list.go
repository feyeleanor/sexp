package sexp

import "fmt"
import "reflect"

func List(items... interface{}) (l *LinearList) {
	l = &LinearList{ &Node{} }
	n := l.Node
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
	return
}

type LinearList struct {
	*Node
}

func (l *LinearList) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *LinearList:	r = reflect.DeepEqual(l, o)
	case LinearList:	r = reflect.DeepEqual(*l, o)
	default:			r = l.Node.Equal(o)
	}
	return 
}

func (l LinearList) Each(f func(interface{})) {
	for n := l.Node; n != nil; n = n.Tail {
		f(n.Head)
	}
}

func (l *LinearList) String() (t string) {
	if !l.IsNil() {
		for n := l.Node; n != nil; n = n.Tail {
			if len(t) > 0 {
				t += " "
			}
			switch h := n.Head.(type) {
			case nil:				t += "nil"
			case *LinearList:		if term := h.String(); term == "()" {
										t += "nil"
									} else {
										t += term
									}
			case *Node:				if term := h.String(); term == "()" {
										t += "nil"
									} else {
										t += term
									}
			case Addressable:		t += printAddress(h)
			default:				t += fmt.Sprintf("%v", h)
			}
		}
	}
	return "(" + t + ")"
}

func (l LinearList) Len() (i int) {
	if !l.IsNil() {
		for n := l.Node; n != nil; n = n.Tail { i++ }
	}
	return
}

func (l LinearList) Depth() (d int) {
	for n := l.Node; n != nil; n = n.Tail {
		if v, ok := n.Head.(Nested); ok {
			if r := v.Depth() + 1; r > d {
				d = r
			}
		}
	}
	return
}

func (l *LinearList) Reverse() {
	if l.Node.Tail != nil {
		var n	*Node
		current := &Node{ Head: l.Node.Head, Tail: l.Node.Tail}
		for ; current != nil; {
			next := current.Tail
			current.Tail = n
			n = current
			current = next				
		}
		*l = LinearList{ Node: n }
	}
}

func (l LinearList) Flatten() {
	
}

func (l LinearList) At(i int) (n interface{}) {
	var c	*Node
	for c = l.Node; i > 0 && c.Tail != nil; i-- {
		c = c.Tail
	}
	if i == 0 {
		n = c.Head
	}
	return
}

func (l LinearList) Set(i int, n interface{}) {
	var c	*Node
	for c = l.Node; i > 0 && c.Tail != nil; i-- {
		c = c.Tail
	}
	if c != nil {
		c.Head = n
	}
}

func (l LinearList) Cut(from, to int) (ok bool) {
	if from < to {
		var target, source	*Node

		for n, i := l.Node, 0; n != nil; n = n.Tail {
			if i == to		{ target = n }
			if i == from	{ source = n }
			i++
		}
		if source != nil && target != nil {
			source.Tail = target
			ok = true
		}
	}
	return
}

func (l LinearList) End() (n *Node) {
	for n = l.Node; n.Tail != nil; n = n.Tail {}
	return
}