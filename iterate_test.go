package sexp

import(
	"github.com/feyeleanor/slices"
	"testing"
)

func TestEach(t *testing.T) {
	count := 0
	S := slices.Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	Each(S, func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	count = 0
	Each(S, func(i int, v interface{}) {
		if i != count {
			t.Fatalf("index %v erroneously reported as %v", count, i)
		}
		if v != count {
			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	count = 0
	Each(S, func(k, v interface{}) {
		if k != count {
			t.Fatalf("index %v erroneously reported as %v", count, k)
		}
		if v != count {
			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	count = 0
	Each(S, func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	count = 0
	Each(S, func(i int, v interface{}) {
		if i != count {
			t.Fatalf("index %v erroneously reported as %v", count, i)
		}
		if v != count {
			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	count = 0
	Each(S, func(k, v interface{}) {
		if k != count {
			t.Fatalf("index %v erroneously reported as %v", count, k)
		}
		if v != count {
			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	count = 0
	I := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	Each(I, func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	count = 0
	Each(I, func(i int, v interface{}) {
		if i != count {
			t.Fatalf("index %v erroneously reported as %v", count, i)
		}
		if v != count {
			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	count = 0
	Each(I, func(k, v interface{}) {
		if k != count {
			t.Fatalf("index %v erroneously reported as %v", count, k)
		}
		if v != count {
			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	count = 0
	Each(I, func(i int, v int) {
		if i != count {
			t.Fatalf("index %v erroneously reported as %v", count, i)
		}
		if v != count {
			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	count = 0
	Each(I, func(k, v int) {
		if k != count {
			t.Fatalf("index %v erroneously reported as %v", count, k)
		}
		if v != count {
			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	M := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	Each(M, func(i int, v interface{}) {
		if i != v {
			t.Fatalf("index %v erroneously reported as %v", i, v)
		}
	})

	Each(M, func(k, v interface{}) {
		if k != v {
			t.Fatalf("index %v erroneously reported as %v", k, v)
		}
	})

	count = 0
	Each(M, func(k, v int) {
		if k != v {
			t.Fatalf("index %v erroneously reported as %v", k, v)
		}
	})

	count = 0
	Each(M, func(k interface{}, v int) {
		if k != v {
			t.Fatalf("index %v erroneously reported as %v", k, v)
		}
	})
}

func TestCycle(t *testing.T) {
	ConfirmCycle := func(s slices.Slice, c int) {
		iterations := 0
		Cycle(s, c, func(i interface{}) {
			iterations++
		})
		if expected := c * s.Len(); iterations != expected {
			t.Fatalf("cycle(%v): iteration count should be %v but is %v", c, expected, iterations)
		}
	}

	S := slices.Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ConfirmCycle(S, 1)
	ConfirmCycle(S, 2)
	ConfirmCycle(S, 3)
}