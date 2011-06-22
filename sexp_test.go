package sexp

import "testing"

func TestEqual(t *testing.T) {
	ConfirmEqual := func(x, y interface{}) {
		if !Equal(x, y) {
			t.Fatalf("Equal(%v, %v) should be true", x, y)
		}
	}
	ConfirmEqual(0, 0)
	ConfirmEqual("a", "a")
	ConfirmEqual(List(0, 1, 2), List(0, 1, 2))
	t.Logf("Write more tests")
}

func TestLen(t *testing.T) {
	ConfirmLen := func(o interface{}, r int) {
		if x := Len(o); x != r {
			t.Fatalf("Len(%v) should be %v but is %v", o, r, x)
		}
	}
	ConfirmLen(0, 0)
	ConfirmLen(([]int)(nil), 0)
	ConfirmLen([]int{}, 0)
	ConfirmLen([]int{0}, 1)
	ConfirmLen([]int{0, 1, 2}, 3)
	ConfirmLen(List(0, 1, 2), 3)
	ConfirmLen(SList(0, 1, 2), 3)
}

func TestCap(t *testing.T) {
	ConfirmCap := func(o interface{}, r int) {
		if x := Cap(o); x != r {
			t.Fatalf("Len(%v) should be %v but is %v", o, r, x)
		}
	}
	ConfirmCap(0, 0)
	ConfirmCap(([]int)(nil), 0)
	ConfirmCap([]int{}, 0)
	ConfirmCap([]int{0}, 1)
	ConfirmCap(List(0, 1, 2), 0)
	ConfirmCap([]int{0, 1, 2}, 3)
	ConfirmCap(SList(0, 1, 2), 3)
}

/*
func TestEach(t *testing.T) { t.Fatal() }
func TestCycle(t *testing.T) { t.Fatal() }
func TestTransform(t *testing.T) { t.Fatal() }
func TestCollect(t *testing.T) { t.Fatal() }
*/

func TestReduce(t *testing.T) {
	ConfirmReduce := func(o, s, r interface{}, f func(m, x interface{}) interface{}) {
		if x := Reduce(o, s, f); !Equal(x, r) {
			t.Fatalf("Reduce(%v, %v, f) should be %v but is %v", o, s, r, x)
		}
	}

	Sum := func(memo, x interface{}) interface{} {
		return memo.(int) + x.(int)
	}

	ConfirmReduce(nil, nil, nil, Sum)
	ConfirmReduce(nil, 0, 0, Sum)
	ConfirmReduce([]int{0}, 0, 0, Sum)
	ConfirmReduce([]int{0, 1}, 0, 1, Sum)
	ConfirmReduce([]int{0, 1, 2}, 0, 3, Sum)
	ConfirmReduce([]int{0, 1, 2, 3}, 0, 6, Sum)
	ConfirmReduce([]int{0, 1, 2, 3, 4}, 0, 10, Sum)
}

func TestWhile(t *testing.T) {
	ConfirmWhile := func(o interface{}, r int, f func(i interface{}) bool) {
		if x := While(o, f); x != r {
			t.Fatalf("While(%v, f) should be %v but is %v", o, r, x)
		}
	}

	IsPositive := func(i interface{}) bool {
		if i, ok := i.(int); ok {
			return i > 0
		}
		return false
	}

	ConfirmWhile(nil, 0, IsPositive)
	ConfirmWhile([]int{}, 0, IsPositive)

	ConfirmWhile([]int{-1}, 0, IsPositive)
	ConfirmWhile([]int{0}, 0, IsPositive)
	ConfirmWhile([]int{1}, 1, IsPositive)

	ConfirmWhile([]int{-1, 0}, 0, IsPositive)
	ConfirmWhile([]int{0, 0}, 0, IsPositive)
	ConfirmWhile([]int{1, 0}, 1, IsPositive)
	ConfirmWhile([]int{1, 2, 0}, 2, IsPositive)
	ConfirmWhile([]int{1, 2, 3, 0}, 3, IsPositive)

	ConfirmWhile([]interface{}{-1, -2, -3, 0}, 0, IsPositive)
	ConfirmWhile([]interface{}{1, 2, 3, 0}, 3, IsPositive)
	ConfirmWhile([]interface{}{1, 2, 3, ""}, 3, IsPositive)
}

