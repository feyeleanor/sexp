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

func TestCycListAppend(t *testing.T) {
	ConfirmAppend := func(c CycList, v interface{}, r CycList) {
		c.Append(v)
		if !c.Equal(r) {
			t.Fatalf("%v should be %v", c, r)
		}
	}
	ConfirmAppend(Loop(), 1, Loop(1))
	ConfirmAppend(Loop(1), 2, Loop(1, 2))
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
	ConfirmFormat(r.Next().node.Head.(CycList), "(0 (0 ...) ...)")

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
		t.Logf("1. c = %v", c)
		t.Logf("1. r = %v", r)
		c.Reverse()
		t.Logf("2. c = %v", c)
		if !c.Equal(r) {
			t.Fatalf("%v should be %v", c, r)
		}
//		t.Logf("3. c = %v", c)
	}

	c := Loop(1)
	ConfirmReverse(c, Loop(1))
	ConfirmReverse(c, Loop(1))

	ConfirmReverse(Loop(1, 2), Loop(2, 1))
	ConfirmReverse(Loop(2, 1), Loop(1, 2))

	c = Loop(1, 2)
/*	ConfirmReverse(c, Loop(2, 1))
//t.Logf("Test:: c = %v", c)
	ConfirmReverse(c, Loop(1, 2))

//	ConfirmReverse(Loop(1, 2, 3), Loop(3, 2, 1))
//	ConfirmReverse(Loop(3, 2, 1), Loop(1, 2, 3))

//	ConfirmReverse(Loop(1, 2, 3, 4), Loop(4, 3, 2, 1))
//	ConfirmReverse(Loop(4, 3, 2, 1), Loop(1, 2, 3, 4))
*/
}

func TestCycListFlatten(t *testing.T) {
	ConfirmFlatten := func(c, r CycList) {
		c.Flatten()
		if !c.Equal(r) {
t.Logf("c = '%v', c.length = '%v'", c, c.length)
t.Logf("r = '%v', r.length = '%v'", r, r.length)
			t.Fatalf("%v should be %v", c, r)
		}
	}
	ConfirmFlatten(Loop(), Loop())
	ConfirmFlatten(Loop(1), Loop(1))
	ConfirmFlatten(Loop(1, Loop(2)), Loop(1, Loop(2)))
	ConfirmFlatten(Loop(1, Loop(2, Loop(3))), Loop(1, Loop(2, Loop(3))))

	ConfirmFlatten(Loop(0, List(1)), Loop(0, 1))
	ConfirmFlatten(Loop(0, List(1, 2), 3), Loop(0, 1, 2, 3))
	ConfirmFlatten(Loop(0, List(1, List(2, 3), 4, List(5, List(6, 7)))), Loop(0, 1, 2, 3, 4, 5, 6, 7))
}