package sexp

import "testing"

func TestNodeIsNil(t *testing.T) {
	ConfirmIsNil := func(n *Node, r bool) {
		if n.IsNil() != r {
			t.Fatalf("%v.IsNil() should be %v", n, r)
		}
	}
	n := &Node{}
	ConfirmIsNil(n, true)

	n.Head = 0
	ConfirmIsNil(n, false)

	n.Tail = n
	ConfirmIsNil(n, false)

	ConfirmIsNil(&Node{ Head: n, Tail: n }, false)
}

func TestNodeNotNil(t *testing.T) {
	ConfirmNotNil := func(n *Node, r bool) {
		if n.NotNil() != r {
			t.Fatalf("%v.NotNil() should be %v", n, r)
		}
	}
	n := &Node{}
	ConfirmNotNil(n, false)

	n.Head = 0
	ConfirmNotNil(n, true)

	n.Tail = n
	ConfirmNotNil(n, true)

	ConfirmNotNil(&Node{ Head: n, Tail: n }, true)
}

func TestEnd(t *testing.T) {
	ConfirmEnd := func(n *Node, r interface{}) {
		if x := n.End(); x.Head != r {
			t.Fatalf("%v.End() should be '%v' but is '%v'", n, r, x.Head)
		}
	}
	ConfirmEnd(List(0).start.End(), 0)
	ConfirmEnd(List(0, 1).start.End(), 1)
	ConfirmEnd(List(0, 1, 2).start.End(), 2)
}

func TestTraverse(t *testing.T) {
	ConfirmTraverse := func(n *Node, i int, r interface{}) {
		if x := n.Traverse(i); x.Head != r {
			t.Fatalf("%v.Traverse(%v) should be '%v' but is '%v'", n, i, r, x.Head)
		}
	}
	RefuteTraverse := func(n *Node, i int) {
		if x := n.Traverse(i); x != nil {
			t.Fatalf("%v.Traverse(%v) should be nil but is '%v'", n, i, x.Head)
		}
	}
	l := List(1, 2, 3, 4, 5)
	ConfirmTraverse(l.start, 0, 1)
	ConfirmTraverse(l.start, 1, 2)
	ConfirmTraverse(l.start, 2, 3)
	ConfirmTraverse(l.start, 3, 4)
	ConfirmTraverse(l.start, 4, 5)
	RefuteTraverse(l.start, -1)
	RefuteTraverse(l.start, 5)
}