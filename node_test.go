package sexp

import "testing"


func TestLastElement(t *testing.T) {
	ConfirmLastElement := func(n ListNode, r interface{}) {
		x := LastElement(n)
		switch x := x.(type) {
		case Equatable: 	if !x.Equal(r) {
								t.Fatalf("Last(%v) Equatable: should be '%v' but is '%v'", n, r, x)
							}
		default:		 	if r != x {
								t.Fatalf("Last(%v) default: should be '%v' but is '%v'", n, r, x)
							}
		}
	}
	ConfirmLastElement(List(0).start, 0)
	ConfirmLastElement(List(0, 1).start, 1)
	ConfirmLastElement(List(0, 1, 2).start, 2)
}

func TestMoveToNode(t *testing.T) {
	ConfirmMoveTo := func(n ListNode, i int, r interface{}) {
		switch x := n.MoveTo(i); {
		case x.Content() != r:	t.Fatalf("%v.MoveTo(%v) should be '%v' but is '%v'", n, i, r, x.Content())
		}
	}
	RefuteMoveTo := func(n ListNode, i int) {
		if x := n.MoveTo(i); x != nil {
			t.Fatalf("%v.MoveTo(%v) should not succeed", n, i)
		}
	}
	l := List(1, 2, 3, 4, 5)
	ConfirmMoveTo(l.start, 0, 1)
	ConfirmMoveTo(l.start, 1, 2)
	ConfirmMoveTo(l.start, 2, 3)
	ConfirmMoveTo(l.start, 3, 4)
	ConfirmMoveTo(l.start, 4, 5)
	RefuteMoveTo(l.start, -1)
	RefuteMoveTo(l.start, 5)
}