func TestUntil(t *testing.T) {
	ConfirmUntil := func(o interface{}, r int, f func(i interface{}) bool) {
		if x := Until(o, f); x != r {
			t.Fatalf("Until(%v, f) should be %v but is %v", o, r, x)
		}
	}

	IsPositive := func(i interface{}) bool {
		if i, ok := i.(int); ok {
			return i > 0
		}
		return false
	}

	ConfirmUntil(nil, 0, IsPositive)
	ConfirmUntil([]int{}, 0, IsPositive)

	ConfirmUntil([]int{-1}, 1, IsPositive)
	ConfirmUntil([]int{0}, 1, IsPositive)
	ConfirmUntil([]int{1}, 0, IsPositive)

	ConfirmUntil([]int{-1, 0}, 2, IsPositive)
	ConfirmUntil([]int{0, 0}, 2, IsPositive)
	ConfirmUntil([]int{1, 0}, 0, IsPositive)
	ConfirmUntil([]int{1, 2, 0}, 0, IsPositive)
	ConfirmUntil([]int{1, 2, 3, 0}, 0, IsPositive)

	ConfirmUntil([]interface{}{"test"}, 1, IsPositive)
	ConfirmUntil([]interface{}{-1, -2, 0}, 3, IsPositive)
	ConfirmUntil([]interface{}{-1, -2, -3, 0}, 4, IsPositive)
	ConfirmUntil([]interface{}{-1, -2, -3, "test"}, 4, IsPositive)
}

func TestAny(t *testing.T) {
	ConfirmAny := func(o interface{}, f func(i interface{}) bool) {
		if !Any(o, f) {
			t.Fatalf("Any(%v, f) should be true but is false", o)
		}
	}

	RefuteAny := func(o interface{}, f func(i interface{}) bool) {
		if Any(o, f) {
			t.Fatalf("Any(%v, f) should be false but is true", o)
		}
	}

	IsPositive := func(i interface{}) bool {
		if i, ok := i.(int); ok {
			return i > 0
		}
		return false
	}

	RefuteAny(nil, IsPositive)
	RefuteAny([]int{}, IsPositive)
	RefuteAny([]int{0}, IsPositive)
	ConfirmAny([]int{0, 1}, IsPositive)
	ConfirmAny([]int{0, 0, 1}, IsPositive)
}

func TestAll(t *testing.T) {
	ConfirmAll := func(o interface{}, f func(i interface{}) bool) {
		if !All(o, f) {
			t.Fatalf("All(%v, f) should be true but is false", o)
		}
	}

	RefuteAll := func(o interface{}, f func(i interface{}) bool) {
		if All(o, f) {
			t.Fatalf("All(%v, f) should be false but is true", o)
		}
	}

	IsPositive := func(i interface{}) bool {
		if i, ok := i.(int); ok {
			return i > 0
		}
		return false
	}

	RefuteAll(nil, IsPositive)
	RefuteAll([]int{}, IsPositive)
	RefuteAll([]int{0}, IsPositive)
	RefuteAll([]int{0, 1}, IsPositive)
	RefuteAll([]int{0, 0, 1}, IsPositive)

	ConfirmAll([]int{1}, IsPositive)
	ConfirmAll([]int{1, 1}, IsPositive)
	ConfirmAll([]int{1, 1, 1}, IsPositive)

	RefuteAll([]interface{}{}, IsPositive)
	RefuteAll([]interface{}{0}, IsPositive)
	RefuteAll([]interface{}{0, 1}, IsPositive)
	RefuteAll([]interface{}{0, 0, 1}, IsPositive)

	ConfirmAll([]interface{}{1}, IsPositive)
	ConfirmAll([]interface{}{1, 1}, IsPositive)
	ConfirmAll([]interface{}{1, 1, 1}, IsPositive)
}

