package sexp

type memo map[uintptr] interface{}

func (m memo) Memorise(s Addressable) (b bool) {
	a := s.Addr()
	if _, present := m[a]; !present {
		m[a] = s
		b = true
	}
	return
}

func (m memo) Replace(s Addressable) (b bool) {
	a := s.Addr()
	if _, present := m[a]; present {
		m[a] = s
		b = true
	}
	return
}

func (m memo) Forget(s Addressable) {
	m[s.Addr()] = nil	
}

func (m memo) Find(s Addressable) (data interface{}) {
	data, _ = m[s.Addr()]
	return
}