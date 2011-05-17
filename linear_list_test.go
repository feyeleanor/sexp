package sexp

import "testing"

func TestLinearListString(t *testing.T) {
	ConfirmFormat := func(l *LinearList, x string) {
		if s := l.String(); s != x {
			t.Fatalf("'%v' erroneously serialised as '%v'", x, s)
		}
	}

	ConfirmFormat(&LinearList{ start: nil, length: 0 }, "()")
	ConfirmFormat(List(0), "(0)")
	ConfirmFormat(List(0, nil), "(0 nil)")
	ConfirmFormat(List(1, List(0, nil)), "(1 (0 nil))")

	ConfirmFormat(List(1, 0, nil), "(1 0 nil)")


	c := List(10, List(0, 1, 2, 3))
	ConfirmFormat(c, "(10 (0 1 2 3))")
	ConfirmFormat(c.start.Tail.Head.(*LinearList), "(0 1 2 3)")
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
								SList(	8,
										SList(	7, 6 )),
								List(	5,
										4,
										List(3, 2),
										1))), 3)
	ConfirmDepth(List(	'B',
						List(	'A',
								SList(	9,
										SList(	8,
												SList( 7, 6 ))),
								List(	5,
										4,
										List(3, 2),
										1))), 4)
}

func TestLinearListEach(t *testing.T) {
	c := List(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	count := 0
	c.Each(func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
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

	l = List(1, List(2, 3), 4)
	ConfirmReverse(l, List(4, List(2, 3), 1))
	ConfirmReverse(l, List(1, List(2, 3), 4))
}

func TestLinearListFlatten(t *testing.T) {
	ConfirmFlatten := func(l, r *LinearList) {
		l.Flatten()
		if !l.Equal(r) {
			t.Logf("%v.Len() = %v, %v.Len() = %v", l, l.Len(), r, r.Len())
			t.Fatalf("'%v' should be '%v'", l, r)
		}
	}
	ConfirmFlatten(List(1), List(1))
	ConfirmFlatten(List(1, List(2)), List(1, 2))
	ConfirmFlatten(List(1, List(2, 3)), List(1, 2, 3))
	ConfirmFlatten(List(1, List(2, List(3))), List(1, 2, 3))
	ConfirmFlatten(List(1, List(2, 3, List(4, List(5)))), List(1, 2, 3, 4, 5))
	ConfirmFlatten(List(1, List(List(2, 3), List(4, List(5)))), List(1, 2, 3, 4, 5))
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

func TestLinearListClone(t *testing.T) {
	ConfirmClone := func(l, r *LinearList) {
		x := l.Clone()
		if !x.Equal(r) {
			t.Fatalf("%v should be %v", x, r)
		}
	}
	ConfirmClone(List(), List())
	ConfirmClone(List(0), List(0))
	ConfirmClone(List(0, 1), List(0, 1))
	ConfirmClone(List(0, List(0, 1)), List(0, List(0, 1)))
}

func TestLinearListAppend(t *testing.T) {
	ConfirmAppend := func(l *LinearList, v interface{}, r *LinearList) {
		l.Append(v)
		if !l.Equal(r) {
			t.Fatalf("%v should be %v", l, r)
		}
	}
	ConfirmAppend(List(), 0, List(0))
	ConfirmAppend(List(0), 1, List(0, 1))
	ConfirmAppend(List(0, 1), 2, List(0, 1, 2))
	ConfirmAppend(List(0, 1, 2), 3, List(0, 1, 2, 3))
}

func TestLinearListAppendSlice(t *testing.T) {
	ConfirmAppendSlice := func(l *LinearList, s []interface{}, r *LinearList) {
		l.AppendSlice(s)
		if !l.Equal(r) {
			t.Fatalf("%v should be %v", l, r)
		}
	}
	ConfirmAppendSlice(List(), []interface{}{ 1 }, List(1))
	ConfirmAppendSlice(List(1), []interface{}{ 2, 3 }, List(1, 2, 3))
}

func TestLinearListDelete(t *testing.T) {
	ConfirmDelete := func(l *LinearList, from, to int, r *LinearList) {
		l.Delete(from, to)
		switch {
		case !l.Equal(r):			t.Fatalf("Delete(%v, %v) should be '%v' and not '%v'", from, to, r, l)
		case l.Len() != r.Len():	t.Fatalf("Delete(%v, %v) length be '%v' and not '%v'", from, to, r.Len(), l.Len())
		}
	}
	ConfirmDelete(List(0, 1, 2, 3), -1, 0, List(0, 1, 2, 3))
	ConfirmDelete(List(0, 1, 2, 3), 0, -1, List(0, 1, 2, 3))
	ConfirmDelete(List(0, 1, 2, 3), 0, 4, List(0, 1, 2, 3))
	ConfirmDelete(List(0, 1, 2, 3), 4, 0, List(0, 1, 2, 3))

	ConfirmDelete(List(0, 1, 2, 3), 0, 0, List(1, 2, 3))
	ConfirmDelete(List(0, 1, 2, 3), 0, 1, List(2, 3))
	ConfirmDelete(List(0, 1, 2, 3), 0, 2, List(3))
	ConfirmDelete(List(0, 1, 2, 3), 0, 3, List())

	ConfirmDelete(List(0, 1, 2, 3), 1, 3, List(0))
	ConfirmDelete(List(0, 1, 2, 3), 2, 3, List(0, 1))
	ConfirmDelete(List(0, 1, 2, 3), 3, 3, List(0, 1, 2))

	ConfirmDelete(List(0, 1, 2, 3), 1, 1, List(0, 2, 3))
	ConfirmDelete(List(0, 1, 2, 3), 1, 2, List(0, 3))
	ConfirmDelete(List(0, 1, 2, 3), 2, 2, List(0, 1, 3))
}

func TestLinearListCut(t *testing.T) {
	ConfirmCut := func(l *LinearList, from, to int, r1, r2 *LinearList) {
		x := l.Cut(from, to)
		switch {
		case !x.Equal(r1):			t.Fatalf("Cut(%v, %v) cut should be '%v' and not '%v'", from, to, r1, x)
		case !l.Equal(r2):			t.Fatalf("Cut(%v, %v) remainder should be '%v' and not '%v'", from, to, r2, l)
		case x.Len() != r1.Len():	t.Fatalf("Cut(%v, %v) cut length should be '%v' and not '%v'", from, to, r1.Len(), x.Len())
		case l.Len() != r2.Len():	t.Fatalf("Cut(%v, %v) remainder length should be '%v' and not '%v'", from, to, r2.Len(), l.Len())
		}
	}
	ConfirmCut(List(0, 1, 2, 3), -1, 0, List(), List(0, 1, 2, 3))
	ConfirmCut(List(0, 1, 2, 3), 1, 0, List(), List(0, 1, 2, 3))
	ConfirmCut(List(0, 1, 2, 3), 0, 4, List(), List(0, 1, 2, 3))

	ConfirmCut(List(0, 1, 2, 3), 0, 0, List(0), List(1, 2, 3))
	ConfirmCut(List(0, 1, 2, 3), 0, 1, List(0, 1), List(2, 3))
	ConfirmCut(List(0, 1, 2, 3), 0, 2, List(0, 1, 2), List(3))
	ConfirmCut(List(0, 1, 2, 3), 0, 3, List(0, 1, 2, 3), List())

	ConfirmCut(List(0, 1, 2, 3), 1, 3, List(1, 2, 3), List(0))
	ConfirmCut(List(0, 1, 2, 3), 2, 3, List(2, 3), List(0, 1))
	ConfirmCut(List(0, 1, 2, 3), 3, 3, List(3), List(0, 1, 2))

	ConfirmCut(List(0, 1, 2, 3), 1, 1, List(1), List(0, 2, 3))
	ConfirmCut(List(0, 1, 2, 3), 1, 2, List(1, 2), List(0, 3))
	ConfirmCut(List(0, 1, 2, 3), 2, 2, List(2), List(0, 1, 3))
}

func TestLinearListInsert(t *testing.T) {
	t.Fatal("Write Tests")
}

func TestLinearListAbsorb(t *testing.T) {
	ConfirmAbsorb := func(l *LinearList, i int, s, r *LinearList) {
		switch ok := l.Absorb(i, s); {
		case !ok:				t.Fatalf("Absorb(%v, ...) should return true", i)
		case s.NotNil():		t.Fatalf("Absorb(%v, ...) source should be nil and not %v", i, s)
		case !l.Equal(r):		t.Fatalf("Absorb(%v, ...) result should be '%v' and not %v", i, r, l)
		}
	}
	RefuteAbsorb := func(l *LinearList, i int, s, r *LinearList) {
		source := s.Clone()
		switch ok := l.Absorb(i, s); {
		case ok:				t.Fatalf("Absorb(%v, ...) should return false", i)
		case !s.Equal(source):	t.Fatalf("Absorb(%v, ...) source should be %v and not %v", i, source, s)
		case !l.Equal(r):		t.Fatalf("Absorb(%v, ...) result should be '%v' and not '%v'", i, r, l)
		}
	}

	RefuteAbsorb(List(), -1, List(-3, -2, -1), List())
	ConfirmAbsorb(List(), 0, List(-3, -2, -1), List(-3, -2, -1))
	RefuteAbsorb(List(), 1, List(-3, -2, -1), List())

	RefuteAbsorb(List(0, 1, 2, 3), -1, List(-3, -2, -1), List(0, 1, 2, 3))
	ConfirmAbsorb(List(0, 1, 2, 3), 0, List(-3, -2, -1), List(-3, -2, -1, 0, 1, 2, 3))
	ConfirmAbsorb(List(0, 1, 2, 3), 1, List(-3, -2, -1), List(0, -3, -2, -1, 1, 2, 3))
	ConfirmAbsorb(List(0, 1, 2, 3), 2, List(-3, -2, -1), List(0, 1, -3, -2, -1, 2, 3))
	ConfirmAbsorb(List(0, 1, 2, 3), 3, List(-3, -2, -1), List(0, 1, 2, -3, -2, -1, 3))
	ConfirmAbsorb(List(0, 1, 2, 3), 4, List(-3, -2, -1), List(0, 1, 2, 3, -3, -2, -1))
	RefuteAbsorb(List(0, 1, 2, 3), 5, List(-3, -2, -1), List(0, 1, 2, 3))
}

func TestLinearListCar(t *testing.T) {
	ConfirmCar := func(l *LinearList, r interface{}) {
		if x := l.Car(); x != r {
			t.Fatalf("Car '%v' should be '%v' but is '%v'", l, r, x)
		}
	}
	ConfirmCar(List(), nil)
	ConfirmCar(List(0), 0)
}

func TestLinearListCdr(t *testing.T) {
	ConfirmCdr := func(l, r *LinearList) {
		if x := l.Cdr(); !x.Equal(r) {
			t.Fatalf("Cdr '%v' should be '%v' but is '%v'", l, r, x)
		}
	}
	ConfirmCdr(List(), List())
	ConfirmCdr(List(0), List())
	ConfirmCdr(List(0, 1), List(1))
	ConfirmCdr(List(0, 1, 2), List(1, 2))
}

func TestLinearListRplaca(t *testing.T) {
	ConfirmRplaca := func(l *LinearList, v interface{}, r interface{}) {
		l.Rplaca(v)
		if !l.Equal(r) {
			t.Fatalf("Rplaca should be '%v' but is '%v'", r, l)
		}
	}
	ConfirmRplaca(List(), 0, List(0))
	ConfirmRplaca(List(0), 1, List(1))
	ConfirmRplaca(List(0, 1), 1, List(1, 1))
}

func TestLinearListRplacd(t *testing.T) {
	ConfirmRplacd := func(l, n, r *LinearList) {
		l.Rplacd(n)
		if !l.Equal(r) {
			t.Fatalf("Rplacd should be '%v' but is '%v'", r, l)
		}
	}
	ConfirmRplacd(List(), List(0), List(0))
	ConfirmRplacd(List(0), List(1), List(0, 1))
	ConfirmRplacd(List(0, 1), List(2), List(0, 2))
	ConfirmRplacd(List(0, 1, 2), List(3, 4), List(0, 3, 4))
}