func TestNone(t *testing.T) {
	ConfirmNone := func(o interface{}, f func(i interface{}) bool) {
		if !None(o, f) {
			t.Fatalf("None(%v, f) should be true but is false", o)
		}
	}

	RefuteNone := func(o interface{}, f func(i interface{}) bool) {
		if None(o, f) {
			t.Fatalf("None(%v, f) should be false but is true", o)
		}
	}

	IsPositive := func(i interface{}) bool {
		if i, ok := i.(int); ok {
			return i > 0
		}
		return false
	}

	ConfirmNone(nil, IsPositive)
	ConfirmNone([]int{}, IsPositive)
	ConfirmNone([]int{0}, IsPositive)
	RefuteNone([]int{0, 1}, IsPositive)
	RefuteNone([]int{0, 0, 1}, IsPositive)

	RefuteNone([]int{1}, IsPositive)
	RefuteNone([]int{1, 1}, IsPositive)
	RefuteNone([]int{1, 1, 1}, IsPositive)

	ConfirmNone([]interface{}{}, IsPositive)
	ConfirmNone([]interface{}{0}, IsPositive)
	RefuteNone([]interface{}{0, 1}, IsPositive)
	RefuteNone([]interface{}{0, 0, 1}, IsPositive)

	RefuteNone([]interface{}{1}, IsPositive)
	RefuteNone([]interface{}{1, 1}, IsPositive)
	RefuteNone([]interface{}{1, 1, 1}, IsPositive)
}

func TestOne(t *testing.T) {
	ConfirmOne := func(o interface{}, f func(i interface{}) bool) {
		if !One(o, f) {
			t.Fatalf("One(%v, f) should be true but is false", o)
		}
	}

	RefuteOne := func(o interface{}, f func(i interface{}) bool) {
		if One(o, f) {
			t.Fatalf("One(%v, f) should be false but is true", o)
		}
	}

	IsPositive := func(i interface{}) bool {
		if i, ok := i.(int); ok {
			return i > 0
		}
		return false
	}

	RefuteOne(nil, IsPositive)
	RefuteOne([]int{}, IsPositive)
	RefuteOne([]int{0}, IsPositive)
	ConfirmOne([]int{0, 1}, IsPositive)
	ConfirmOne([]int{0, 0, 1}, IsPositive)
	RefuteOne([]int{0, 0, 1, 1}, IsPositive)

	ConfirmOne([]int{1}, IsPositive)
	RefuteOne([]int{1, 1}, IsPositive)
	RefuteOne([]int{1, 1, 1}, IsPositive)

	RefuteOne([]interface{}{}, IsPositive)
	RefuteOne([]interface{}{0}, IsPositive)
	ConfirmOne([]interface{}{0, 1}, IsPositive)
	ConfirmOne([]interface{}{0, 0, 1}, IsPositive)
	RefuteOne([]interface{}{0, 0, 1, 1}, IsPositive)

	ConfirmOne([]interface{}{1}, IsPositive)
	RefuteOne([]interface{}{1, 1}, IsPositive)
	RefuteOne([]interface{}{1, 1, 1}, IsPositive)
}

func TestCount(t *testing.T) {
	ConfirmCount := func(o, r interface{}, f func(interface{}) bool) {
		if x := Count(o, f); !Equal(r, x) {
			t.Fatalf("Count(%v, f) should be %v but is %v", o, r, x)
		}
	}

	IsPositive := func(i interface{}) bool {
		if i, ok := i.(int); ok {
			return i > 0
		}
		return false
	}

	ConfirmCount([]int{}, 0, IsPositive)
	ConfirmCount([]interface{}{}, 0, IsPositive)

	ConfirmCount([]int{0}, 0, IsPositive)
	ConfirmCount([]int{1}, 1, IsPositive)
	ConfirmCount([]interface{}{"test"}, 0, IsPositive)
	ConfirmCount([]interface{}{1}, 1, IsPositive)


	ConfirmCount([]int{0, 1}, 1, IsPositive)
	ConfirmCount([]int{1, 2}, 2, IsPositive)
	ConfirmCount([]interface{}{"test", 1}, 1, IsPositive)
	ConfirmCount([]interface{}{1, 2}, 2, IsPositive)
}

