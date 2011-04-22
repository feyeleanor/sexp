package sexp

import "testing"

func TestSCons(t *testing.T) {
	sxp := SCons(nil, nil)
	switch {
	case len(sxp) != 2:		t.Fatalf("SCons(nil nil) should allocate 2 cells, not %v cells", len(sxp))
	case sxp[0] != nil:		t.Fatalf("SCons(nil nil) element 0 should be nil and not %v", sxp[0])
	case sxp[1] != nil:		t.Fatalf("SCons(nil nil) element 1 should be nil and not %v", sxp[1])
	}

	sxp = SCons(1, nil)
	switch {
	case len(sxp) != 2:		t.Fatalf("SCons(1 nil) should allocate 2 cells, not %v cells", len(sxp))
	case sxp[0] != 1:		t.Fatalf("SCons(1 nil) element 0 should be 1 and not %v", sxp[0])
	case sxp[1] != nil:		t.Fatalf("SCons(1 nil) element 1 should be nil and not %v", sxp[1])
	}

	sxp = SCons(1, 2)
	switch {
	case len(sxp) != 2:		t.Fatalf("SCons(1 2) should allocate 2 cells, not %v cells", len(sxp))
	case sxp[0] != 1:		t.Fatalf("SCons(1 2) element 0 should be 1 and not %v", sxp[0])
	case sxp[1] != 2:		t.Fatalf("SCons(1 2) element 1 should be 2 and not %v", sxp[1])
	}

	sxp = SCons(1, 2, 3)
	switch {
	case len(sxp) != 3:		t.Fatalf("SCons(1 2 3) should allocate 3 cells, not %v cells", len(sxp))
	case sxp[0] != 1:		t.Fatalf("SCons(1 2 3) element 0 should be 1 and not %v", sxp[0])
	case sxp[1] != 2:		t.Fatalf("SCons(1 2 3) element 1 should be 2 and not %v", sxp[1])
	case sxp[2] != 3:		t.Fatalf("SCons(1 2 3) element 2 should be 3 and not %v", sxp[2])
	}

	sxp = SCons(1, SCons(10, 20), 3)
	rxp := SEXP{ 10, 20 }
	switch {
	case len(sxp) != 3:			t.Fatalf("SCons(1 (10 20) 3) should allocate 3 cells, not %v cells", len(sxp))
	case sxp[0] != 1:			t.Fatalf("SCons(1 (10 20) 3) element 0 should be 1 and not %v", sxp[0])
	case !rxp.Equal(sxp[1]):	t.Fatalf("SCons(1 (10 20) 3) element 1 should be (10 20) and not %v", sxp[1])
	case sxp[2] != 3:			t.Fatalf("SCons(1 (10 20) 3) element 2 should be 3 and not %v", sxp[2])
	}


	sxp = SCons(1, SCons(10, SCons(-10, -30)), 3)
	rxp = SEXP{ 10, SEXP{ -10, -30 } }
	switch {
	case len(sxp) != 3:			t.Fatalf("SCons(1 (10 20) 3) should allocate 3 cells, not %v cells", len(sxp))
	case sxp[0] != 1:			t.Fatalf("SCons(1 (10 20) 3) element 0 should be 1 and not %v", sxp[0])
	case !rxp.Equal(sxp[1]):	t.Fatalf("SCons(1 (10 20) 3) element 1 should be (10 20) and not %v", sxp[1])
	case sxp[2] != 3:			t.Fatalf("SCons(1 (10 20) 3) element 2 should be 3 and not %v", sxp[2])
	}
}

func TestString(t *testing.T) {
	FormatError := func(x, y interface{}) { t.Fatalf("%v erroneously serialised as %v", x, y) }
	sxp := SEXP{ 0 }
	if s := sxp.String(); s != "(0)" { FormatError("(0)", s) }

	sxp = SEXP{ 0, 1 }
	if s := sxp.String(); s != "(0 1)" { FormatError("(0 1)", s) }

	sxp = SEXP{ SEXP{ 0, 1 }, 1 }
	if s := sxp.String(); s != "((0 1) 1)" { FormatError("((0 1) 1)", s) }

	sxp = SEXP{ SEXP{ 0, 1 }, SEXP{ 0, 1 } }
	if s := sxp.String(); s != "((0 1) (0 1))" { FormatError("((0 1) (0 1))", s) }
}

func TestLen(t *testing.T) {
	sxp := SEXP{ 0 }
	if sxp.Len() != 1 { t.Fatalf("With 1 element in an SEXP the length should be 1 but is %v", sxp.Len()) }

	sxp = SEXP{ 0, 1 }
	if sxp.Len() != 2 { t.Fatalf("With 2 element in an SEXP the length should be 2 but is %v", sxp.Len()) }

	sxp = SEXP{ SEXP{ 0, 1 }, 2 }
	if sxp.Len() != 2 { t.Fatalf("With 1 nested SEXP the length should be 2 but is %v", sxp.Len()) }

	sxp = SCons(0, 1)
	if sxp.Len() != 2 { t.Fatalf("With 0 nested SCons cells the length should be 2 but is %v", sxp.Len()) }

	sxp = SCons(SCons(0, 1), 2)
	if sxp.Len() != 2 { t.Fatalf("With 1 nested SCons cells the length should be 2 but is %v", sxp.Len()) }

	sxp = SCons(0, 1, SCons(2, SCons(3, 4, 5)), SCons(6, 7, 8, 9))
	if sxp.Len() != 4 { t.Fatalf("With 2 nested SCons cells the length should be 3 but is %v", sxp.Len()) }

	sxp = SCons(0, 1, SCons(2, SCons(3, 4, 5)), sxp, SCons(6, 7, 8, 9))
	if sxp.Len() != 5 { t.Fatalf("With 2 nested SCons cells plus recursion the length should be 5 but is %v", sxp.Len()) }
}

