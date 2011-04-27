package sexp

import "fmt"
import "testing"

func TestCycListIsNil(t *testing.T) {
	c := CList()
	if !c.IsNil() { t.Fatalf("%v.IsNil() should be true", c) }

	c = CList(0)
	if c.IsNil() { t.Fatalf("%v.IsNil() should be false", c) }

	c = &CycList{&Node{ Tail: c.Node }}
	if c.IsNil() { t.Fatalf("%v.IsNil() should be false", c) }

	c = &CycList{&Node{ Head: c, Tail: c.Node }}
	if c.IsNil() { t.Fatalf("%v.IsNil() should be false", c) }
}

func TestCycListString(t *testing.T) {
	ConfirmFormat := func(c *CycList, x string) {
		if s := c.String(); s != x {
			t.Fatalf("'%v' erroneously serialised as '%v'", x, s)
		}
	}

	ConfirmFormat(&CycList{ &Node{nil, nil} }, "()")
	ConfirmFormat(CList(0), "(0)")
	ConfirmFormat(CList(0, nil), "(0 nil)")
	ConfirmFormat(CList(1, CList(0, nil)), "(1 (0 nil))")

	ConfirmFormat(CList(1, 0, nil), "(1 0 nil)")

	c := CList(0)
	c.Node.Tail = c.Node
	ConfirmFormat(c, "(0 ...)")

	c = CList(0, c)
	ConfirmFormat(c, "(0 (0 ...))")

	c.Node.Tail.Head = c
	ConfirmFormat(c, "(0 (...))")

	c.Node.Tail.Tail = c.Node
	ConfirmFormat(c, "(0 (...) ...)")

	r := CList(10, c)
	ConfirmFormat(r, "(10 (0 (...) ...))")
	ConfirmFormat(&CycList{ r.Node.Tail }, "((0 (...) ...))")
	ConfirmFormat(r.Node.Tail.Head.(*CycList), "(0 (...) ...)")

	c.Node.Tail = CList(-1, -2, c).Node
	ConfirmFormat(r, "(10 (0 -1 -2 (...)))")

	c.Node.Tail = CList(3, 4, c).Node
	ConfirmFormat(r, "(10 (0 3 4 (...)))")

	c.Node.Tail.Tail = CList(3).Node
	ConfirmFormat(r, "(10 (0 3 3))")

t.Fatal()
	c.Node.Tail.Tail = CList(c, CList(-1, -2, c).Node).Node
	ConfirmFormat(c, "(0 -1 -2 (...) (-1 -2 ...))")

t.Fatal()

	ConfirmFormat(r, fmt.Sprintf("(10 (1 0 cons(%v) -1 -2 ...))", printAddress(c)))

//	r = Cons(10, Cons(9, c))
	r = CList(10, 9, c)
	ConfirmFormat(r, fmt.Sprintf("(10 9 (1 0 cons(%v) -1 -2 ...))", printAddress(c)))

//	r = Cons(10, Cons(9, Cons(8, c)))
	r = CList(10, 9, 8, c)
	ConfirmFormat(r, fmt.Sprintf("(10 9 8 (1 0 cons(%v) -1 -2 ...))", printAddress(c)))

//	r = Cons(10, Cons(9, Cons(8, Cons(7, c))))
	r = CList(10, 9, 8, 7, c)
	ConfirmFormat(r, fmt.Sprintf("(10 9 8 7 (1 0 cons(%v) -1 -2 ...))", printAddress(c)))
}

func TestCycList(t *testing.T) {
	ConfirmFormat := func(c *CycList, x string) {
		if s := c.String(); s != x {
			t.Fatalf("'%v' erroneously serialised as '%v'", x, s)
		}
	}
	ConfirmFormat(CList(), "()")
	ConfirmFormat(CList(1), "(1)")
	ConfirmFormat(CList(2, 1), "(2 1)")
	ConfirmFormat(CList(3, 2, 1), "(3 2 1)")
	ConfirmFormat(CList(4, 3, 2, 1), "(4 3 2 1)")

	c := CList(4, 3, 2, 1)
	ConfirmFormat(c, "(4 3 2 1)")
	ConfirmFormat(CList(5, c, 0), "(5 (4 3 2 1) 0)")
	c = CList(5, c, 0)
	ConfirmFormat(c, "(5 (4 3 2 1) 0)")
	c.Tail.Tail.Tail = c.Node
	ConfirmFormat(c, "(5 (4 3 2 1) 0 ...)")
	c.Tail.Head = c
	ConfirmFormat(c, "(5 (...) 0 ...)")
}