func TestDensity(t *testing.T) {
	ConfirmDensity := func(o interface{}, r float64, f func(x interface{}) bool) {
		tol := 0.0001
		if d := Density(o, f); (d - r > tol) && (r - d < tol) {
			t.Fatalf("Density(%v, f) should be %v with a tolerance of %v but is %v", o, r, tol, d)
		}
	}

	IsPositive := func(i interface{}) bool {
		if i, ok := i.(int); ok {
			return i > 0
		}
		return false
	}

	ConfirmDensity(nil, 0.0, IsPositive)
	ConfirmDensity([]int{}, 0.0, IsPositive)
	ConfirmDensity([]int{0}, 0.0, IsPositive)
	ConfirmDensity([]int{1}, 1.0, IsPositive)

	ConfirmDensity([]int{0, 1}, 0.5, IsPositive)
	ConfirmDensity([]int{1, 0}, 0.5, IsPositive)

	ConfirmDensity([]int{0, 0, 1}, 0.3333, IsPositive)
	ConfirmDensity([]int{0, 1, 0}, 0.3333, IsPositive)
	ConfirmDensity([]int{1, 0, 0}, 0.3333, IsPositive)

	ConfirmDensity([]int{1, 0, 1}, 0.6666, IsPositive)
	ConfirmDensity([]int{1, 1, 0}, 0.6666, IsPositive)
	ConfirmDensity([]int{0, 1, 1}, 0.6666, IsPositive)

	ConfirmDensity([]int{0, 0, 0, 1}, 0.25, IsPositive)
	ConfirmDensity([]int{0, 0, 1, 0}, 0.25, IsPositive)
	ConfirmDensity([]int{0, 1, 0, 0}, 0.25, IsPositive)
	ConfirmDensity([]int{1, 0, 0, 0}, 0.25, IsPositive)

	ConfirmDensity([]int{1, 1, 0, 1}, 0.75, IsPositive)
	ConfirmDensity([]int{1, 0, 1, 1}, 0.75, IsPositive)
	ConfirmDensity([]int{0, 1, 1, 1}, 0.75, IsPositive)
	ConfirmDensity([]int{1, 1, 1, 0}, 0.75, IsPositive)
}

func TestDense(t *testing.T) {
	tol := 0.0001
	ConfirmDense := func(o interface{}, d float64, f func(interface{}) bool) {
		if !Dense(o, d, tol, f) {
			t.Fatalf("Dense(%v, %v, %v, f) should be true but is false", o, d, tol)
		}
	}

	RefuteDense := func(o interface{}, d float64, f func(interface{}) bool) {
		if Dense(o, d, tol, f) {
			t.Fatalf("Dense(%v, %v, %v, f) should be false but is true", o, d, tol)
		}
	}

	IsPositive := func(i interface{}) bool {
		if i, ok := i.(int); ok {
			return i > 0
		}
		return false
	}

	RefuteDense(nil, 0.0, IsPositive)
	RefuteDense(nil, 0.5, IsPositive)
	RefuteDense(nil, 1.0, IsPositive)

	RefuteDense([]int{}, 0.0, IsPositive)
	RefuteDense([]int{}, 0.5, IsPositive)
	RefuteDense([]int{}, 1.0, IsPositive)

	RefuteDense([]int{0}, 0.0, IsPositive)
	RefuteDense([]int{0}, 0.5, IsPositive)
	RefuteDense([]int{0}, 1.0, IsPositive)

	ConfirmDense([]int{0, 1}, 0.0, IsPositive)
	ConfirmDense([]int{0, 1}, 0.45, IsPositive)
	RefuteDense([]int{0, 1}, 0.55, IsPositive)
	RefuteDense([]int{0, 1}, 1.0, IsPositive)

	ConfirmDense([]int{0, 0, 1}, 0.0, IsPositive)
	RefuteDense([]int{0, 0, 1}, 0.5, IsPositive)
	RefuteDense([]int{0, 0, 1}, 1.0, IsPositive)

	ConfirmDense([]int{0, 1, 1}, 0.0, IsPositive)
	ConfirmDense([]int{0, 1, 1}, 0.5, IsPositive)
	RefuteDense([]int{0, 1, 1}, 1.0, IsPositive)

	ConfirmDense([]int{0, 0, 0, 1}, 0.0, IsPositive)
	RefuteDense([]int{0, 0, 0, 1}, 0.5, IsPositive)
	RefuteDense([]int{0, 0, 0, 1}, 1.0, IsPositive)

	ConfirmDense([]int{0, 0, 1, 1}, 0.0, IsPositive)
	ConfirmDense([]int{0, 0, 1, 1}, 0.45, IsPositive)
	RefuteDense([]int{0, 0, 1, 1}, 0.55, IsPositive)
	RefuteDense([]int{0, 0, 1, 1}, 1.0, IsPositive)

	ConfirmDense([]int{0, 1, 1, 1}, 0.0, IsPositive)
	ConfirmDense([]int{0, 1, 1, 1}, 0.5, IsPositive)
	RefuteDense([]int{0, 1, 1, 1}, 1.0, IsPositive)

	ConfirmDense([]int{1, 1, 1, 1}, 0.0, IsPositive)
	ConfirmDense([]int{1, 1, 1, 1}, 0.5, IsPositive)
	ConfirmDense([]int{1, 1, 1, 1}, 0.99, IsPositive)
	RefuteDense([]int{1, 1, 1, 1}, 1.0, IsPositive)
}

