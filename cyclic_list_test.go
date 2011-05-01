package sexp

//import "fmt"
import "testing"

func TestCycListIsNil(t *testing.T) {
	t.Log("Write Tests")
}

func TestCycListLen(t *testing.T) {
	ConfirmLen := func(c CycList, x int) {
		if i := c.Len(); i != x {
			t.Fatalf("'%v' length should be %v but is %v", c.String(), x, i)
		}
	}
	ConfirmLen(Loop(), 0)
	ConfirmLen(Loop(4), 1)
	ConfirmLen(Loop(4, 3, 2, 1), 4)
	ConfirmLen(Loop(4, Loop(3), 2, 1), 4)
}

func TestCycListEach(t *testing.T) {
	c := Loop(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	count := 0
	func() {
		defer func() {
			if x := recover(); x == nil {
				t.Fatalf("Each terminated without raising an exception")
			}
		}()
		c.Each(func(i interface{}) {
			if i != count { t.Fatalf("element %v erroneously reported as %v", count, i) }
			count++
			if count == c.Len() {
				panic(count)
			}
		})
	}()

}

func TestCycListAt(t *testing.T) {
	ConfirmAt := func(c CycList, i int, v interface{}) {
		if x, _ := c.At(i); x != v {
			t.Fatalf("List[%v] should be %v but is %v", i, v, x)
		}
	}
	RefuteAt := func(c CycList, i int) {
		if v, ok := c.At(i); ok {
			t.Fatalf("List[%v] erroneously returned %v", i, v)
		}
	}
	c := Loop(10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	RefuteAt(c, -1)
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
	RefuteAt(c, 10)
}

func TestCycListSet(t *testing.T) {
	ConfirmSet := func(c CycList, i int, v interface{}) {
		c.Set(i, v)
		if x, _ := c.At(i); x != v {
			t.Fatalf("List[%v] should be %v but is %v", i, v, x)
		}
	}
	RefuteSet := func(c CycList, i int, v interface{}) {
		c.Set(i, v)
		if x, ok := c.At(i); ok {
			t.Fatalf("List[%v] erroneously returned %v", i, x)
		}
	}
	c := Loop(10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	RefuteSet(c, -1, 10)
	ConfirmSet(c, 0, 10)
	ConfirmSet(c, 1, 10)
	ConfirmSet(c, 2, 10)
	ConfirmSet(c, 3, 10)
	ConfirmSet(c, 4, 10)
	ConfirmSet(c, 5, 10)
	ConfirmSet(c, 6, 10)
	ConfirmSet(c, 7, 10)
	ConfirmSet(c, 8, 10)
	ConfirmSet(c, 9, 10)
	RefuteSet(c, 10, 10)
}

func TestCycListNext(t *testing.T) {
	ConfirmNext := func(c, r CycList) {
		x := c.Next()
		if !x.Equal(r) {
			t.Fatalf("%v should be %v", x, r)
		}
	}
	ConfirmNext(Loop(), Loop())
	ConfirmNext(Loop(0), Loop(0))
	ConfirmNext(Loop(0, 1), Loop(1, 0))
	ConfirmNext(Loop(0, 1, 2), Loop(1, 2, 0))
	ConfirmNext(Loop(0, 1, 2, 3), Loop(1, 2, 3, 0))
	ConfirmNext(Loop(0, 1, Loop(2), 3), Loop(1, Loop(2), 3, 0))
}

func TestCycListEnd(t *testing.T) {
	ConfirmEnd := func(c, r CycList) {
		x := c.End()
		if !x.Equal(r) {
			t.Fatalf("%v should be %v", x, r)
		}
	}
	ConfirmEnd(Loop(), Loop())
	ConfirmEnd(Loop(0), Loop(0))
	ConfirmEnd(Loop(0, 1), Loop(1, 0))
	ConfirmEnd(Loop(0, 1, 2), Loop(2, 0, 1))
	ConfirmEnd(Loop(0, 1, 2, 3), Loop(3, 0, 1, 2))
	ConfirmEnd(Loop(0, 1, Loop(2), 3), Loop(3, 0, 1, Loop(2)))
}

func TestCycListString(t *testing.T) {
	ConfirmFormat := func(c CycList, x string) {
		if s := c.String(); s != x {
			t.Fatalf("'%v' erroneously serialised as '%v'", x, s)
		}
	}

	ConfirmFormat(Loop(), "()")
	ConfirmFormat(Loop(0), "(0 ...)")
	ConfirmFormat(Loop(0, nil), "(0 nil ...)")
	ConfirmFormat(Loop(0, Loop(0)), "(0 (0 ...) ...)")
	ConfirmFormat(Loop(1, Loop(0, nil)), "(1 (0 nil ...) ...)")

	ConfirmFormat(Loop(1, 0, nil), "(1 0 nil ...)")

	r := Loop(10, Loop(0, Loop(0)))
	ConfirmFormat(r, "(10 (0 (0 ...) ...) ...)")
	ConfirmFormat(r.Next(), "((0 (0 ...) ...) 10 ...)")
	ConfirmFormat(r.Next().Head.(CycList), "(0 (0 ...) ...)")

	r = Loop(r, 0, Loop(-1, -2, r))
	ConfirmFormat(r, "((10 (0 (0 ...) ...) ...) 0 (-1 -2 (10 (0 (0 ...) ...) ...) ...) ...)")
}

func TestLoop(t *testing.T) {
	ConfirmFormat := func(c CycList, x string) {
		if s := c.String(); s != x {
			t.Fatalf("'%v' erroneously serialised as '%v'", x, s)
		}
	}
	ConfirmFormat(Loop(), "()")
	ConfirmFormat(Loop(1), "(1 ...)")
	ConfirmFormat(Loop(2, 1), "(2 1 ...)")
	ConfirmFormat(Loop(3, 2, 1), "(3 2 1 ...)")
	ConfirmFormat(Loop(4, 3, 2, 1), "(4 3 2 1 ...)")

	c := Loop(4, 3, 2, 1)
	ConfirmFormat(c, "(4 3 2 1 ...)")
	ConfirmFormat(Loop(5, c, 0), "(5 (4 3 2 1 ...) 0 ...)")
	c = Loop(5, c, 0)
	ConfirmFormat(c, "(5 (4 3 2 1 ...) 0 ...)")
}

func TestCycListDepth(t *testing.T) {
	ConfirmDepth := func(c CycList, x int) {
		if i := c.Depth(); i != x {
			t.Fatalf("'%v' depth should be %v but is %v", c.String(), x, i)
		}
	}
	ConfirmDepth(Loop(	4, 3, 2, 1), 0)
	ConfirmDepth(Loop(	5,
						Loop(4, 3),
						2), 1)
	ConfirmDepth(Loop(	6,
						Loop(	5,
								Loop(4, 3, 2)),
						1), 2)
	ConfirmDepth(Loop(	7,
						Loop(	6,
								Loop(	5,
										4,
										Loop(3, 2),
										1)),
								0), 3)
	ConfirmDepth(Loop(	8,
						Loop(	7,
								Loop(	6,
										5,
										Loop(4, 3),
								2)),
								Loop(	1,
										Loop(0, -1))), 3)
	ConfirmDepth(Loop(	9,
						Loop(	8,
								Loop(	7,
										Loop(	6, 5)),
										Loop(	4,
												3,
												Loop(2, 1),
												0))), 3)
	ConfirmDepth(Loop(	'A',
						Loop(	9,
								SList(	8,
										SList(	7, 6 )),
								Loop(	5,
										4,
										Loop(3, 2),
										1))), 3)
	ConfirmDepth(Loop(	'B',
						Loop(	'A',
								SList(	9,
										SList(	8,
												SList( 7, 6 ))),
								Loop(	5,
										4,
										Loop(3, 2),
										1))), 4)
}

func TestCycListReverse(t *testing.T) {
	ConfirmReverse := func(c, r CycList) {
		c.Reverse()
		if !c.Equal(r) {
			t.Fatalf("%v should be %v", c, r)
		}
	}

	c := Loop(1)
	ConfirmReverse(c, Loop(1))
	ConfirmReverse(c, Loop(1))

	ConfirmReverse(Loop(1, 2), Loop(2, 1))
	ConfirmReverse(Loop(2, 1), Loop(1, 2))

	ConfirmReverse(Loop(1, 2, 3), Loop(3, 2, 1))
/*	ConfirmReverse(Loop(3, 2, 1), Loop(1, 2, 3))

	ConfirmReverse(Loop(1, 2, 3, 4), Loop(4, 3, 2, 1))
	ConfirmReverse(Loop(4, 3, 2, 1), Loop(1, 2, 3, 4))
*/
}

func TestCycListFlatten(t *testing.T) {
	t.Log("Write Tests")
}

func TestCycListLink(t *testing.T) {
	t.Log("Implement CycList::Link()")
/*	ConfirmLink := func(c CycList, to, from int, r CycList) {
		switch {
		case !c.Link(to, from):		t.Fatalf("c.Link() failed")
		case !c.Equal(r):			t.Fatalf("%v should be %v", c, r)
		}
	}
	c := Loop(0)
	c.Tail = c.Node
	ConfirmLink(Loop(0, 1, 2, 3), 0, 0, c)
	ConfirmLink(Loop(0, 1, 2, 3), 1, 0, Loop(0, 1, 2, 3))
	ConfirmLink(Loop(0, 1, 2, 3), 2, 0, Loop(0, 2, 3))
	ConfirmLink(Loop(0, 1, 2, 3), 3, 0, Loop(0, 3))

	c = Loop(0, 1)
	c.Tail.Tail = c.Node
	ConfirmLink(Loop(0, 1, 2, 3), 0, 1, c)

	c = Loop(0, 1, 2)
	c.Tail.Tail.Tail = c.Node
	ConfirmLink(Loop(0, 1, 2, 3), 0, 2, c)

	c = Loop(0, 1, 2, 3)
	c.Tail.Tail.Tail.Tail = c.Node
	ConfirmLink(Loop(0, 1, 2, 3), 0, 3, c)
*/}