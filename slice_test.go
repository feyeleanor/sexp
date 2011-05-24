package sexp

import "testing"

func TestSliceIsNil(t *testing.T) {
	ConfirmIsNil := func(s *Slice) {
		if !s.IsNil() {
			t.Fatalf("%v should be nil", s)
		}
	}
	RefuteIsNil := func(s *Slice) {
		if s.IsNil() {
			t.Fatalf("%v should not be nil", s)
		}
	}
	ConfirmIsNil((*Slice)(nil))
	ConfirmIsNil(&Slice{})
	ConfirmIsNil(SList())
	ConfirmIsNil((*Slice)(&[]interface{}{}))
	RefuteIsNil(&Slice{ []interface{}{} })
	RefuteIsNil(SList(&[]interface{}{}))
	RefuteIsNil(SList(nil))
	RefuteIsNil(SList(0, 1))
}

func TestSList(t *testing.T) {
	sxp := SList(nil, nil)
	switch {
	case sxp.Len() != 2:			t.Fatalf("SList(nil nil) should allocate 2 cells, not %v cells", sxp.Len())
	case sxp.At(0) != nil:			t.Fatalf("SList(nil nil) element 0 should be nil and not %v", sxp.At(0))
	case sxp.At(1) != nil:			t.Fatalf("SList(nil nil) element 1 should be nil and not %v", sxp.At(1))
	}

	sxp = SList(1, nil)
	switch {
	case sxp.Len() != 2:			t.Fatalf("SList(1 nil) should allocate 2 cells, not %v cells", sxp.Len())
	case sxp.At(0) != 1:			t.Fatalf("SList(1 nil) element 0 should be 1 and not %v", sxp.At(0))
	case sxp.At(1) != nil:			t.Fatalf("SList(1 nil) element 1 should be nil and not %v", sxp.At(1))
	}

	sxp = SList(1, 2)
	switch {
	case sxp.Len() != 2:			t.Fatalf("SList(1 2) should allocate 2 cells, not %v cells", sxp.Len())
	case sxp.At(0) != 1:			t.Fatalf("SList(1 2) element 0 should be 1 and not %v", sxp.At(0))
	case sxp.At(1) != 2:			t.Fatalf("SList(1 2) element 1 should be 2 and not %v", sxp.At(1))
	}

	sxp = SList(1, 2, 3)
	switch {
	case sxp.Len() != 3:			t.Fatalf("SList(1 2 3) should allocate 3 cells, not %v cells", sxp.Len())
	case sxp.At(0) != 1:			t.Fatalf("SList(1 2 3) element 0 should be 1 and not %v", sxp.At(0))
	case sxp.At(1) != 2:			t.Fatalf("SList(1 2 3) element 1 should be 2 and not %v", sxp.At(1))
	case sxp.At(2) != 3:			t.Fatalf("SList(1 2 3) element 2 should be 3 and not %v", sxp.At(2))
	}

	sxp = SList(1, SList(10, 20), 3)
	rxp := SList(10, 20)
	switch {
	case sxp.Len() != 3:			t.Fatalf("SList(1 (10 20) 3) should allocate 3 cells, not %v cells", sxp.Len())
	case sxp.At(0) != 1:			t.Fatalf("SList(1 (10 20) 3) element 0 should be 1 and not %v", sxp.At(0))
	case !rxp.Equal(sxp.At(1)):		t.Fatalf("SList(1 (10 20) 3) element 1 should be (10 20) and not %v", sxp.At(1))
	case sxp.At(2) != 3:			t.Fatalf("SList(1 (10 20) 3) element 2 should be 3 and not %v", sxp.At(2))
	}


	sxp = SList(1, SList(10, SList(-10, -30)), 3)
	rxp = SList(10, SList(-10, -30))
	switch {
	case sxp.Len() != 3:			t.Fatalf("SList(1 (10 20) 3) should allocate 3 cells, not %v cells", sxp.Len())
	case sxp.At(0) != 1:			t.Fatalf("SList(1 (10 20) 3) element 0 should be 1 and not %v", sxp.At(0))
	case !rxp.Equal(sxp.At(1)):		t.Fatalf("SList(1 (10 20) 3) element 1 should be (10 20) and not %v", sxp.At(1))
	case sxp.At(2) != 3:			t.Fatalf("SList(1 (10 20) 3) element 2 should be 3 and not %v", sxp.At(2))
	}
}