func TestMost(t *testing.T) {
	tol := 0.0001
	ConfirmMost := func(o interface{}, f func(interface{}) bool) {
		if !Most(o, tol, f) {
			t.Fatalf("Most(%v, %v, f) should be true but is false", o, tol)
		}
	}

	RefuteMost := func(o interface{}, f func(interface{}) bool) {
		if Most(o, tol, f) {
			t.Fatalf("Most(%v, %v, f) should be false but is true", o, tol)
		}
	}

	IsPositive := func(i interface{}) bool {
		if i, ok := i.(int); ok {
			return i > 0
		}
		return false
	}

	RefuteMost(nil, IsPositive)
	RefuteMost([]int{}, IsPositive)
	RefuteMost([]int{0}, IsPositive)

	RefuteMost([]int{0, 0}, IsPositive)
	RefuteMost([]int{0, 1}, IsPositive)
	ConfirmMost([]int{1, 1}, IsPositive)

	RefuteMost([]int{0, 0, 1}, IsPositive)
	ConfirmMost([]int{0, 1, 1}, IsPositive)
	ConfirmMost([]int{1, 1, 1}, IsPositive)

	RefuteMost([]int{0, 0, 0, 1}, IsPositive)
	RefuteMost([]int{0, 0, 1, 1}, IsPositive)
	ConfirmMost([]int{0, 1, 1, 1}, IsPositive)
	ConfirmMost([]int{1, 1, 1, 1}, IsPositive)
}

func TestReverse(t *testing.T) {
	ConfirmReverse := func(o, r interface{}) {
		Reverse(o)
		if !Equal(o, r) {
			t.Fatalf("Reverse(o) should be %v but is %v", r, o)
		}
	}
	ConfirmReverse(SList(0, 1, 2, 3, 4, 5), SList(5, 4, 3, 2, 1, 0))
	ConfirmReverse(List(0, 1, 2, 3, 4, 5), List(5, 4, 3, 2, 1, 0))
	ConfirmReverse([]int{0, 1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1, 0})
}

