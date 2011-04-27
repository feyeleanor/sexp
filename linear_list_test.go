package sexp

import "testing"

func TestLinearListIsNil(t *testing.T) {
	l := List()
	if !l.IsNil() { t.Fatalf("%v.IsNil() should be true", l) }

	l = List(0)
	if l.IsNil() { t.Fatalf("%v.IsNil() should be false", l) }

	l = List(&Node{})
	if l.IsNil() { t.Fatalf("%v.IsNil() should be false", l) }
}

func TestLinearListString(t *testing.T) {
	ConfirmFormat := func(l *LinearList, x string) {
		if s := l.String(); s != x {
			t.Fatalf("'%v' erroneously serialised as '%v'", x, s)
		}
	}

	ConfirmFormat(&LinearList{ &Node{nil, nil} }, "()")
	ConfirmFormat(List(0), "(0)")
	ConfirmFormat(List(0, nil), "(0 nil)")
	ConfirmFormat(List(1, List(0, nil)), "(1 (0 nil))")

	ConfirmFormat(List(1, 0, nil), "(1 0 nil)")


	c := List(10, List(0, 1, 2, 3))
	ConfirmFormat(c, "(10 (0 1 2 3))")
	ConfirmFormat(c.Tail.Head.(*LinearList), "(0 1 2 3)")
}

func TestLinearListList(t *testing.T) {
	ConfirmFormat := func(l *LinearList, x string) {
		if s := l.String(); s != x {
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
}

func TestLinearListLen(t *testing.T) {
	ConfirmLen := func(l *LinearList, x int) {
		if i := l.Len(); i != x {
			t.Fatalf("'%v' length should be %v but is %v", l.String(), x, i)
		}
	}
	ConfirmLen(List(4, 3, 2, 1), 4)
	ConfirmLen(List(4, List(3, 3, 3), 2, 1), 4)
}

func TestLinearListDepth(t *testing.T) {
	ConfirmDepth := func(l *LinearList, x int) {
		if i := l.Depth(); i != x {
			t.Fatalf("'%v' depth should be %v but is %v", l.String(), x, i)
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

func TestLinearListEach(t *testing.T) {
	c := List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	count := 0
	c.Each(func(i interface{}) {
		if i != count { t.Fatalf("element %v erroneously reported as %v", count, i) }
		count++
	})
}

func TestLinearListReverse(t *testing.T) {
	ConfirmReverse := func(l, r *LinearList) {
		l.Reverse()
		if !r.Equal(l) {
			t.Fatalf("'%v' should be '%v'", l, r)
		}
	}
	l := List(1)
	ConfirmReverse(l, List(1))
	ConfirmReverse(l, List(1))

	l = List(1, 2)
	ConfirmReverse(l, List(2, 1))
	ConfirmReverse(l, List(1, 2))

	l = List(1, 2, 3)
	ConfirmReverse(l, List(3, 2, 1))
	ConfirmReverse(l, List(1, 2, 3))

	l = List(1, 2, 3, 4)
	ConfirmReverse(l, List(4, 3, 2, 1))
	ConfirmReverse(l, List(1, 2, 3, 4))
}

func TestLinearListFlatten(t *testing.T) {
	t.Fatal()
}

func TestLinearListAt(t *testing.T) {
	ConfirmAt := func(l *LinearList, i int, v interface{}) {
		if l.At(i) != v {
			t.Fatalf("List[%v] should be %v but is %v", i, v, l.At(i))
		}
	}
	l := List(10, 11, 12, 13, 14, 15, 16, 17)
	ConfirmAt(l, 0, 10)
	ConfirmAt(l, 1, 11)
	ConfirmAt(l, 2, 12)
	ConfirmAt(l, 3, 13)
	ConfirmAt(l, 4, 14)
	ConfirmAt(l, 5, 15)
	ConfirmAt(l, 6, 16)
	ConfirmAt(l, 7, 17)
}

func TestLinearListSet(t *testing.T) {
	ConfirmSet := func(l *LinearList, i int, v interface{}) {
		l.Set(i, v)
		if l.At(i) != v {
			t.Fatalf("List[%v] should be %v but is %v", i, v, l.At(i))
		}
	}
	l := List(10, 11, 12, 13, 14, 15, 16, 17)
	ConfirmSet(l, 0, 20)
	ConfirmSet(l, 1, 21)
	ConfirmSet(l, 2, 22)
	ConfirmSet(l, 3, 23)
	ConfirmSet(l, 4, 24)
	ConfirmSet(l, 5, 25)
	ConfirmSet(l, 6, 26)
	ConfirmSet(l, 7, 27)
}

func TestLinearListCut(t *testing.T) {
	ConfirmCut := func(l *LinearList, from, to int, r *LinearList) {
		switch {
		case !l.Cut(from, to):		t.Fatalf("c.Cut() failed")
		case !l.Equal(r):			t.Fatalf("%v should be %v", l, r)
		}
	}
	ConfirmCut(List(0, 1, 2, 3), 0, 1, List(0, 1, 2, 3))
	ConfirmCut(List(0, 1, 2, 3), 0, 2, List(0, 2, 3))
	ConfirmCut(List(0, 1, 2, 3), 0, 3, List(0, 3))
}

func TestLinearListEnd(t *testing.T) {
	ConfirmEnd := func(l, r *LinearList) {
		x := l.End()
		if !r.Equal(x) {
			t.Fatalf("%v should be %v", x, r)
		}
	}
	ConfirmEnd(List(0), List(0))
	ConfirmEnd(List(0, 1), List(1))
	ConfirmEnd(List(0, 1, 2), List(2))
	ConfirmEnd(List(0, 1, 2, 3), List(3))
}