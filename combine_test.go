package sexp

import "testing"

func TestCombine(t *testing.T) {
	ConfirmCombine := func(l, r interface{}, f func(interface{}, interface{}) interface{}, result interface{}) {
		if x := Combine(l, r, f); !Equal(x, result) {
			t.Fatalf("Combine(%v, %v, f) should be %v but is %v", l, r, result, x)
		}
	}
	Add := func(l, r interface{}) interface{} {
		return l.(int) + r.(int)
	}
	ConfirmCombine([]int{0, 1}, []int{3, 4, 5}, Add, []int{3, 5})
	ConfirmCombine([]int{0, 1, 2}, []int{3, 4, 5}, Add, []int{3, 5, 7})
	ConfirmCombine([]int{0, 1, 2, 3}, []int{3, 4}, Add, []int{3, 5})

	ConfirmCombine(map[int]int{0: 0, 1: 1}, map[int]int{0: 3, 1: 4, 2: 5}, Add, map[int]int{0: 3, 1: 5})
	ConfirmCombine(map[int]int{0: 0, 1: 1, 2: 2}, map[int]int{0: 3, 1: 4, 2: 5}, Add, map[int]int{0: 3, 1: 5, 2: 7})
	ConfirmCombine(map[int]int{0: 0, 1: 1, 2: 2, 3: 3}, map[int]int{0: 3, 1: 4}, Add, map[int]int{0: 3, 1: 5})

//	ConfirmCombine([]int{0, 1, 2}, map[int]int{0: 3, 1: 4, 2: 5}, Add, []int{3, 5, 7})
//	ConfirmCombine(map[int]int{0: 0, 1: 1, 2: 2}, []int{3, 4, 5}, Add, map[int]int{0: 3, 1: 5, 2: 7})
}