func TestDepth(t *testing.T) {
	ConfirmDepth := func(o interface{}, d int) {
		if x := Depth(o); d != x {
			t.Fatalf("Depth(%v) should be %v but is %v", o, d, x)
		}
	}

	ConfirmDepth(nil, 0)
	ConfirmDepth(1, 0)
	ConfirmDepth([]int{}, 0)
	ConfirmDepth([]int{0, 1, 2}, 1)


	ConfirmDepth(List(4, 3, 2, 1), 1)
	ConfirmDepth(List(	5,
						List(4, 3),
						2), 2)
	ConfirmDepth(List(	6,
						List(	5,
								List(4, 3, 2)),
						1), 3)
	ConfirmDepth(List(	7,
						List(	6,
								List(	5,
										4,
										List(3, 2),
										1)),
								0), 4)
	ConfirmDepth(List(	8,
						List(	7,
								List(	6,
										5,
										List(4, 3),
								2)),
								List(	1,
										List(0, -1))), 4)
	ConfirmDepth(List(	9,
						List(	8,
								List(	7,
										List(	6, 5)),
										List(	4,
												3,
												List(2, 1),
												0))), 4)
	ConfirmDepth(List(	'A',
						List(	9,
								SList(	8,
										SList(	7, 6 )),
								List(	5,
										4,
										List(3, 2),
										1))), 4)
	ConfirmDepth(List(	'B',
						List(	'A',
								SList(	9,
										SList(	8,
												SList( 7, 6 ))),
								List(	5,
										4,
										List(3, 2),
										1))), 5)

	ConfirmDepth(Loop(	4, 3, 2, 1), 1)
	ConfirmDepth(Loop(	5,
						Loop(4, 3),
						2), 2)
	ConfirmDepth(Loop(	6,
						Loop(	5,
								Loop(4, 3, 2)),
						1), 3)
	ConfirmDepth(Loop(	7,
						Loop(	6,
								Loop(	5,
										4,
										Loop(3, 2),
										1)),
								0), 4)
	ConfirmDepth(Loop(	8,
						Loop(	7,
								Loop(	6,
										5,
										Loop(4, 3),
								2)),
								Loop(	1,
										Loop(0, -1))), 4)
	ConfirmDepth(Loop(	9,
						Loop(	8,
								Loop(	7,
										Loop(	6, 5)),
										Loop(	4,
												3,
												Loop(2, 1),
												0))), 4)
	ConfirmDepth(Loop(	'A',
						Loop(	9,
								SList(	8,
										SList(	7, 6 )),
								Loop(	5,
										4,
										Loop(3, 2),
										1))), 4)
	ConfirmDepth(Loop(	'B',
						Loop(	'A',
								SList(	9,
										SList(	8,
												SList( 7, 6 ))),
								Loop(	5,
										4,
										Loop(3, 2),
										1))), 5)
}

/*
func TestFlatten(t *testing.T) { t.Fatal() }
func TestAppend(t *testing.T) { t.Fatal() }
func TestAppendContainer(t *testing.T) { t.Fatal() }
func TestPrepend(t *testing.T) { t.Fatal() }
func TestPrependContainer(t *testing.T) { t.Fatal() }
*/