func TestCycListLen(t *testing.T) {
	ConfirmLen := func(c *CycList, x int, b bool) {
		defer func() {
			r := recover()
			switch {
			case r == nil:	if b { t.Fatalf("'%v' recursion should be false but is true", c.String()) }
			default:		switch {
							case !b:		t.Fatalf("'%v' recursion should be true but is false", c.String())
							case x != r:	t.Fatalf("'%v' length should be %v but is %v", c.String(), r, x)
							}
			}
		}()
		i := c.Len()
		if i != x { t.Fatalf("'%v' length should be %v but is %v", c.String(), x, i) }
	}
	ConfirmLen(CList(4, 3, 2, 1), 4, false)
	ConfirmLen(CList(4, CList(3), 2, 1), 4, false)

	c := CList(4, 3, 2, 1)
	c.Tail.Tail.Tail.Tail = c.Node
	ConfirmLen(c, 4, true)
	c.Tail.Tail.Tail = c.Node
	ConfirmLen(c, 3, true)
	c.Tail.Tail = c.Node
	ConfirmLen(c, 2, true)
	c.Tail = c.Node
	ConfirmLen(c, 1, true)

	c = CList(4, 3, 2, 1)
	c.Tail.Head = c.Node
	ConfirmLen(c, 4, false)
}

func TestCycListDepth(t *testing.T) {
	ConfirmDepth := func(c *CycList, x int) {
		if i := c.Depth(); i != x {
			t.Fatalf("'%v' depth should be %v but is %v", c.String(), x, i)
		}
	}
	ConfirmDepth(CList(	4, 3, 2, 1), 0)
	ConfirmDepth(CList(	5,
						CList(4, 3),
						2), 1)
	ConfirmDepth(CList(	6,
						CList(	5,
								CList(4, 3, 2)),
						1), 2)
	ConfirmDepth(CList(	7,
						CList(	6,
								CList(	5,
										4,
										CList(3, 2),
										1)),
								0), 3)
	ConfirmDepth(CList(	8,
						CList(	7,
								CList(	6,
										5,
										CList(4, 3),
								2)),
								CList(	1,
										CList(0, -1))), 3)
	ConfirmDepth(CList(	9,
						CList(	8,
								CList(	7,
										CList(	6, 5)),
										CList(	4,
												3,
												CList(2, 1),
												0))), 3)
	ConfirmDepth(CList(	'A',
						CList(	9,
								SCons(	8,
										SCons(7, 6)),
								CList(	5,
										4,
										CList(3, 2),
										1))), 3)
	ConfirmDepth(CList(	'B',
						CList(	'A',
								SCons(	9,
										SCons(	8,
												SCons(7, 6))),
								CList(	5,
										4,
										CList(3, 2),
										1))), 4)
}

func TestCycListEach(t *testing.T) {
	c := CList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	count := 0
	c.Each(func(i interface{}) {
		if i != count { t.Fatalf("element %v erroneously reported as %v", count, i) }
		count++
	})
}

func TestCycListReverse(t *testing.T) {
	ConfirmReverse := func(c, r *CycList) {
		c.Reverse()
		if !c.Equal(r) {
			t.Fatalf("%v should be %v", c, r)
		}
	}
	c := CList(1)
	ConfirmReverse(c, CList(1))
	ConfirmReverse(c, CList(1))

	c = CList(1, 2)
	ConfirmReverse(c, CList(2, 1))
	ConfirmReverse(c, CList(1, 2))

	c = CList(1, 2, 3)
	ConfirmReverse(c, CList(3, 2, 1))
	ConfirmReverse(c, CList(1, 2, 3))

	c = CList(1, 2, 3, 4)
	ConfirmReverse(c, CList(4, 3, 2, 1))
	ConfirmReverse(c, CList(1, 2, 3, 4))

	c = CList(0, 1)
	c.Tail.Tail = c.Node
	r := CList(1, 0)
	r.Tail.Tail = r.Node
	ConfirmReverse(c, r)
}

func TestCycListFlatten(t *testing.T) {
	t.Fatal()
}

