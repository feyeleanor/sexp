package sexp

import "testing"

func TestCons(t *testing.T) {
	sxp := Cons(nil)
	switch {
	case len(sxp) != 2:		t.Fatalf("Cons(nil) should allocate 2 cells, not %v cells", len(sxp))
	case sxp[0] != nil:		t.Fatalf("Cons(nil) element 0 should be nil and not %v", sxp[0])
	case sxp[1] != nil:		t.Fatalf("Cons(nil) element 1 should be nil and not %v", sxp[1])
	}

	sxp = Cons(1)
	switch {
	case len(sxp) != 2:		t.Fatalf("Cons(1) should allocate 2 cells, not %v cells", len(sxp))
	case sxp[0] != 1:		t.Fatalf("Cons(1) element 0 should be 1 and not %v", sxp[0])
	case sxp[1] != nil:		t.Fatalf("Cons(1) element 1 should be nil and not %v", sxp[1])
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

func TestReverse(t *testing.T) {
	sxp := Cons(1, 2, 3, 4, 5)
	rxp := Cons(5, 4, 3, 2, 1)
	sxp.Reverse()
	if !rxp.Equal(sxp) { t.Fatalf("Reversal failed: %v", sxp) }
}

func TestCar(t *testing.T) {
	sxp := Cons(1, 2, 3)
	if h := sxp.Car(); h != 1 { t.Fatalf("head should be 1 but is %v", h) }
}

func TestCaar(t *testing.T) {
	sxp := Cons(1, 2, 3)
	if h := sxp.Caar(); h != nil { t.Fatalf("head should be nil but is %v", h) }

	sxp = Cons(Cons(10, 20), 2, 3)
	if h := sxp.Caar(); h != 10 { t.Fatalf("head should be 10 but is %v", h) }
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
	t.Fatal()
}

func TestRplacd(t *testing.T) {
	t.Fatal()
}