func TestBlockCopy(t *testing.T) {
	ConfirmBlockCopy := func(i interface{}, d, s, n int, r interface{}) {
		BlockCopy(i, d, s, n)
		if !Equal(i, r) {
			t.Fatalf("BlockCopy(i, %v, %v, %v) should be %v but is %v", d, s, n, r, i)
		}
	}

	//	SList() returns a Slice which supports the Blitter interface
	ConfirmBlockCopy(SList(0, 1, 2, 3, 4, 5), 0, 1, -1, SList(0, 1, 2, 3, 4, 5))
	ConfirmBlockCopy(SList(0, 1, 2, 3, 4, 5), 0, 1, 0, SList(0, 1, 2, 3, 4, 5))
	ConfirmBlockCopy(SList(0, 1, 2, 3, 4, 5), 1, 0, 0, SList(0, 1, 2, 3, 4, 5))

	ConfirmBlockCopy(SList(0, 1, 2, 3, 4, 5), -1, 0, 3, SList(0, 1, 2, 3, 4, 5))
	ConfirmBlockCopy(SList(0, 1, 2, 3, 4, 5), 0, -1, 3, SList(0, 1, 2, 3, 4, 5))

	ConfirmBlockCopy(SList(0, 1, 2, 3, 4, 5), 3, 0, 3, SList(0, 1, 2, 0, 1, 2))
	ConfirmBlockCopy(SList(0, 1, 2, 3, 4, 5), 0, 3, 3, SList(3, 4, 5, 3, 4, 5))
	ConfirmBlockCopy(SList(0, 1, 2, 3, 4, 5), 0, 0, 3, SList(0, 1, 2, 3, 4, 5))

	ConfirmBlockCopy(SList(0, 1, 2, 3, 4, 5), 3, 0, 4, SList(0, 1, 2, 0, 1, 2))
	ConfirmBlockCopy(SList(0, 1, 2, 3, 4, 5), 0, 3, 4, SList(3, 4, 5, 3, 4, 5))

	//	List() returns a LinearList which supports the Indexable interface
	ConfirmBlockCopy(List(0, 1, 2, 3, 4, 5), 0, 1, -1, List(0, 1, 2, 3, 4, 5))
	ConfirmBlockCopy(List(0, 1, 2, 3, 4, 5), 0, 1, 0, List(0, 1, 2, 3, 4, 5))
	ConfirmBlockCopy(List(0, 1, 2, 3, 4, 5), 1, 0, 0, List(0, 1, 2, 3, 4, 5))

	ConfirmBlockCopy(List(0, 1, 2, 3, 4, 5), -1, 0, 3, List(0, 1, 2, 3, 4, 5))
	ConfirmBlockCopy(List(0, 1, 2, 3, 4, 5), 0, -1, 3, List(0, 1, 2, 3, 4, 5))

	ConfirmBlockCopy(List(0, 1, 2, 3, 4, 5), 3, 0, 3, List(0, 1, 2, 0, 1, 2))
	ConfirmBlockCopy(List(0, 1, 2, 3, 4, 5), 0, 3, 3, List(3, 4, 5, 3, 4, 5))
	ConfirmBlockCopy(List(0, 1, 2, 3, 4, 5), 0, 0, 3, List(0, 1, 2, 3, 4, 5))

	ConfirmBlockCopy(List(0, 1, 2, 3, 4, 5), 3, 0, 4, List(0, 1, 2, 0, 1, 2))
	ConfirmBlockCopy(List(0, 1, 2, 3, 4, 5), 0, 3, 4, List(3, 4, 5, 3, 4, 5))

	//	[]int{} slices are handled using reflection
	ConfirmBlockCopy([]int{0, 1, 2, 3, 4, 5}, 0, 1, -1, []int{0, 1, 2, 3, 4, 5})
	ConfirmBlockCopy([]int{0, 1, 2, 3, 4, 5}, 0, 1, 0, []int{0, 1, 2, 3, 4, 5})
	ConfirmBlockCopy([]int{0, 1, 2, 3, 4, 5}, 1, 0, 0, []int{0, 1, 2, 3, 4, 5})

	ConfirmBlockCopy([]int{0, 1, 2, 3, 4, 5}, -1, 0, 3, []int{0, 1, 2, 3, 4, 5})
	ConfirmBlockCopy([]int{0, 1, 2, 3, 4, 5}, 0, -1, 3, []int{0, 1, 2, 3, 4, 5})

	ConfirmBlockCopy([]int{0, 1, 2, 3, 4, 5}, 3, 0, 3, []int{0, 1, 2, 0, 1, 2})
	ConfirmBlockCopy([]int{0, 1, 2, 3, 4, 5}, 0, 3, 3, []int{3, 4, 5, 3, 4, 5})
	ConfirmBlockCopy([]int{0, 1, 2, 3, 4, 5}, 0, 0, 3, []int{0, 1, 2, 3, 4, 5})

	ConfirmBlockCopy([]int{0, 1, 2, 3, 4, 5}, 3, 0, 4, []int{0, 1, 2, 0, 1, 2})
	ConfirmBlockCopy([]int{0, 1, 2, 3, 4, 5}, 0, 3, 4, []int{3, 4, 5, 3, 4, 5})
}

func TestBlockClear(t *testing.T) {
	ConfirmBlockClear := func(i interface{}, d, n int, r interface{}) {
		BlockClear(i, d, n)
		if !Equal(i, r) {
			t.Fatalf("BlockClear(i, %v, %v, %v) should be %v but is %v", d, n, r, i)
		}
	}

	//	SList() returns a Slice which supports the Blitter interface
	ConfirmBlockClear(SList(0, 1, 2, 3, 4, 5), 3, 3, SList(0, 1, 2, nil, nil, nil))
	ConfirmBlockClear(SList(0, 1, 2, 3, 4, 5), 0, 3, SList(nil, nil, nil, 3, 4, 5))

	//	List() returns a LinearList which supports the Indexable interface
	ConfirmBlockClear(List(0, 1, 2, 3, 4, 5), 3, 3, List(0, 1, 2, nil, nil, nil))
	ConfirmBlockClear(List(0, 1, 2, 3, 4, 5), 0, 3, List(nil, nil, nil, 3, 4, 5))

	//	[]int{} slices are handled using reflection
	ConfirmBlockClear([]int{0, 1, 2, 3, 4, 5}, 3, 3, []int{0, 1, 2, 0, 0, 0})
	ConfirmBlockClear([]int{0, 1, 2, 3, 4, 5}, 0, 3, []int{0, 0, 0, 3, 4, 5})
}

