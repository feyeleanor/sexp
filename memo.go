package sexp

type Addressable interface {
	Addr() uintptr
}

type memo map[uintptr] interface{}

func (m memo) Memorise(s Addressable) (b bool) {
	a := s.Addr()
	if b = (m[a] == nil); !b {
		m[a] = s
	}
	return
}

func (m memo) Forget(s Addressable) {
	m[s.Addr()] = nil	
}