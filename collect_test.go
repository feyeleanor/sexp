package sexp

import(
	"github.com/feyeleanor/slices"
	"testing"
)

func TestCollectSlice(t *testing.T) {
	ConfirmCollect := func(s, r, f interface{}) {
		if x, _ := Collect(s, f); !Equal(x, r) {
			t.Fatalf("collected slice should be %v but is %v", r, x)
		}
	}

	S := slices.Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	R := slices.Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ConfirmCollect(S, R, func(i interface{}) interface{} {
		return i
	})

	R = slices.Slice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ConfirmCollect(S, R, func(i int, v interface{}) interface{} {
		return v.(int) + 1
	})

	R = slices.Slice{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
	ConfirmCollect(S, R, func(k, v interface{}) interface{} {
		return v.(int) * 2
	})

	R = slices.Slice{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ConfirmCollect(S, R, func(k, v interface{}) interface{} {
		return 0
	})

	ConfirmCollect([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81}, func(x int) int {
		return x * x
	})

	ConfirmCollect([]float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []float32{0, 1, 4, 9, 16, 25, 36, 49, 64, 81}, func(x float32) float32 {
		return x * x
	})

	ConfirmCollect(slices.Slice{0, 1.0, 2, 3, 4, 5, 6, 7, 8, 9}, slices.Slice{0, 1.0, 4, 9, 16, 25, 36, 49, 64, 81}, func(x interface{}) (r interface{}) {
		switch x := x.(type) {
		case int:			r = x * x
		case float32:		r = x * x
		case float64:		r = x * x
		default:			t.Fatalf("Typecast failed")
		}
		return
	})
}

func TestCollectIntSlice(t *testing.T) {
	var count	int

	ConfirmCollect := func(s, r []int, f interface{}) {
		count = 0
		if x, _ := Collect(s, f); !Equal(x, r) {
			t.Fatalf("collected slice should be %v but is %v", r, x)
		}
	}

	I := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	R := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ConfirmCollect(I, R, func(i interface{}) interface{} {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
		return count
	})

	ConfirmCollect(I, R, func(i int, v interface{}) interface{} {
		switch {
		case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
		return count
	})

	ConfirmCollect(I, R, func(k, v interface{}) interface{} {
		switch {
		case k != count:			t.Fatalf("index %v erroneously reported as %v", count, k)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
		return count
	})

	ConfirmCollect(I, R, func(i, v int) int {
		switch {
		case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
		return count
	})
}

func TestCollectMap(t *testing.T) {
	ConfirmCollect := func(m, r, f interface{}) {
		if x, _ := Collect(m, f); !Equal(x, r) {
			t.Fatalf("collected map should be %v [%T] but is %v [%T]", r, r, x, x)
		}
	}

	M1 := map[int] int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	R1 := map[int] int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5, 5: 6, 6: 7, 7: 8, 8: 9, 9: 10}
	ConfirmCollect(M1, R1, func(k, v int) int {
		if k != v {
			t.Fatalf("index %v erroneously reported as %v", k, v)
		}
		return k + 1
	})

	ConfirmCollect(M1, R1, func(i int, v interface{}) int {
		if i != v {
			t.Fatalf("index %v erroneously reported as %v", i, v)
		}
		return i + 1
	})

	ConfirmCollect(M1, R1, func(k, v interface{}) int {
		if k != v {
			t.Fatalf("index %v erroneously reported as %v", k, v)
		}
		return k.(int) + 1
	})

	ConfirmCollect(M1, R1, func(k interface{}, v int) int {
		if k != v {
			t.Fatalf("index %v erroneously reported as %v", k, v)
		}
		return k.(int) + 1
	})

	M2 := map[int] interface{}{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	R2 := map[int] interface{}{0: 1, 1: 2, 2: 3, 3: 4, 4: 5, 5: 6, 6: 7, 7: 8, 8: 9, 9: 10}
	ConfirmCollect(M2, R2, func(i int, v interface{}) interface{} {
		if i != v {
			t.Fatalf("index %v erroneously reported as %v", i, v)
		}
		return i + 1
	})

	ConfirmCollect(M2, R2, func(k, v interface{}) interface{} {
		if k != v {
			t.Fatalf("index %v erroneously reported as %v", k, v)
		}
		return k.(int) + 1
	})
}