func TestCycListAt(t *testing.T) {
	ConfirmAt := func(c *CycList, i int, v interface{}) {
		if c.At(i) != v {
			t.Fatalf("List[%v] should be %v but is %v", i, v, c.At(i))
		}
	}
	c := CList(10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	c.Link(0, 9)
	ConfirmAt(c, 0, 10)
	ConfirmAt(c, 1, 11)
	ConfirmAt(c, 2, 12)
	ConfirmAt(c, 3, 13)
	ConfirmAt(c, 4, 14)
	ConfirmAt(c, 5, 15)
	ConfirmAt(c, 6, 16)
	ConfirmAt(c, 7, 17)
	ConfirmAt(c, 8, 18)
	ConfirmAt(c, 9, 19)
	ConfirmAt(c, 10, 10)
	ConfirmAt(c, 11, 11)
	ConfirmAt(c, 12, 12)
	ConfirmAt(c, 13, 13)
	ConfirmAt(c, 14, 14)
	ConfirmAt(c, 15, 15)
	ConfirmAt(c, 16, 16)
	ConfirmAt(c, 17, 17)
	ConfirmAt(c, 18, 18)
	ConfirmAt(c, 19, 19)
}

func TestCycListSet(t *testing.T) {
	ConfirmAt := func(c *CycList, i int, v interface{}) {
		if c.At(i) != v {
			t.Fatalf("List[%v] should be %v but is %v", i, v, c.At(i))
		}
	}
	ConfirmSet := func(c *CycList, i int, v interface{}) {
		c.Set(i, v)
		if c.At(i) != v {
			t.Fatalf("List[%v] should be %v but is %v", i, v, c.At(i))
		}
	}
	c := CList(10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	c.Link(0, 9)
	ConfirmSet(c, 0, 0)
	ConfirmAt(c, 0, 0)
	ConfirmSet(c, 1, 1)
	ConfirmAt(c, 1, 1)
	ConfirmSet(c, 2, 2)
	ConfirmAt(c, 2, 2)
	ConfirmSet(c, 3, 3)
	ConfirmAt(c, 3, 3)
	ConfirmSet(c, 4, 4)
	ConfirmAt(c, 4, 4)
	ConfirmSet(c, 5, 5)
	ConfirmAt(c, 5, 5)
	ConfirmSet(c, 6, 6)
	ConfirmAt(c, 6, 6)
	ConfirmSet(c, 7, 7)
	ConfirmAt(c, 7, 7)
	ConfirmSet(c, 8, 8)
	ConfirmAt(c, 8, 8)
	ConfirmSet(c, 9, 9)
	ConfirmAt(c, 9, 9)
	ConfirmSet(c, 10, 10)
	ConfirmAt(c, 0, 10)
	ConfirmSet(c, 11, 11)
	ConfirmAt(c, 1, 11)
	ConfirmSet(c, 12, 12)
	ConfirmAt(c, 2, 12)
	ConfirmSet(c, 13, 13)
	ConfirmAt(c, 3, 13)
	ConfirmSet(c, 14, 14)
	ConfirmAt(c, 4, 14)
	ConfirmSet(c, 15, 15)
	ConfirmAt(c, 5, 15)
	ConfirmSet(c, 16, 16)
	ConfirmAt(c, 6, 16)
	ConfirmSet(c, 17, 17)
	ConfirmAt(c, 7, 17)
	ConfirmSet(c, 18, 18)
	ConfirmAt(c, 8, 18)
	ConfirmSet(c, 19, 19)
	ConfirmAt(c, 9, 19)
}

func TestCycListLink(t *testing.T) {
	ConfirmLink := func(c *CycList, to, from int, r *CycList) {
		switch {
		case !c.Link(to, from):		t.Fatalf("c.Link() failed")
		case !c.Equal(r):			t.Fatalf("%v should be %v", c, r)
		}
	}
	c := CList(0)
	c.Tail = c.Node
	ConfirmLink(CList(0, 1, 2, 3), 0, 0, c)
	ConfirmLink(CList(0, 1, 2, 3), 1, 0, CList(0, 1, 2, 3))
	ConfirmLink(CList(0, 1, 2, 3), 2, 0, CList(0, 2, 3))
	ConfirmLink(CList(0, 1, 2, 3), 3, 0, CList(0, 3))

	c = CList(0, 1)
	c.Tail.Tail = c.Node
	ConfirmLink(CList(0, 1, 2, 3), 0, 1, c)

	c = CList(0, 1, 2)
	c.Tail.Tail.Tail = c.Node
	ConfirmLink(CList(0, 1, 2, 3), 0, 2, c)

	c = CList(0, 1, 2, 3)
	c.Tail.Tail.Tail.Tail = c.Node
	ConfirmLink(CList(0, 1, 2, 3), 0, 3, c)
}

func TestCycListEnd(t *testing.T) {
	ConfirmEnd := func(c, r *CycList) {
		x := c.End()
		if !r.Equal(x) {
			t.Fatalf("%v should be %v", x, r)
		}
	}
	ConfirmEnd(CList(0), CList(0))
	ConfirmEnd(CList(0, 1), CList(1))
	ConfirmEnd(CList(0, 1, 2), CList(2))
	ConfirmEnd(CList(0, 1, 2, 3), CList(3))

	c := CList(0, 1, 2, 3, 4, 5)
	c.Tail = c.Node
	r := CList(0)
	r.Tail = r.Node
	ConfirmEnd(c, r)

	c = CList(0, 1, 2, 3, 4, 5)
	c.Tail.Tail = c.Node
	r = CList(0, 1)
	r.Tail.Tail = r.Node
	ConfirmEnd(c, r)

	c = CList(0, 1, 2, 3, 4, 5)
	c.Tail.Tail = c.Tail
	r = CList(1)
	r.Tail = r.Node
	ConfirmEnd(c, r)

	c = CList(0, 1, 2, 3, 4, 5)
	c.Tail.Tail.Tail = c.Tail
	r = CList(1, 2)
	r.Tail.Tail = r.Node
	ConfirmEnd(c, r)
}