func TestSliceString(t *testing.T) {
	ConfirmString := func(s *Slice, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v erroneously serialised as '%v'", r, x)
		}
	}

	ConfirmString(SList(), "()")
	ConfirmString(SList(0), "(0)")
	ConfirmString(SList(0, 1), "(0 1)")
	ConfirmString(SList(SList(0, 1), 1), "((0 1) 1)")
	ConfirmString(SList(SList(0, 1), SList(0, 1)), "((0 1) (0 1))")
}

func TestSliceLen(t *testing.T) {
	ConfirmLength := func(s *Slice, i int) {
		if x := s.Len(); x != i {
			t.Fatalf("%v.Len() should be %v but is %v", s, i, x)
		}
	}
	
	ConfirmLength(SList(0), 1)
	ConfirmLength(SList(0, 1), 2)
	ConfirmLength(SList(SList(0, 1), 2), 2)
	ConfirmLength(SList(0, 1), 2)
	ConfirmLength(SList(SList(0, 1), 2), 2)

	sxp := SList(0, 1, SList(2, SList(3, 4, 5)), SList(6, 7, 8, 9))
	ConfirmLength(sxp, 4)
	ConfirmLength(SList(0, 1, SList(2, SList(3, 4, 5)), sxp, SList(6, 7, 8, 9)), 5)
}

func TestSliceEach(t *testing.T) {
	c := SList(0, 1, 2, 3, 4, 5, 6, 7, 8 ,9)
	count := 0
	c.Each(func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})
}

