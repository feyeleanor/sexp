package sexp

import "testing"

func BenchmarkCons2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Cons(0, 1)
	}
}

func BenchmarkCons10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
}

func BenchmarkCons2x2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Cons(0, Cons(0, 1))
	}
}

func BenchmarkCons2x10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Cons(0, Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	}
}

func BenchmarkCons10x2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Cons(Cons(0, 1), Cons(1, 2), Cons(2, 3), Cons(3, 4), Cons(4, 5), Cons(5, 6), Cons(6, 7), Cons(7, 8), Cons(8, 9), Cons(9, 0))
	}
}

func BenchmarkCons10x10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Cons(	Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)	)
	}
}

func BenchmarkLen1(b *testing.B) {
	v := SExp{ 0 }
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkLen1x1(b *testing.B) {
	v := SExp{ 0, SExp{ 0 } }
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkLen1x10(b *testing.B) {
	v := SExp{ Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9) }
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkLen10(b *testing.B) {
	v := SExp{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkLen10x2(b *testing.B) {
	v := Cons(Cons(0, 1), Cons(1, 2), Cons(2, 3), Cons(3, 4), Cons(4, 5), Cons(5, 6), Cons(6, 7), Cons(7, 8), Cons(8, 9), Cons(9, 0))
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkLen10x10(b *testing.B) {
	v := Cons(	Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)	)
	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkDepth1(b *testing.B) {
	v := SExp{ 0 }
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkDepth1x1(b *testing.B) {
	v := SExp{ 0, SExp{ 0 } }
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkDepth1x10(b *testing.B) {
	v := SExp{ Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9) }
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkDepth10(b *testing.B) {
	v := SExp{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkDepth10x2(b *testing.B) {
	v := Cons(Cons(0, 1), Cons(1, 2), Cons(2, 3), Cons(3, 4), Cons(4, 5), Cons(5, 6), Cons(6, 7), Cons(7, 8), Cons(8, 9), Cons(9, 0))
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkDepth10x10(b *testing.B) {
	v := Cons(	Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)	)
	for i := 0; i < b.N; i++ {
		_ = v.Depth()
	}
}

func BenchmarkBounds1(b *testing.B) {
	v := SExp{ 0 }
	for i := 0; i < b.N; i++ {
		_, _ = v.Bounds()
	}
}

func BenchmarkBounds1x1(b *testing.B) {
	v := SExp{ 0, SExp{ 0 } }
	for i := 0; i < b.N; i++ {
		_, _ = v.Bounds()
	}
}

func BenchmarkBounds1x10(b *testing.B) {
	v := SExp{ Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9) }
	for i := 0; i < b.N; i++ {
		_, _ = v.Bounds()
	}
}

func BenchmarkBounds10(b *testing.B) {
	v := SExp{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	for i := 0; i < b.N; i++ {
		_, _ = v.Bounds()
	}
}

func BenchmarkBounds10x2(b *testing.B) {
	v := Cons(Cons(0, 1), Cons(1, 2), Cons(2, 3), Cons(3, 4), Cons(4, 5), Cons(5, 6), Cons(6, 7), Cons(7, 8), Cons(8, 9), Cons(9, 0))
	for i := 0; i < b.N; i++ {
		_, _ = v.Bounds()
	}
}

func BenchmarkBounds10x10(b *testing.B) {
	v := Cons(	Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
				Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)	)
	for i := 0; i < b.N; i++ {
		_, _ = v.Bounds()
	}
}

func BenchmarkReverse10(b *testing.B) {
	b.StopTimer()
		v := Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v.Reverse()
	}
}

func BenchmarkReverse10x10(b *testing.B) {
	b.StopTimer()
		v := Cons(	Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
					Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)	)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v.Reverse()
	}
}

func BenchmarkCar(b *testing.B) {}

func BenchmarkCaar(b *testing.B) {}

func BenchmarkCdr(b *testing.B) {}

func BenchmarkCddr(b *testing.B) {}

func BenchmarkRplaca(b *testing.B) {}

func BenchmarkRplacd(b *testing.B) {}