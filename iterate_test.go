package sexp

import(
	"github.com/feyeleanor/slices"
	"testing"
)

func TestEachSlice(t *testing.T) {
	var count	int

	ConfirmEach := func(s slices.Slice, f interface{}) {
		count = 0
		Each(s, f)
		if count != len(s) {
			t.Fatalf("total iterations should be %v but is %v", len(s), count)
		}
	}

	S := slices.Slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ConfirmEach(S, func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	ConfirmEach(S, func(i int, v interface{}) {
		switch {
		case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	ConfirmEach(S, func(k, v interface{}) {
		switch {
		case k != count:			t.Fatalf("index %v erroneously reported as %v", count, k)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	ConfirmEach(S, func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	ConfirmEach(S, func(i int, v interface{}) {
		switch {
		case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	ConfirmEach(S, func(k, v interface{}) {
		switch {
		case k != count:			t.Fatalf("index %v erroneously reported as %v", count, k)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})
}

func TestEachIntSlice(t *testing.T) {
	var count	int

	ConfirmEach := func(s []int, f interface{}) {
		count = 0
		Each(s, f)
		if count != len(s) {
			t.Fatalf("total iterations should be %v but is %v", len(s), count)
		}
	}

	I := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ConfirmEach(I, func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	ConfirmEach(I, func(i int, v interface{}) {
		switch {
		case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	ConfirmEach(I, func(k, v interface{}) {
		switch {
		case k != count:			t.Fatalf("index %v erroneously reported as %v", count, k)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	ConfirmEach(I, func(i int, v int) {
		switch {
		case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	ConfirmEach(I, func(k, v int) {
		switch {
		case k != count:			t.Fatalf("index %v erroneously reported as %v", count, k)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})
}

func TestEachMap(t *testing.T) {
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

	Each(M, func(k, v int) {
		if k != v {
			t.Fatalf("index %v erroneously reported as %v", k, v)
		}
	})

	Each(M, func(k interface{}, v int) {
		if k != v {
			t.Fatalf("index %v erroneously reported as %v", k, v)
		}
	})
}

func TestEachChannel(t *testing.T) {
	var index	int

	ConfirmEach := func(s []int, f interface{}) {
		var v		int
		var count	int

		C := make(chan int)
		go func() {
			for count, v = range s {
				C <- v
			}
			close(C)
		}()

		index = 0
		Each(C, f)
		if count != len(s) - 1 {
			t.Fatalf("total iterations should be %v but is %v", len(s) - 1, count)
		}
	}

	S := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
	ConfirmEach(S, func(v interface{}) {
		if v != S[index] {
			t.Fatalf("index %v erroneously reported as %v", index, v)
		}
		index++
	})

	ConfirmEach(S, func(i int, v interface{}) {
		switch {
		case i != index:		t.Fatalf("index %v erroneously reported as %v", index, i)
		case v != S[i]:			t.Fatalf("value %v erroneously reported as %v", S[9], v)
		}
		index++
	})
}

func TestEachFunction(t *testing.T) {
	F := func(v interface{}) (r interface{}) {
		if v.(int) < 10 {
			r = v
		}
		return
	}

	count := 0
	Each(F, func(v interface{}) {
		if v != count {
			t.Fatalf("index %v erroneously reported as %v", count, v)
		}
		count++
	})

	Each(F, func(i interface{}, v interface{}) {
		if i != v {
			t.Fatalf("index %v erroneously reported as %v", i, v)
		}
	})

	Each(F, func(i int, v interface{}) {
		if i != v {
			t.Fatalf("index %v erroneously reported as %v", i, v)
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