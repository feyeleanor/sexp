package sexp

import "fmt"

type Node struct {
	Head		interface{}
	Tail		*Node
}

func (n *Node) End() (r *Node) {
	if n.NotNil() {
		for r = n; r.Tail != nil; r = r.Tail {}
	}
	return
}

func (n *Node) Traverse(i int) (r *Node) {
	if i >= 0 {
		for r = n; i > 0 && r != nil; i-- {
			r = r.Tail
		}
	}
	return
}

func (n *Node) Append(x interface{}) {
	n.Tail = &Node{ Head: x }
}

func (n *Node) Prepend(x interface{}) {
	*n = Node{ Head: x, Tail: n }
}

func (n Node) IsNil() bool {
	return n.Head == nil && n.Tail == nil
}

func (n Node) NotNil() bool {
	return n.Head != nil || n.Tail != nil
}

func (n Node) equal(o Node) (r bool) {
	if v, ok := n.Head.(Equatable); ok {
		r = v.Equal(o.Head)
	} else {
		r = n.Head == o.Head
	}
	return
}

func (n *Node) Equal(o interface{}) (r bool) {
	if n != nil {
		switch o := o.(type) {
		case *Node:			r = n.equal(*o)
		case Node:			r = n.equal(o)
		default:			r = n.equal(Node{ Head: o })
		}
	}
	return
}

func (n *Node) String() (t string) {
	if n.IsNil() {
		t = "nil"
	} else {
	 	t = fmt.Sprint(n.Head)
	}
	return
}

func (n *Node) Car() interface{} {
	return n.Head
}

func (n *Node) Cdr() *Node {
	return n.Tail
}

func (n *Node) Rplaca(i interface{}) {
	n.Head = i
}

func (n *Node) Rplacd(next *Node) {
	n.Tail = next
}