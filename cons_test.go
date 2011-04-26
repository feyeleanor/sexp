package sexp

import "fmt"
import "testing"

func TestConsCellIsNil(t *testing.T) {
	c := ConsCell{}
	if !c.IsNil() { t.Fatalf("%v.IsNil() should be true", c) }

	c = ConsCell{ Head: 0 }
	if c.IsNil() { t.Fatalf("%v.IsNil() should be false", c) }

	c = ConsCell{ Tail: &c }
	if c.IsNil() { t.Fatalf("%v.IsNil() should be false", c) }

	c = ConsCell{ Head: c, Tail: &c }
	if c.IsNil() { t.Fatalf("%v.IsNil() should be false", c) }
}

func TestConsCellString(t *testing.T) {
	ConfirmFormat := func(c *ConsCell, x string) {
		if s := c.String(); s != x {
			t.Fatalf("'%v' erroneously serialised as '%v'", x, s)
		}
	}

	ConfirmFormat(Cons(nil, nil), "()")
	ConfirmFormat(ConsNil(), "()")
	ConfirmFormat(Cons(0, nil), "(0)")
	ConfirmFormat(Cons(0, ConsNil()), "(0 nil)")
	ConfirmFormat(Cons(1, Cons(0, ConsNil())), "(1 (0 nil))")

	ConfirmFormat(List(1, 0, ConsNil()), "(1 0 nil)")

	c := List(0)
	c.Tail = c
	ConfirmFormat(c, "(0 ...)")

	c = List(0, c)
	ConfirmFormat(c, "(0 (0 ...))")

	c.Tail.Head = c
	ConfirmFormat(c, "(0 (...))")

	c.Tail.Tail = c
	ConfirmFormat(c, "(0 (...) ...)")

	r := List(10, c)
	ConfirmFormat(r, "(10 (0 (...) ...))")
	ConfirmFormat(r.Tail, "((0 (...) ...))")
	ConfirmFormat(r.Tail.Head.(*ConsCell), "(0 (...) ...)")

	c.Tail = List(-1, -2, c)
	ConfirmFormat(r, "(10 (0 -1 -2 (...)))")

	c.Tail = List(3, 4, c)
	ConfirmFormat(r, "(10 (0 3 4 (...)))")

	c.Tail.Tail = List(3)
	ConfirmFormat(r, "(10 (0 3 3))")

t.Fatal()
	c.Tail.Tail = List(c, List(-1, -2, c))
	ConfirmFormat(c, "(0 -1 -2 (...) (-1 -2 ...))")

t.Fatal()

	ConfirmFormat(r, fmt.Sprintf("(10 (1 0 cons(%v) -1 -2 ...))", printAddress(c)))

	r = Cons(10, Cons(9, c))
	ConfirmFormat(r, fmt.Sprintf("(10 9 (1 0 cons(%v) -1 -2 ...))", printAddress(c)))

	r = Cons(10, Cons(9, Cons(8, c)))
	ConfirmFormat(r, fmt.Sprintf("(10 9 8 (1 0 cons(%v) -1 -2 ...))", printAddress(c)))

	r = Cons(10, Cons(9, Cons(8, Cons(7, c))))	
	ConfirmFormat(r, fmt.Sprintf("(10 9 8 7 (1 0 cons(%v) -1 -2 ...))", printAddress(c)))
}

func TestConsCellList(t *testing.T) {
	ConfirmFormat := func(c *ConsCell, x string) {
		if s := c.String(); s != x {
			t.Fatalf("'%v' erroneously serialised as '%v'", x, s)
		}
	}
	ConfirmFormat(List(), "()")
	ConfirmFormat(List(1), "(1)")
	ConfirmFormat(List(2, 1), "(2 1)")
	ConfirmFormat(List(3, 2, 1), "(3 2 1)")
	ConfirmFormat(List(4, 3, 2, 1), "(4 3 2 1)")

	c := List(4, 3, 2, 1)
	ConfirmFormat(c, "(4 3 2 1)")
	ConfirmFormat(List(5, c, 0), "(5 (4 3 2 1) 0)")
	c = List(5, c, 0)
	ConfirmFormat(c, "(5 (4 3 2 1) 0)")
	c.Tail.Tail.Tail = c
	ConfirmFormat(c, "(5 (4 3 2 1) 0 ...)")
	c.Tail.Head = c
	ConfirmFormat(c, "(5 (...) 0 ...)")
}

func TestConsLen(t *testing.T) {
	ConfirmLen := func(c *ConsCell, x int, b bool) {
		switch i, r := c.Len(); {
		case r != b:	t.Fatalf("'%v' recursion should be %v but is %v", c.String(), b, r)
		case i != x:	t.Fatalf("'%v' length should be %v but is %v", c.String(), x, i)
		}
	}
	ConfirmLen(List(4, 3, 2, 1), 4, false)
	ConfirmLen(List(4, List(3), 2, 1), 4, false)

	c := List(4, 3, 2, 1)
	c.Tail.Tail.Tail.Tail = c
	ConfirmLen(c, 4, true)
	c.Tail.Tail.Tail = c
	ConfirmLen(c, 3, true)
	c.Tail.Tail = c
	ConfirmLen(c, 2, true)
	c.Tail = c
	ConfirmLen(c, 1, true)

	c = List(4, 3, 2, 1)
	c.Tail.Head = c
	ConfirmLen(c, 4, false)
}

