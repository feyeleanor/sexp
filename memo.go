package sexp

type memo map[uintptr] interface{}

func (m memo) Memorise(s *SEXP) (b bool) {
	a := s.address()
	if b = (m[a] == nil); !b {
		m[a] = s
	}
	return
}

func (m memo) Forget(s SEXP) {
	m[s.address()] = nil	
}