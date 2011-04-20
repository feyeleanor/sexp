package sexp

import "testing"

func TestCons(t *testing.T) {
	sxp := Cons(nil, nil)
	switch {
	case len(sxp) != 2:		t.Fatalf("Cons(nil nil) should allocate 2 cells, not %v cells", len(sxp))
	case sxp[0] != nil:		t.Fatalf("Cons(nil nil) element 0 should be nil and not %v", sxp[0])
	case sxp[1] != nil:		t.Fatalf("Cons(nil nil) element 1 should be nil and not %v", sxp[1])
	}

	sxp = Cons(1, nil)
	switch {
	case len(sxp) != 2:		t.Fatalf("Cons(1 nil) should allocate 2 cells, not %v cells", len(sxp))
	case sxp[0] != 1:		t.Fatalf("Cons(1 nil) element 0 should be 1 and not %v", sxp[0])
	case sxp[1] != nil:		t.Fatalf("Cons(1 nil) element 1 should be nil and not %v", sxp[1])
	}

	sxp = Cons(1, 2)
	switch {
	case len(sxp) != 2:		t.Fatalf("Cons(1 2) should allocate 2 cells, not %v cells", len(sxp))
	case sxp[0] != 1:		t.Fatalf("Cons(1 2) element 0 should be 1 and not %v", sxp[0])
	case sxp[1] != 2:		t.Fatalf("Cons(1 2) element 1 should be 2 and not %v", sxp[1])
	}

	sxp = Cons(1, 2, 3)
	switch {
	case len(sxp) != 3:		t.Fatalf("Cons(1 2 3) should allocate 3 cells, not %v cells", len(sxp))
	case sxp[0] != 1:		t.Fatalf("Cons(1 2 3) element 0 should be 1 and not %v", sxp[0])
	case sxp[1] != 2:		t.Fatalf("Cons(1 2 3) element 1 should be 2 and not %v", sxp[1])
	case sxp[2] != 3:		t.Fatalf("Cons(1 2 3) element 2 should be 3 and not %v", sxp[2])
	}

	sxp = Cons(1, Cons(10, 20), 3)
	rxp := SExp{ 10, 20 }
	switch {
	case len(sxp) != 3:			t.Fatalf("Cons(1 (10 20) 3) should allocate 3 cells, not %v cells", len(sxp))
	case sxp[0] != 1:			t.Fatalf("Cons(1 (10 20) 3) element 0 should be 1 and not %v", sxp[0])
	case !rxp.Equal(sxp[1]):	t.Fatalf("Cons(1 (10 20) 3) element 1 should be (10 20) and not %v", sxp[1])
	case sxp[2] != 3:			t.Fatalf("Cons(1 (10 20) 3) element 2 should be 3 and not %v", sxp[2])
	}


	sxp = Cons(1, Cons(10, Cons(-10, -30)), 3)
	rxp = SExp{ 10, SExp{ -10, -30 } }
	switch {
	case len(sxp) != 3:			t.Fatalf("Cons(1 (10 20) 3) should allocate 3 cells, not %v cells", len(sxp))
	case sxp[0] != 1:			t.Fatalf("Cons(1 (10 20) 3) element 0 should be 1 and not %v", sxp[0])
	case !rxp.Equal(sxp[1]):	t.Fatalf("Cons(1 (10 20) 3) element 1 should be (10 20) and not %v", sxp[1])
	case sxp[2] != 3:			t.Fatalf("Cons(1 (10 20) 3) element 2 should be 3 and not %v", sxp[2])
	}
}

func TestString(t *testing.T) {
	FormatError := func(x, y interface{}) { t.Fatalf("%v erroneously serialised as %v", x, y) }
	sxp := SExp{ 0 }
	if s := sxp.String(); s != "(0)" { FormatError("(0)", s) }

	sxp = SExp{ 0, 1 }
	if s := sxp.String(); s != "(0 1)" { FormatError("(0 1)", s) }

	sxp = SExp{ SExp{ 0, 1 }, 1 }
	if s := sxp.String(); s != "((0 1) 1)" { FormatError("((0 1) 1)", s) }

	sxp = SExp{ SExp{ 0, 1 }, SExp{ 0, 1 } }
	if s := sxp.String(); s != "((0 1) (0 1))" { FormatError("((0 1) (0 1))", s) }
}

