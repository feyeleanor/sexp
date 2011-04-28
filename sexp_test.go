package sexp

import "testing"

func TestSList(t *testing.T) {
	sxp := SList(nil, nil)
	switch {
	case len(sxp) != 2:		t.Fatalf("SList(nil nil) should allocate 2 cells, not %v cells", len(sxp))
	case sxp[0] != nil:		t.Fatalf("SList(nil nil) element 0 should be nil and not %v", sxp[0])
	case sxp[1] != nil:		t.Fatalf("SList(nil nil) element 1 should be nil and not %v", sxp[1])
	}

	sxp = SList(1, nil)
	switch {
	case len(sxp) != 2:		t.Fatalf("SList(1 nil) should allocate 2 cells, not %v cells", len(sxp))
	case sxp[0] != 1:		t.Fatalf("SList(1 nil) element 0 should be 1 and not %v", sxp[0])
	case sxp[1] != nil:		t.Fatalf("SList(1 nil) element 1 should be nil and not %v", sxp[1])
	}

	sxp = SList(1, 2)
	switch {
	case len(sxp) != 2:		t.Fatalf("SList(1 2) should allocate 2 cells, not %v cells", len(sxp))
	case sxp[0] != 1:		t.Fatalf("SList(1 2) element 0 should be 1 and not %v", sxp[0])
	case sxp[1] != 2:		t.Fatalf("SList(1 2) element 1 should be 2 and not %v", sxp[1])
	}

	sxp = SList(1, 2, 3)
	switch {
	case len(sxp) != 3:		t.Fatalf("SList(1 2 3) should allocate 3 cells, not %v cells", len(sxp))
	case sxp[0] != 1:		t.Fatalf("SList(1 2 3) element 0 should be 1 and not %v", sxp[0])
	case sxp[1] != 2:		t.Fatalf("SList(1 2 3) element 1 should be 2 and not %v", sxp[1])
	case sxp[2] != 3:		t.Fatalf("SList(1 2 3) element 2 should be 3 and not %v", sxp[2])
	}

	sxp = SList(1, SList(10, 20), 3)
	rxp := Slice{ 10, 20 }
	switch {
	case len(sxp) != 3:			t.Fatalf("SList(1 (10 20) 3) should allocate 3 cells, not %v cells", len(sxp))
	case sxp[0] != 1:			t.Fatalf("SList(1 (10 20) 3) element 0 should be 1 and not %v", sxp[0])
	case !rxp.Equal(sxp[1]):	t.Fatalf("SList(1 (10 20) 3) element 1 should be (10 20) and not %v", sxp[1])
	case sxp[2] != 3:			t.Fatalf("SList(1 (10 20) 3) element 2 should be 3 and not %v", sxp[2])
	}


	sxp = SList(1, SList(10, SList(-10, -30)), 3)
	rxp = Slice{ 10, Slice{ -10, -30 } }
	switch {
	case len(sxp) != 3:			t.Fatalf("SList(1 (10 20) 3) should allocate 3 cells, not %v cells", len(sxp))
	case sxp[0] != 1:			t.Fatalf("SList(1 (10 20) 3) element 0 should be 1 and not %v", sxp[0])
	case !rxp.Equal(sxp[1]):	t.Fatalf("SList(1 (10 20) 3) element 1 should be (10 20) and not %v", sxp[1])
	case sxp[2] != 3:			t.Fatalf("SList(1 (10 20) 3) element 2 should be 3 and not %v", sxp[2])
	}
}

func TestSliceString(t *testing.T) {
	FormatError := func(x, y interface{}) { t.Fatalf("%v erroneously serialised as %v", x, y) }
	sxp := Slice{ 0 }
	if s := sxp.String(); s != "(0)" { FormatError("(0)", s) }

	sxp = Slice{ 0, 1 }
	if s := sxp.String(); s != "(0 1)" { FormatError("(0 1)", s) }

	sxp = Slice{ Slice{ 0, 1 }, 1 }
	if s := sxp.String(); s != "((0 1) 1)" { FormatError("((0 1) 1)", s) }

	sxp = Slice{ Slice{ 0, 1 }, Slice{ 0, 1 } }
	if s := sxp.String(); s != "((0 1) (0 1))" { FormatError("((0 1) (0 1))", s) }
}

func TestSliceLen(t *testing.T) {
	sxp := Slice{ 0 }
	if sxp.Len() != 1 { t.Fatalf("With 1 element in an Slice the length should be 1 but is %v", sxp.Len()) }

	sxp = Slice{ 0, 1 }
	if sxp.Len() != 2 { t.Fatalf("With 2 element in an Slice the length should be 2 but is %v", sxp.Len()) }

	sxp = Slice{ Slice{ 0, 1 }, 2 }
	if sxp.Len() != 2 { t.Fatalf("With 1 nested Slice the length should be 2 but is %v", sxp.Len()) }

	sxp = SList(0, 1)
	if sxp.Len() != 2 { t.Fatalf("With 0 nested SList cells the length should be 2 but is %v", sxp.Len()) }

	sxp = SList(SList(0, 1), 2)
	if sxp.Len() != 2 { t.Fatalf("With 1 nested SList cells the length should be 2 but is %v", sxp.Len()) }

	sxp = SList(0, 1, SList(2, SList(3, 4, 5)), SList(6, 7, 8, 9))
	if sxp.Len() != 4 { t.Fatalf("With 2 nested SList cells the length should be 3 but is %v", sxp.Len()) }

	sxp = SList(0, 1, SList(2, SList(3, 4, 5)), sxp, SList(6, 7, 8, 9))
	if sxp.Len() != 5 { t.Fatalf("With 2 nested SList cells plus recursion the length should be 5 but is %v", sxp.Len()) }
}