func TestSliceBlit(t *testing.T) {
	ConfirmBlit := func(s *Slice, destination, source, count int, r *Slice) {
		s.Blit(destination, source, count)
		if !r.Equal(s) {
			t.Fatalf("Blit(%v, %v, %v) should be %v but is %v", destination, source, count, r, s)
		}
	}

	ConfirmBlit(SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 0, 4, SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlit(SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 9, 9, 4, SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlit(SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 2, 4, SList(0, 1, 2, 3, 4, 2, 3, 4, 5, 9))
	ConfirmBlit(SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 2, 5, 4, SList(0, 1, 5, 6, 7, 8, 6, 7, 8, 9))
}

func TestSliceOverwrite(t *testing.T) {
	ConfirmOverwrite := func(s *Slice, offset int, v, r *Slice) {
		s.Overwrite(offset, *v)
		if !r.Equal(s) {
			t.Fatalf("Overwrite(%v, %v) should be %v but is %v", offset, v, r, s)
		}
	}

	ConfirmOverwrite(SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, SList(10, 9, 8, 7), SList(10, 9, 8, 7, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, SList(10, 9, 8, 7), SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, SList(11, 12, 13, 14), SList(0, 1, 2, 3, 4, 11, 12, 13, 14, 9))
}

func TestSliceReallocate(t *testing.T) {
	ConfirmReallocate := func(s *Slice, capacity int, r *Slice) {
		o := s.String()
		switch s.Reallocate(capacity); {
		case s == nil:				t.Fatalf("%v.Reallocate(%v) created a nil value for Slice", o, capacity)
		case s.Cap() != capacity:	t.Fatalf("%v.Reallocate(%v) capacity should be %v but is %v", o, capacity, s.Cap())
		case !r.Equal(s):			t.Fatalf("%v.Reallocate(%v) should be %v but is %v", o, r, s)
		}
	}

	ConfirmReallocate(SList(), 10, SList())
	ConfirmReallocate(SList(0, 1, 2, 3, 4), 10, SList(0, 1, 2, 3, 4))
	ConfirmReallocate(SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, SList(0, 1, 2, 3, 4))
}

func TestSliceDepth(t *testing.T) {
	ConfirmDepth := func(s *Slice, i int) {
		if x := s.Depth(); x != i {
			t.Fatalf("%v.Depth() should be %v but is %v", s, i, x)
		}
	}
	ConfirmDepth(SList(0, 1), 0)
	ConfirmDepth(SList(SList(0, 1), 2), 1)
	ConfirmDepth(SList(0, SList(1, 2)), 1)
	ConfirmDepth(SList(0, 1, SList(2, SList(3, 4, 5))), 2)

	sxp := SList(0, 1,
				SList(2, SList(3, 4, 5)),
				SList(6, SList(7, SList(8, SList(9, 0)))),
				SList(2, SList(3, 4, 5)))
	ConfirmDepth(sxp, 4)

	rxp := SList(0, sxp, sxp)
	ConfirmDepth(rxp, 5)
	ConfirmDepth(SList(rxp, sxp), 6)
	t.Log("Need tests for circular recursive Slice?")
}

func TestSliceReverse(t *testing.T) {
	sxp := SList(1, 2, 3, 4, 5)
	rxp := SList(5, 4, 3, 2, 1)
	sxp.Reverse()
	if !rxp.Equal(sxp) {
		t.Fatalf("Reversal failed: %v", sxp)
	}
}

func TestSliceAppend(t *testing.T) {
	ConfirmAppend := func(s *Slice, v interface{}, r *Slice) {
		s.Append(v)
		if !r.Equal(s) {
			t.Fatalf("Append(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppend(SList(), 0, SList(0))
	ConfirmAppend(SList(), SList(0, 1), SList(SList(0, 1)))
	ConfirmAppend(SList(0, 1, 2), SList(3, 4), SList(0, 1, 2, SList(3, 4)))
}

func TestSliceAppendSlice(t *testing.T) {
	ConfirmAppendSlice := func(s, v, r *Slice) {
		s.AppendSlice(*v)
		if !r.Equal(s) {
			t.Fatalf("AppendSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmAppendSlice(SList(), SList(0), SList(0))
	ConfirmAppendSlice(SList(), SList(0, 1), SList(0, 1))
	ConfirmAppendSlice(SList(0, 1, 2), SList(3, 4), SList(0, 1, 2, 3, 4))
}

func TestSlicePrepend(t *testing.T) {
	ConfirmPrepend := func(s *Slice, v interface{}, r *Slice) {
		s.Prepend(v)
		if !r.Equal(s) {
			t.Fatalf("Prepend(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrepend(SList(), 0, SList(0))
	ConfirmPrepend(SList(0), 1, SList(1, 0))
	ConfirmPrepend(SList(0, 1, 2), SList(3, 4), SList(SList(3, 4), 0, 1, 2))
}

func TestSlicePrependSlice(t *testing.T) {
	ConfirmPrependSlice := func(s, v, r *Slice) {
		s.PrependSlice(*v)
		if !r.Equal(s) {
			t.Fatalf("PrependSlice(%v) should be %v but is %v", v, r, s)
		}
	}

	ConfirmPrependSlice(SList(), SList(0), SList(0))
	ConfirmPrependSlice(SList(), SList(0, 1), SList(0, 1))
	ConfirmPrependSlice(SList(0, 1, 2), SList(3, 4), SList(3, 4, 0, 1, 2))
}

func TestSliceRepeat(t *testing.T) {
	ConfirmRepeat := func(s *Slice, count int, r *Slice) {
		x := s.Repeat(count)
		if !x.Equal(r) {
			t.Fatalf("%v.Repeat(%v) should be %v but is %v", s, count, r, x)
		}
	}

	ConfirmRepeat(SList(), 5, SList())
	ConfirmRepeat(SList(0), 1, SList(0))
	ConfirmRepeat(SList(0), 2, SList(0, 0))
	ConfirmRepeat(SList(0), 3, SList(0, 0, 0))
	ConfirmRepeat(SList(0), 4, SList(0, 0, 0, 0))
	ConfirmRepeat(SList(0), 5, SList(0, 0, 0, 0, 0))
}

func TestSliceFlatten(t *testing.T) {
	ConfirmFlatten := func(s, r *Slice) {
		s.Flatten()
		if !s.Equal(r) {
			t.Fatalf("%v should be %v", s, r)
		}
	}
	ConfirmFlatten(SList(), SList())
	ConfirmFlatten(SList(1), SList(1))
	ConfirmFlatten(SList(1, SList(2)), SList(1, 2))
	ConfirmFlatten(SList(1, SList(2, SList(3))), SList(1, 2, 3))
	ConfirmFlatten(SList(1, 2, SList(3, SList(4, 5), SList(6, SList(7, 8, 9), SList(10, 11)))), SList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11))

	ConfirmFlatten(SList(0, List(1, 2, SList(3, 4))), SList(0, List(1, 2, SList(3, 4))))
	ConfirmFlatten(SList(0, List(1, 2, List(3, 4))), SList(0, List(1, 2, 3, 4)))

	ConfirmFlatten(SList(0, Loop(1, 2)), SList(0, Loop(1, 2)))
	ConfirmFlatten(SList(0, List(1, Loop(2, 3))), SList(0, List(1, Loop(2, 3))))

	ConfirmFlatten(SList(0, List(1, 2, Loop(3, 4))), SList(0, List(1, 2, Loop(3, 4))))
	ConfirmFlatten(SList(3, 4, SList(5, 6, 7)), SList(3, 4, 5, 6, 7))
	ConfirmFlatten(SList(0, Loop(1, 2, SList(3, 4, SList(5, 6, 7)))), SList(0, Loop(1, 2, SList(3, 4, 5, 6, 7))))

	sxp := SList(1, 2, SList(3, SList(4, 5), SList(6, SList(7, 8, 9), SList(10, 11))))
	rxp := SList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	sxp.Flatten()
	if !rxp.Equal(sxp) {
		t.Fatalf("Flatten failed: %v", sxp)
	}

	rxp = SList(1, 2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 3, 4, 5, 6, 7, 8, 9, 10, 11, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	sxp = SList(1, 2, sxp, SList(3, SList(4, 5), SList(6, SList(7, 8, 9), SList(10, 11), sxp)))
	sxp.Flatten()
	if !rxp.Equal(sxp) {
		t.Fatalf("Flatten failed with explicit expansions: %v", sxp)
	}
}

func TestSliceCar(t *testing.T) {
	ConfirmCar := func(s *Slice, r interface{}) {
		var ok bool
		n := s.Car()
		switch n := n.(type) {
		case Equatable:		ok = n.Equal(r)
		default:			ok = n == r
		}
		if !ok {
			t.Fatalf("head should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCar(SList(1, 2, 3), 1)
	ConfirmCar(SList(SList(10, 20), 2, 3), SList(10, 20))
}

func TestSliceCaar(t *testing.T) {
	ConfirmCaar := func(s *Slice, r interface{}) {
		var ok bool
		n := s.Caar()
		switch n := n.(type) {
		case Equatable:		ok = n.Equal(r)
		default:			ok = n == r
		}
		if !ok {
			t.Fatalf("head should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCaar(SList(1, 2), nil)
	ConfirmCaar(SList(1, 2, 3), nil)
	ConfirmCaar(SList(SList(10, 20), 2, 3), 10)
	ConfirmCaar(SList(SList(SList(10, 20), 20), 2, 3), SList(10, 20))
}

func TestSliceCdr(t *testing.T) {
	ConfirmCdr := func(s, r *Slice) {
		if n := s.Cdr(); !n.Equal(r) {
			t.Fatalf("tail should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCdr(SList(1, 2, 3), SList(2, 3))
}

func TestSliceCddr(t *testing.T) {
	ConfirmCddr := func(s, r *Slice) {
		if n := s.Cddr(); !n.Equal(r) {
			t.Fatalf("tail should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCddr(SList(1, 2, 3), SList(3))
	ConfirmCddr(SList(1, 2, SList(10, 20)), SList(10, 20))
	ConfirmCddr(SList(1, SList(10, 20)), SList(20))
}

func TestSliceRplaca(t *testing.T) {
	ConfirmRplaca := func(s *Slice, v interface{}, r *Slice) {
		s.Rplaca(v)
		if !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplaca(SList(1, 2, 3, 4, 5), 0, SList(0, 2, 3, 4, 5))
	ConfirmRplaca(SList(1, 2, 3, 4, 5), SList(1, 2, 3), SList(SList(1, 2, 3), 2, 3, 4, 5))
}

func TestSliceRplacd(t *testing.T) {
	ConfirmRplacd := func(s *Slice, v interface{}, r *Slice) {
		s.Rplacd(v)
		if !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplacd(SList(1, 2, 3, 4, 5), nil, SList(1))
	ConfirmRplacd(SList(1, 2, 3, 4, 5), 10, SList(1, 10))
	ConfirmRplacd(SList(1, 2, 3, 4, 5), SList(5, 4, 3, 2), SList(1, 5, 4, 3, 2))
	ConfirmRplacd(SList(1, 2, 3, 4, 5), SList(2, 4, 8, 16, 32), SList(1, 2, 4, 8, 16, 32))
}