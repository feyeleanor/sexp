package sexp

import "testing"

func TestNodeIsNil(t *testing.T) {
	n := &Node{}
	if !n.IsNil() { t.Fatalf("%v.IsNil() should be true", n) }

	n.Head = 0
	if n.IsNil() { t.Fatalf("%v.IsNil() should be false", n) }

	n.Tail = n
	if n.IsNil() { t.Fatalf("%v.IsNil() should be false", n) }

	n = &Node{ Head: n, Tail: n }
	if n.IsNil() { t.Fatalf("%v.IsNil() should be false", n) }
}