func TestDepth(t *testing.T) {
	sxp := SEXP{ 0, 1 }
	if sxp.Depth() != 0 { t.Fatalf("With 0 nested SEXP cells the depth should be 0 but is %v", sxp.Depth()) }

	sxp = SEXP{ SEXP{ 0, 1 }, 2 }
	if sxp.Depth() != 1 { t.Fatalf("With 1 nested SEXP cells the depth should be 1 but is %v", sxp.Depth()) }

	sxp = SCons(0, SCons(1, 2))
	if sxp.Depth() != 1 { t.Fatalf("With 1 nested SCons cells the depth should be 1 but is %v", sxp.Depth()) }

	sxp = SCons(0, 1,
				SCons(	2,
						SCons(3, 4, 5)	))
	if sxp.Depth() != 2 { t.Fatalf("With 2 nested SCons cells the depth should be 2 but is %v", sxp.Depth()) }

	sxp = SCons(0, 1,
				SCons(	2,
						SCons(3, 4, 5)	),
				SCons(	6,
						SCons(	7,
								SCons(	8,
										SCons(9, 0)	))),
				SCons(	2,
						SCons(3, 4, 5)	))
	if sxp.Depth() != 4 { t.Fatalf("With 4 nested SCons cells the depth should be 4 but is %v", sxp.Depth()) }

	rxp := SCons(0, sxp, sxp)
	if rxp.Depth() != 5 { t.Fatalf("With 5 nested SCons cells the depth should be 5 but is %v", rxp.Depth()) }

	sxp = SCons(rxp, sxp)
	if sxp.Depth() != 6 { t.Fatalf("With 6 nested SCons cells and circular references the depth should be 6 but is %v", sxp.Depth()) }

	t.Log("Need tests for circular recursive SEXP")
}

func TestReverse(t *testing.T) {
	sxp := SCons(1, 2, 3, 4, 5)
	rxp := SCons(5, 4, 3, 2, 1)
	sxp.Reverse()
	if !rxp.Equal(sxp) { t.Fatalf("Reversal failed: %v", sxp) }
}

func TestFlatten(t *testing.T) {
	sxp := SCons(1, 2, SCons(3, SCons(4, 5), SCons(6, SCons(7, 8, 9), SCons(10, 11))))
	rxp := SCons(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	sxp.Flatten()
	if !rxp.Equal(sxp) { t.Fatalf("Flatten failed: %v", sxp) }

	fxp := SCons(1, 2, sxp, 3, 4, 5, 6, 7, 8, 9, 10, 11, sxp)
	rxp = SCons(1, 2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 3, 4, 5, 6, 7, 8, 9, 10, 11, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	sxp = SCons(1, 2, sxp, SCons(3, SCons(4, 5), SCons(6, SCons(7, 8, 9), SCons(10, 11), sxp)))
	sxp.Flatten()
	switch {
	case !rxp.Equal(sxp):						t.Fatalf("Flatten failed with explicit expansions: %v", sxp)
	case !sxp.Equal(fxp.flatten(make(memo))):	t.Fatalf("Flatten failed with flattened expansions: %v", sxp)
	}
}

func TestCar(t *testing.T) {
	sxp := SCons(1, 2, 3)
	if h := sxp.Car(); h != 1 { t.Fatalf("head should be 1 but is %v", h) }

	c := SCons(10, 20)
	sxp = SCons(c, 2, 3)
	if h := sxp.Car(); !c.Equal(h) { t.Fatalf("head should be (10 20) but is %v", h) }
}

func TestCaar(t *testing.T) {
	sxp := SCons(1, 2)
	if h := sxp.Caar(); h != nil { t.Fatalf("head should be nil but is %v", h) }

	sxp = SCons(1, 2, 3)
	if h := sxp.Caar(); h != nil { t.Fatalf("head should be nil but is %v", h) }

	sxp = SCons(SCons(10, 20), 2, 3)
	if h := sxp.Caar(); h != 10 { t.Fatalf("head should be 10 but is %v", h) }

	sxp = SCons(SCons(SCons(10, 20), 20), 2, 3)
	if h := sxp.Caar(); !SCons(10, 20).Equal(h) { t.Fatalf("head should be (10 20) but is %v", h) }
}

func TestCdr(t *testing.T) {
	sxp := SCons(1, 2, 3)
	rxp := SEXP{ 2, 3 }
	if r := sxp.Cdr(); !r.Equal(rxp) { t.Fatalf("tail should be %v but is %v", rxp, r) }
}

func TestCddr(t *testing.T) {
	sxp := SCons(1, 2, 3)
	rxp := SEXP{ 3 }
	if r := sxp.Cddr(); !r.Equal(rxp) { t.Fatalf("tail should be %v but is %v", rxp, r) }

	sxp = SCons(1, 2, SCons(10, 20))
	rxp = SCons(10, 20)
	if r := sxp.Cddr(); !r.Equal(rxp) { t.Fatalf("tail should be %v but is %v", rxp, r) }

	sxp = SCons(1, SCons(10, 20))
	rxp = SEXP{ 20 }
	if r := sxp.Cddr(); !r.Equal(rxp) { t.Fatalf("tail should be %v but is %v", rxp, r) }
}

func TestRplaca(t *testing.T) {
	t.Log("Write Tests")
}

func TestRplacd(t *testing.T) {
	t.Log("Write Tests")
}