func TestReallocate(t *testing.T) {
	ConfirmReallocate := func(i interface{}, l, c int, r interface{}) {
		if x := Reallocate(i, l, c); !Equal(x, r) {
			t.Fatalf("Reallocate(%v, %v, %v) should be %v but is %v", i, l, c, r, x)
		}
	}
	ConfirmReallocate([]int{0, 1, 2}, 2, 3, []int{0, 1})
	ConfirmReallocate([]int{0, 1, 2}, 3, 3, []int{0, 1, 2})
	ConfirmReallocate([]int{0, 1, 2}, 4, 3, []int{0, 1, 2})
	ConfirmReallocate([]int{0, 1, 2}, 4, 4, []int{0, 1, 2, 0})

	ConfirmReallocate(SList(0, 1, 2), 2, 3, SList(0, 1))
	ConfirmReallocate(SList(0, 1, 2), 3, 3, SList(0, 1, 2))
	ConfirmReallocate(SList(0, 1, 2), 3, 4, SList(0, 1, 2))
	ConfirmReallocate(SList(0, 1, 2), 4, 4, SList(0, 1, 2, nil))
}

func TestExpand(t *testing.T) {
	ConfirmExpand := func(i interface{}, x, n int, r interface{}) {
		if y := Expand(i, x, n); !Equal(y, r) {
			t.Fatalf("Expand(%v, %v, %v) should be %v but is %v", i, x, n, r, y)
		}
	}

	ConfirmExpand([]int{}, 0, -1, []int{})

	ConfirmExpand([]int{}, -1, 1, []int{})
	ConfirmExpand([]int{}, 0, 1, []int{0})
	ConfirmExpand([]int{}, 1, 1, []int{})

	ConfirmExpand([]int{0, 1, 2}, 1, 3, []int{0, 0, 0, 0, 1, 2})
	ConfirmExpand([]int{0, 1, 2}, 3, 3, []int{0, 1, 2, 0, 0, 0})
	ConfirmExpand([]int{0, 1, 2}, 4, 3, []int{0, 1, 2})

	ConfirmExpand(List(0, 1, 2, 3, 4, 5), 1, 3, List(0, nil, nil, nil, 1, 2, 3, 4, 5))



//	ConfirmExpand(SList(0, 1, 2, 3, 4, 5), 1, 3, SList(0, nil, nil, nil, 1, 2, 3, 4, 5))

}


func TestFeed(t *testing.T) {
	ConfirmFeed := func(c, r interface{}, f func(interface{}) interface{}) {
		channel := make(chan interface{})
		Feed(c, channel, f)
		o := make([]interface{}, 0, 0)
		for i := Len(c); i > 0; i-- {
			o = append(o, <-channel)
		}
		switch {
		case Len(o) != Len(r):			t.Fatalf("Feed(%v, <-, f()) expected %v results but got %v results", c, Len(r), Len(o))
		case !Equal(o, r):				t.Fatalf("Feed(%v, <-, f()) should generate %v but generated %v", c, r, o)
		}
	}

	i := 0
	ConfirmFeed([]int{0, 1, 2}, []int{0, 1, 4}, func(x interface{}) (r interface{}) {
		r = i * x.(int)
		i++
		return
	})

	//	test cases for Feeder
}

func TestPipe(t *testing.T) {
	ConfirmPipe := func(c, r interface{}, f func(interface{}) interface{}) {
		o := make([]interface{}, 0, 0)
		for x := range Pipe(c, f) {
			o = append(o, x)
		}
		switch {
		case Len(o) != Len(r):			t.Fatalf("Pipe(%v, <-, f()) expected %v results but got %v results", c, Len(r), Len(o))
		case !Equal(o, r):				t.Fatalf("Pipe(%v, <-, f()) should generate %v but generated %v", c, r, o)
		}
	}

	i := 0
	ConfirmPipe([]int{0, 1, 2}, []int{0, 1, 4}, func(x interface{}) (r interface{}) {
		r = i * x.(int)
		i++
		return
	})

	//	test cases for Piper
}