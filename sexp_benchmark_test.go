package sexp

import "testing"

func BenchmarkSCons2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SCons(0, 1)
	}
}

func BenchmarkSCons10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
}

func BenchmarkSCons2x2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SCons(0, SCons(0, 1))
	}
}

func BenchmarkSCons2x10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SCons(0, SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	}
}

func BenchmarkSCons10x2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SCons(SCons(0, 1), SCons(1, 2), SCons(2, 3), SCons(3, 4), SCons(4, 5), SCons(5, 6), SCons(6, 7), SCons(7, 8), SCons(8, 9), SCons(9, 0))
	}
}

func BenchmarkSCons10x10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SCons(	SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)	)
	}
}

func BenchmarkLen1(b *testing.B) {
	v := SEXP{ 0 }
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkLen1x1(b *testing.B) {
	v := SEXP{ 0, SEXP{ 0 } }
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkLen1x10(b *testing.B) {
	v := SEXP{ SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9) }
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkLen10(b *testing.B) {
	v := SEXP{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkLen10x2(b *testing.B) {
	v := SCons(SCons(0, 1), SCons(1, 2), SCons(2, 3), SCons(3, 4), SCons(4, 5), SCons(5, 6), SCons(6, 7), SCons(7, 8), SCons(8, 9), SCons(9, 0))
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkLen10x10(b *testing.B) {
	v := SCons(	SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)	)
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkDepth1(b *testing.B) {
	v := SEXP{ 0 }
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkDepth1x1(b *testing.B) {
	v := SEXP{ 0, SEXP{ 0 } }
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkDepth1x10(b *testing.B) {
	v := SEXP{ SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9) }
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkDepth10(b *testing.B) {
	v := SEXP{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkDepth10x2(b *testing.B) {
	v := SCons(SCons(0, 1), SCons(1, 2), SCons(2, 3), SCons(3, 4), SCons(4, 5), SCons(5, 6), SCons(6, 7), SCons(7, 8), SCons(8, 9), SCons(9, 0))
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkDepth10x10(b *testing.B) {
	v := SCons(	SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)	)
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkReverse10(b *testing.B) {
	b.StopTimer()
		v := SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v.Reverse()
	}
}

func BenchmarkReverse10x10(b *testing.B) {
	b.StopTimer()
		v := SCons(	SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)	)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v.Reverse()
	}
}
func BenchmarkFlatten1(b *testing.B) {
	v := SEXP{ 0 }
	for i := 0; i < b.N; i++ {
		v.Flatten()
	}
}

func BenchmarkFlatten1x1(b *testing.B) {
	v := SEXP{ 0, SEXP{ 0 } }
	for i := 0; i < b.N; i++ {
		v.Flatten()
	}
}

func BenchmarkFlatten1x10(b *testing.B) {
	v := SEXP{ SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9) }
	for i := 0; i < b.N; i++ {
		v.Flatten()
	}
}

func BenchmarkFlatten10(b *testing.B) {
	v := SEXP{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	for i := 0; i < b.N; i++ {
		v.Flatten()
	}
}

func BenchmarkFlatten10x2(b *testing.B) {
	v := SCons(SCons(0, 1), SCons(1, 2), SCons(2, 3), SCons(3, 4), SCons(4, 5), SCons(5, 6), SCons(6, 7), SCons(7, 8), SCons(8, 9), SCons(9, 0))
	for i := 0; i < b.N; i++ {
		v.Flatten()
	}
}

func BenchmarkFlatten10x10(b *testing.B) {
	v := SCons(	SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				SCons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)	)
	for i := 0; i < b.N; i++ {
		v.Flatten()
	}
}

func BenchmarkCar(b *testing.B) {
	v := SCons(0, 1)
	for i := 0; i < b.N; i++ {
		_ = v.Car()
	}
}

func BenchmarkCaar(b *testing.B) {
	v := SCons(SCons(0, 1), 2)
	for i := 0; i < b.N; i++ {
		_ = v.Caar()
	}
}

func BenchmarkCdr(b *testing.B) {
	v := SCons(0, 1)
	for i := 0; i < b.N; i++ {
		_ = v.Cdr()
	}
}

func BenchmarkCddr(b *testing.B) {
	v := SCons(0, SCons(1, 2))
	for i := 0; i < b.N; i++ {
		_ = v.Cddr()
	}
}

func BenchmarkRplaca(b *testing.B) {}

func BenchmarkRplacd(b *testing.B) {}