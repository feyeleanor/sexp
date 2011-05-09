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