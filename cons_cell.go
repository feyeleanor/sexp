package sexp

import "fmt"

type ConsCell struct {
	Head		interface{}
	Tail		*ConsCell
}

func (n *ConsCell) End() (r *ConsCell) {
	if n.NotNil() {
		for r = n; r.Tail != nil; r = r.Tail {}
	}
	return
}

func (n *ConsCell) MoveTo(i int) (r *ConsCell) {
	if i < 0 {
		panic(i)
	}
	for r = n; i > 0 && r != nil; i-- {
		r = r.Tail
	}
	if i != 0 || r == nil {
		panic(i)
	}
	return
}

func (n *ConsCell) Append(x interface{}) {
	n.Tail = &ConsCell{ Head: x }
}

func (n *ConsCell) Prepend(x interface{}) {
	*n = ConsCell{ Head: x, Tail: n }
}

func (n ConsCell) IsNil() bool {
	return n.Head == nil && n.Tail == nil
}

func (n ConsCell) NotNil() bool {
	return n.Head != nil || n.Tail != nil
}

func (n ConsCell) equal(o ConsCell) (r bool) {
	if v, ok := n.Head.(Equatable); ok {
		r = v.Equal(o.Head)
	} else {
		r = n.Head == o.Head
	}
	return
}

func (n *ConsCell) Equal(o interface{}) (r bool) {
	if n != nil {
		switch o := o.(type) {
		case *ConsCell:			r = n.equal(*o)
		case ConsCell:			r = n.equal(o)
		default:			r = n.equal(ConsCell{ Head: o })
		}
	}
	return
}

func (n *ConsCell) String() (t string) {
	if n.IsNil() {
		t = "nil"
	} else {
	 	t = fmt.Sprint(n.Head)
	}
	return
}

func (n *ConsCell) Car() interface{} {
	return n.Head
}

func (n *ConsCell) Cdr() *ConsCell {
	return n.Tail
}

func (n *ConsCell) Rplaca(i interface{}) {
	n.Head = i
}

func (n *ConsCell) Rplacd(next *ConsCell) {
	n.Tail = next
}