func TestLen(t *testing.T) {
	sxp := SExp{ 0 }
	if sxp.Len() != 1 { t.Fatalf("With 1 element in an SExp the length should be 1 but is %v", sxp.Len()) }

	sxp = Cons(0, 1)
	if sxp.Len() != 2 { t.Fatalf("With 0 nested Cons cells the length should be 2 but is %v", sxp.Len()) }

	sxp = Cons(Cons(0, 1), 2)
	if sxp.Len() != 3 { t.Fatalf("With 1 nested Cons cells the length should be 3 but is %v", sxp.Len()) }

	sxp = Cons(0, 1, Cons(2, Cons(3, 4, 5)), Cons(6, 7, 8, 9))
	if sxp.Len() != 10 { t.Fatalf("With 3 nested Cons cells the length should be 10 but is %v", sxp.Len()) }

	sxp = Cons(0, 1, Cons(2, Cons(3, 4, 5)), sxp, Cons(6, 7, 8, 9))
	if sxp.Len() != 20 { t.Fatalf("With 3 nested Cons cells plus recursion the length should be 20 but is %v", sxp.Len()) }

	t.Log("Need tests for circular recursive SExp")
}

func TestDepth(t *testing.T) {
	sxp := Cons(0, 1)
	if sxp.Depth() != 0 { t.Fatalf("With 0 nested Cons cells the depth should be 0 but is %v", sxp.Depth()) }

	sxp = Cons(0, Cons(1, 2))
	if sxp.Depth() != 1 { t.Fatalf("With 1 nested Cons cells the depth should be 1 but is %v", sxp.Depth()) }

	sxp = Cons(	0, 1,
				Cons(	2,
						Cons(3, 4, 5)	))
	if sxp.Depth() != 2 { t.Fatalf("With 2 nested Cons cells the depth should be 2 but is %v", sxp.Depth()) }

	sxp = Cons(	0, 1,
				Cons(	2,
						Cons(3, 4, 5)	),
				Cons(	6,
						Cons(	7,
								Cons(	8,
										Cons(9, 0)	))),
				Cons(	2,
						Cons(3, 4, 5)	))
	if sxp.Depth() != 4 { t.Fatalf("With 4 nested Cons cells the depth should be 4 but is %v", sxp.Depth()) }

	rxp := Cons(0, sxp, sxp)
	if rxp.Depth() != 5 { t.Fatalf("With 5 nested Cons cells the depth should be 5 but is %v", rxp.Depth()) }

	sxp = Cons(rxp, sxp)
	if sxp.Depth() != 6 { t.Fatalf("With 6 nested Cons cells and circular references the depth should be 6 but is %v", sxp.Depth()) }

	t.Log("Need tests for circular recursive SExp")
}

func TestBounds(t *testing.T) {
	sxp := Cons(0, 1)
	switch l, d := sxp.Bounds(); {
	case d != 0:		t.Fatalf("With 0 nested Cons cells the depth should be 0 but is %v", d)
	case l != 2:		t.Fatalf("With 0 nested Cons cells the length should be 2 but is %v", l)
	}

	sxp = Cons(0, Cons(1, 2))
	switch l, d := sxp.Bounds(); {
	case d != 1:		t.Fatalf("With 1 nested Cons cells the depth should be 1 but is %v", d)
	case l != 3:		t.Fatalf("With 1 nested Cons cells the length should be 3 but is %v", l)
	}

	sxp = Cons(	0, 1,
				Cons(	2,
						Cons(3, 4, 5)	))
	switch l, d := sxp.Bounds(); {
	case d != 2:		t.Fatalf("With 2 nested Cons cells the depth should be 2 but is %v", d)
	case l != 6:		t.Fatalf("With 2 nested Cons cells the length should be 6 but is %v", l)
	}

	sxp = Cons(	0, 1,
				Cons(	2,
						Cons(3, 4, 5)	),
					Cons(	6,
						Cons(	7,
								Cons(	8,
										Cons(9, 0)	))),
					Cons(	2,
						Cons(3, 4, 5)	))
	switch l, d := sxp.Bounds(); {
	case d != 4:		t.Fatalf("With 4 nested Cons cells the depth should be 4 but is %v", d)
	case l != 15:		t.Fatalf("With 4 nested Cons cells the length should be 15 but is %v", l)
	}

	rxp := Cons(0, sxp, sxp)
	switch l, d := rxp.Bounds(); {
	case d != 5:		t.Fatalf("With 5 nested Cons cells the depth should be 5 but is %v", d)
	case l != 31:		t.Fatalf("With 5 nested Cons cells the length should be 31 but is %v", l)
	}

	t.Log("Need tests for circular recursive SExp")
}