func TestConsDepth(t *testing.T) {
	ConfirmDepth := func(c *ConsCell, x int) {
		if i := c.Depth(); i != x {
			t.Fatalf("'%v' depth should be %v but is %v", c.String(), x, i)
		}
	}
	ConfirmDepth(List(	4, 3, 2, 1), 0)
	ConfirmDepth(List(	5,
						List(4, 3),
						2), 1)
	ConfirmDepth(List(	6,
						List(	5,
								List(4, 3, 2)),
						1), 2)
	ConfirmDepth(List(	7,
						List(	6,
								List(	5,
										4,
										List(3, 2),
										1)),
								0), 3)
	ConfirmDepth(List(	8,
						List(	7,
								List(	6,
										5,
										List(4, 3),
								2)),
								List(	1,
										List(0, -1))), 3)
	ConfirmDepth(List(	9,
						List(	8,
								List(	7,
										List(	6, 5)),
										List(	4,
												3,
												List(2, 1),
												0))), 3)
	ConfirmDepth(List(	'A',
						List(	9,
								SCons(	8,
										SCons(7, 6)),
								List(	5,
										4,
										List(3, 2),
										1))), 3)
	ConfirmDepth(List(	'B',
						List(	'A',
								SCons(	9,
										SCons(	8,
												SCons(7, 6))),
								List(	5,
										4,
										List(3, 2),
										1))), 4)
}

func TestConsEach(t *testing.T) {
	c := List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	count := 0
	c.Each(func(i interface{}) {
		if i != count { t.Fatalf("element %v erroneously reported as %v", count, i) }
		count++
	})
}

func TestConsReverse(t *testing.T) {
	ConfirmReverse := func(c, r *ConsCell) {
		c.Reverse()
		if !c.Equal(r) {
			t.Fatalf("%v should be %v", c, r)
		}
	}
	c := List(1)
	ConfirmReverse(c, List(1))
	ConfirmReverse(c, List(1))

	c = List(1, 2)
	ConfirmReverse(c, List(2, 1))
	ConfirmReverse(c, List(1, 2))

	c = List(1, 2, 3)
	ConfirmReverse(c, List(3, 2, 1))
	ConfirmReverse(c, List(1, 2, 3))

	c = List(1, 2, 3, 4)
	ConfirmReverse(c, List(4, 3, 2, 1))
	ConfirmReverse(c, List(1, 2, 3, 4))
}

func TestConsFlatten(t *testing.T) {
	t.Fatal()
}

func TestConsAt(t *testing.T) {
	t.Fatal()
}

func TestConsSet(t *testing.T) {
	t.Fatal()
}

func TestConsLink(t *testing.T) {
	ConfirmLink := func(c *ConsCell, to, from int, r *ConsCell) {
		switch {
		case !c.Link(to, from):		t.Fatalf("c.Link() failed")
		case !c.Equal(r):			t.Fatalf("%v should be %v", c, r)
		}
	}
	c := List(0)
	c.Tail = c
	ConfirmLink(List(0, 1, 2, 3), 0, 0, c)
	ConfirmLink(List(0, 1, 2, 3), 1, 0, List(0, 1, 2, 3))
	ConfirmLink(List(0, 1, 2, 3), 2, 0, List(0, 2, 3))
	ConfirmLink(List(0, 1, 2, 3), 3, 0, List(0, 3))

	c = List(0, 1)
	c.Tail.Tail = c
	ConfirmLink(List(0, 1, 2, 3), 0, 1, c)

	c = List(0, 1, 2)
	c.Tail.Tail.Tail = c
	ConfirmLink(List(0, 1, 2, 3), 0, 2, c)

	c = List(0, 1, 2, 3)
	c.Tail.Tail.Tail.Tail = c
	ConfirmLink(List(0, 1, 2, 3), 0, 3, c)
}

func TestEnd(t *testing.T) {
	ConfirmEnd := func(c, r *ConsCell) {
		x := c.End()
		if !r.Equal(x) {
			t.Fatalf("%v should be %v", x, r)
		}
	}
	ConfirmEnd(List(0), List(0))
	ConfirmEnd(List(0, 1), List(1))
	ConfirmEnd(List(0, 1, 2), List(2))
	ConfirmEnd(List(0, 1, 2, 3), List(3))

	c := List(0, 1, 2, 3, 4, 5)
	c.Tail = c
	r := List(0)
	r.Tail = r
	ConfirmEnd(c, r)

	c = List(0, 1, 2, 3, 4, 5)
	c.Tail.Tail = c
	r = List(0, 1)
	r.Tail.Tail = r
	ConfirmEnd(c, r)

	c = List(0, 1, 2, 3, 4, 5)
	c.Tail.Tail = c.Tail
	r = List(1)
	r.Tail = r
	ConfirmEnd(c, r)

	c = List(0, 1, 2, 3, 4, 5)
	c.Tail.Tail.Tail = c.Tail
	r = List(1, 2)
	r.Tail.Tail = r
	ConfirmEnd(c, r)
}