package sexp

import(
//	"github.com/feyeleanor/slices"
//	"strconv"
	"testing"
)

type partially_iterable_slice []interface{}

func (s partially_iterable_slice) While(r bool, f interface{}) (ok bool, count int) {
	switch f := f.(type) {
	case func(interface{}) bool:					for _, v := range s {
														if f(v) != r {
															break 
														}
														count++
													}
													ok = true

	case func(int, interface{}) bool:				for i, v := range s {
														if f(i, v) != r {
															break
														}
														count++
													}
													ok = true

	case func(interface{}, interface{}) bool:		for i, v := range s {
														if f(i, v) != r {
															break
														}
														count++
													}
													ok = true
	}
	return
}

func TestWhileSlice(t *testing.T) {
	ConfirmWhileTrue := func(s interface{}, r int) {
		switch ok, count := While(s, func(v interface{}) bool { return v == true }); {
		case !ok:			t.Fatalf("failed to perform partial iteration over %v", s)
		case count != r:	t.Fatalf("total iterations should be %v but are %v", r, count)
		}

		switch ok, count := While(s, func(i int, v interface{}) bool { return v == true }); {
		case !ok:			t.Fatalf("failed to perform int indexed partial iteration over %v", s)
		case count != r:	t.Fatalf("total iterations should be %v but are %v", r, count)
		}

		switch ok, count := While(s, func(i, v interface{}) bool { return v == true }); {
		case !ok:			t.Fatalf("failed to perform intterface{} indexed partial iteration over %v", s)
		case count != r:	t.Fatalf("total iterations should be %v but are %v", r, count)
		}
	}

	ConfirmWhileTrue(partially_iterable_slice{false}, 0)
	ConfirmWhileTrue(partially_iterable_slice{false, true}, 0)
	ConfirmWhileTrue(partially_iterable_slice{true, false}, 1)
	ConfirmWhileTrue(partially_iterable_slice{true, true}, 2)

	ConfirmWhileTrue(indexable_slice{false}, 0)
	ConfirmWhileTrue(indexable_slice{false, true}, 0)
	ConfirmWhileTrue(indexable_slice{true, false}, 1)
	ConfirmWhileTrue(indexable_slice{true, true}, 2)
}

func TestUntilSlice(t *testing.T) {
	ConfirmWhileFalse := func(s interface{}, r int) {
		switch ok, count := Until(s, func(v interface{}) bool { return v == true }); {
		case !ok:			t.Fatalf("failed to perform partial iteration over %v", s)
		case count != r:	t.Fatalf("total iterations should be %v but are %v", r, count)
		}

		switch ok, count := Until(s, func(i int, v interface{}) bool { return v == true }); {
		case !ok:			t.Fatalf("failed to perform int indexed partial iteration over %v", s)
		case count != r:	t.Fatalf("total iterations should be %v but are %v", r, count)
		}

		switch ok, count := Until(s, func(i, v interface{}) bool { return v == true }); {
		case !ok:			t.Fatalf("failed to perform intterface{} indexed partial iteration over %v", s)
		case count != r:	t.Fatalf("total iterations should be %v but are %v", r, count)
		}
	}

	ConfirmWhileFalse(indexable_slice{false}, 1)
	ConfirmWhileFalse(indexable_slice{false, true}, 1)
	ConfirmWhileFalse(indexable_slice{true, false}, 0)
	ConfirmWhileFalse(indexable_slice{true, true}, 0)
}