func TestReverse(t *testing.T) {
	sxp := Cons(1, 2, 3, 4, 5)
	rxp := Cons(5, 4, 3, 2, 1)
	sxp.Reverse()
	if !rxp.Equal(sxp) { t.Fatalf("Reversal failed: %v", sxp) }
}

func TestFlatten(t *testing.T) {
	sxp := Cons(1, 2, Cons(3, Cons(4, 5), Cons(6, Cons(7, 8, 9), Cons(10, 11))))
	rxp := Cons(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	sxp.Flatten()
	if !rxp.Equal(sxp) { t.Fatalf("Flatten failed: %v", sxp) }

	fxp := Cons(1, 2, sxp, 3, 4, 5, 6, 7, 8, 9, 10, 11, sxp)
	rxp = Cons(1, 2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 3, 4, 5, 6, 7, 8, 9, 10, 11, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	sxp = Cons(1, 2, sxp, Cons(3, Cons(4, 5), Cons(6, Cons(7, 8, 9), Cons(10, 11), sxp)))
	sxp.Flatten()
	switch {
	case !rxp.Equal(sxp):						t.Fatalf("Flatten failed with explicit expansions: %v", sxp)
	case !sxp.Equal(fxp.flatten(make(memo))):	t.Fatalf("Flatten failed with flattened expansions: %v", sxp)
	}
}

func TestCar(t *testing.T) {
	sxp := Cons(1, 2, 3)
	if h := sxp.Car(); h != 1 { t.Fatalf("head should be 1 but is %v", h) }

	c := Cons(10, 20)
	sxp = Cons(c, 2, 3)
	if h := sxp.Car(); !c.Equal(h) { t.Fatalf("head should be (10 20) but is %v", h) }
}

func TestCaar(t *testing.T) {
	sxp := Cons(1, 2, 3)
	if h := sxp.Caar(); h != nil { t.Fatalf("head should be nil but is %v", h) }

	sxp = Cons(Cons(10, 20), 2, 3)
	if h := sxp.Caar(); h != 10 { t.Fatalf("head should be 10 but is %v", h) }

	sxp = Cons(Cons(Cons(10, 20), 20), 2, 3)
	if h := sxp.Caar(); !Cons(10, 20).Equal(h) { t.Fatalf("head should be (10 20) but is %v", h) }
}

func TestCdr(t *testing.T) {
	sxp := Cons(1, 2, 3)
	rxp := SExp{ 2, 3 }
	if r := sxp.Cdr(); !r.Equal(rxp) { t.Fatalf("tail should be %v but is %v", rxp, r) }
}

func TestCddr(t *testing.T) {
	sxp := Cons(1, 2, 3)
	rxp := SExp{ 3 }
	if r := sxp.Cddr(); !r.Equal(rxp) { t.Fatalf("tail should be %v but is %v", rxp, r) }

	sxp = Cons(1, 2, Cons(10, 20))
	rxp = Cons(10, 20)
	if r := sxp.Cddr(); !r.Equal(rxp) { t.Fatalf("tail should be %v but is %v", rxp, r) }

	sxp = Cons(1, Cons(10, 20))
	rxp = SExp{ 20 }
	if r := sxp.Cddr(); !r.Equal(rxp) { t.Fatalf("tail should be %v but is %v", rxp, r) }
}

func TestRplaca(t *testing.T) {
	t.Log("Write Tests")
}

func TestRplacd(t *testing.T) {
	t.Log("Write Tests")
}