func TestSliceDepth(t *testing.T) {
	sxp := Slice{ 0, 1 }
	if sxp.Depth() != 0 { t.Fatalf("With 0 nested Slice cells the depth should be 0 but is %v", sxp.Depth()) }

	sxp = Slice{ Slice{ 0, 1 }, 2 }
	if sxp.Depth() != 1 { t.Fatalf("With 1 nested Slice cells the depth should be 1 but is %v", sxp.Depth()) }

	sxp = SList(0, SList(1, 2))
	if sxp.Depth() != 1 { t.Fatalf("With 1 nested SList cells the depth should be 1 but is %v", sxp.Depth()) }

	sxp = SList(0, 1,
				SList(	2,
						SList(3, 4, 5)	))
	if sxp.Depth() != 2 { t.Fatalf("With 2 nested SList cells the depth should be 2 but is %v", sxp.Depth()) }

	sxp = SList(0, 1,
				SList(	2,
						SList(3, 4, 5)	),
				SList(	6,
						SList(	7,
								SList(	8,
										SList(9, 0)	))),
				SList(	2,
						SList(3, 4, 5)	))
	if sxp.Depth() != 4 { t.Fatalf("With 4 nested SList cells the depth should be 4 but is %v", sxp.Depth()) }

	rxp := SList(0, sxp, sxp)
	if rxp.Depth() != 5 { t.Fatalf("With 5 nested SList cells the depth should be 5 but is %v", rxp.Depth()) }

	sxp = SList(rxp, sxp)
	if sxp.Depth() != 6 { t.Fatalf("With 6 nested SList cells and circular references the depth should be 6 but is %v", sxp.Depth()) }

	t.Log("Need tests for circular recursive Slice")
}

func TestSliceReverse(t *testing.T) {
	sxp := SList(1, 2, 3, 4, 5)
	rxp := SList(5, 4, 3, 2, 1)
	sxp.Reverse()
	if !rxp.Equal(sxp) { t.Fatalf("Reversal failed: %v", sxp) }
}

func TestSliceFlatten(t *testing.T) {
	sxp := SList(1, 2, SList(3, SList(4, 5), SList(6, SList(7, 8, 9), SList(10, 11))))
	rxp := SList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	sxp.Flatten()
	if !rxp.Equal(sxp) { t.Fatalf("Flatten failed: %v", sxp) }

	fxp := SList(1, 2, sxp, 3, 4, 5, 6, 7, 8, 9, 10, 11, sxp)
	rxp = SList(1, 2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 3, 4, 5, 6, 7, 8, 9, 10, 11, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	sxp = SList(1, 2, sxp, SList(3, SList(4, 5), SList(6, SList(7, 8, 9), SList(10, 11), sxp)))
	sxp.Flatten()
	switch {
	case !rxp.Equal(sxp):						t.Fatalf("Flatten failed with explicit expansions: %v", sxp)
	case !sxp.Equal(fxp.flatten(make(memo))):	t.Fatalf("Flatten failed with flattened expansions: %v", sxp)
	}
}

func TestSliceCar(t *testing.T) {
	sxp := SList(1, 2, 3)
	if h := sxp.Car(); h != 1 { t.Fatalf("head should be 1 but is %v", h) }

	c := SList(10, 20)
	sxp = SList(c, 2, 3)
	if h := sxp.Car(); !c.Equal(h) { t.Fatalf("head should be (10 20) but is %v", h) }
}

func TestSliceCaar(t *testing.T) {
	sxp := SList(1, 2)
	if h := sxp.Caar(); h != nil { t.Fatalf("head should be nil but is %v", h) }

	sxp = SList(1, 2, 3)
	if h := sxp.Caar(); h != nil { t.Fatalf("head should be nil but is %v", h) }

	sxp = SList(SList(10, 20), 2, 3)
	if h := sxp.Caar(); h != 10 { t.Fatalf("head should be 10 but is %v", h) }

	sxp = SList(SList(SList(10, 20), 20), 2, 3)
	if h := sxp.Caar(); !SList(10, 20).Equal(h) { t.Fatalf("head should be (10 20) but is %v", h) }
}

func TestSliceCdr(t *testing.T) {
	sxp := SList(1, 2, 3)
	rxp := Slice{ 2, 3 }
	if r := sxp.Cdr(); !r.Equal(rxp) { t.Fatalf("tail should be %v but is %v", rxp, r) }
}

func TestSliceCddr(t *testing.T) {
	sxp := SList(1, 2, 3)
	rxp := Slice{ 3 }
	if r := sxp.Cddr(); !r.Equal(rxp) { t.Fatalf("tail should be %v but is %v", rxp, r) }

	sxp = SList(1, 2, SList(10, 20))
	rxp = SList(10, 20)
	if r := sxp.Cddr(); !r.Equal(rxp) { t.Fatalf("tail should be %v but is %v", rxp, r) }

	sxp = SList(1, SList(10, 20))
	rxp = Slice{ 20 }
	if r := sxp.Cddr(); !r.Equal(rxp) { t.Fatalf("tail should be %v but is %v", rxp, r) }
}

func TestSliceRplaca(t *testing.T) {
	t.Log("Write Tests")
}

func TestSliceRplacd(t *testing.T) {
	t.Log("Write Tests")
}