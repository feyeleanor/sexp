package sexp

import(
	"github.com/feyeleanor/slices"
	"testing"
)

func TestTransformSlice(t *testing.T) {
	var count	int

	ConfirmTransform := func(s, r, f interface{}) {
		count = 0
		Transform(s, f)
		if !Equal(s, r) {
			t.Fatalf("transformed slice should be %v but is %v", r, s)
		}
	}

	S := slices.Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	R := slices.Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ConfirmTransform(S, R, func(i interface{}) interface{} {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
		return i
	})

	R = slices.Slice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ConfirmTransform(S, R, func(i int, v interface{}) interface{} {
		switch {
		case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", i, v)
		}
		count++
		return v.(int) + 1
	})

	R = slices.Slice{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	ConfirmTransform(S, R, func(k, v interface{}) interface{} {
		switch {
		case k != count:			t.Fatalf("index %v erroneously reported as %v", count, k)
		case v != count + 1:		t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
		return v.(int) * 2
	})

	R = slices.Slice{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
	ConfirmTransform(S, R, func(i interface{}) interface{} {
		count++
		return i.(int) - 2
	})

	R = slices.Slice{0, 4, 12, 24, 40, 60, 84, 112, 144, 180}
	ConfirmTransform(S, R, func(i int, v interface{}) interface{} {
		count++
		return v.(int) * count
	})

	R = slices.Slice{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ConfirmTransform(S, R, func(k, v interface{}) interface{} {
		count++
		return 0
	})

	ConfirmTransform([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81}, func(x int) int {
		count++
		return x * x
	})

	ConfirmTransform([]float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []float32{0, 1, 4, 9, 16, 25, 36, 49, 64, 81}, func(x float32) float32 {
		count++
		return x * x
	})

	ConfirmTransform(slices.Slice{0, 1.0, 2, 3, 4, 5, 6, 7, 8, 9}, slices.Slice{0, 1.0, 4, 9, 16, 25, 36, 49, 64, 81}, func(x interface{}) (r interface{}) {
		switch x := x.(type) {
		case int:			r = x * x
		case float32:		r = x * x
		case float64:		r = x * x
		default:			t.Fatalf("Typecast failed")
		}
		count++
		return
	})
}

func TestTransformIntSlice(t *testing.T) {
	ConfirmTransform := func(s, r []int, f interface{}) {
		if Transform(s, f); !Equal(s, r) {
			t.Fatalf("transformed slice should be %v but is %v", r, s)
		}
	}

	I := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	R := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ConfirmTransform(I, R, func(i interface{}) interface{} {
		return i.(int) + 1
	})

	I = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ConfirmTransform(I, R, func(i int, v interface{}) interface{} {
		return v.(int) + 1
	})

	I = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ConfirmTransform(I, R, func(k, v interface{}) interface{} {
		return v.(int) + 1
	})

	I = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ConfirmTransform(I, R, func(i, v int) int {
		return v + 1
	})
}

func TestTransformMap(t *testing.T) {
	ConfirmTransform := func(m, r, f interface{}) {
		if Transform(m, f); !Equal(m, r) {
			t.Fatalf("transformed map should be %v but is %v", r, m)
		}
	}

	M1 := map[int] int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	R1 := map[int] int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5, 5: 6, 6: 7, 7: 8, 8: 9, 9: 10}
	ConfirmTransform(M1, R1, func(k, v int) int {
		return v + 1
	})

	M1 = map[int] int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	ConfirmTransform(M1, R1, func(i int, v interface{}) int {
		return v.(int) + 1
	})

	M1 = map[int] int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	ConfirmTransform(M1, R1, func(k, v interface{}) int {
		return v.(int) + 1
	})

	M1 = map[int] int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	ConfirmTransform(M1, R1, func(k interface{}, v int) int {
		return v + 1
	})

	M2 := map[int] interface{}{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	R2 := map[int] interface{}{0: 1, 1: 2, 2: 3, 3: 4, 4: 5, 5: 6, 6: 7, 7: 8, 8: 9, 9: 10}
	ConfirmTransform(M2, R2, func(i int, v interface{}) interface{} {
		return v.(int) + 1
	})

	M2 = map[int] interface{}{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	ConfirmTransform(M2, R2, func(k, v interface{}) interface{} {
		return v.(int) + 1
	})
}