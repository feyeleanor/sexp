package sexp

import "testing"

//	Write benchmarks for Equal()
//	Write benchmarks for Len()
//	Write benchmarks for Cap()
//	Write benchmarks for Each()
//	Write benchmarks for Transform()

func BenchmarkReverseReversible(b *testing.B) {
	s := SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	for i := 0; i < b.N; i++ {
		Reverse(s)
	}
}

func BenchmarkReverseIndexable(b *testing.B) {
	s := indexableSlice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		Reverse(s)
	}
}

func BenchmarkReverseReflected(b *testing.B) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		Reverse(s)
	}
}

func BenchmarkDepthNested(b *testing.B) {
	l := SList(0, SList(1, SList(2, SList(3, SList(4, SList(5, SList()))))))
	for i := 0; i < b.N; i++ {
		_ = Depth(l)
	}
}

func BenchmarkDepthReflected(b *testing.B) {
	s := []interface{}{0, []interface{}{1, []interface{}{2, []interface{}{3, []interface{}{4, []interface{}{5, []interface{}{}}}}}}}
	for i := 0; i < b.N; i++ {
		_ = Depth(s)
	}	
}

//	Write benchmarks for Flatten()
//	Write benchmarks for Append()
//	Write benchmarks for AppendContainer()
//	Write benchmarks for Prepend()
//	Write benchmarks for PrependContainer()

func BenchmarkBlockCopyBlitter(b *testing.B) {
	s := SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	for i := 0; i < b.N; i++ {
		BlockCopy(s, 0, 5, 5)
	}
}

func BenchmarkBlockCopyIndexable(b *testing.B) {
	s := indexableSlice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		BlockCopy(s, 0, 5, 5)
	}
}

func BenchmarkBlockCopyReflected(b *testing.B) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		BlockCopy(s, 0, 5, 5)
	}
}

func BenchmarkBlockClearBlitter(b *testing.B) {
	s := SList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	for i := 0; i < b.N; i++ {
		BlockClear(s, 0, 5)
	}
}

func BenchmarkBlockClearIndexable(b *testing.B) {
	l := List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	for i := 0; i < b.N; i++ {
		BlockClear(l, 0, 5)
	}
}

func BenchmarkBlockClearReflected(b *testing.B) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		BlockClear(s, 0, 5)
	}
}

//	Write benchmarks for Reallocate()
//	Write benchmarks for Expand()
//	Write benchmarks for Feed()
//	Write benchmarks for Pipe()