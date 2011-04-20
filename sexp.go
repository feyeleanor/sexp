package sexp

import "fmt"
import "reflect"
import "unsafe"


type memo map[uintptr] interface{}

func (m memo) Memorise(s *SExp) (b bool) {
	a := s.address()
	if b = (m[a] == nil); !b {
		m[a] = s
	}
	return
}

func (m memo) Forget(s SExp) {
	m[s.address()] = nil	
}


func Cons(a, b interface{}, n... interface{}) (s SExp) {
	length := len(n) + 2
	s = make(SExp, length, length)
	s[0] = a
	s[1] = b
	if len(n) > 0 {
		copy(s[2:], n)
	}
	return
}


type SExp []interface{}

func (s SExp) String() (t string) {
	for _, v := range s {
		if len(t) == 0 {
			t = fmt.Sprintf("%v", v)
		} else {
			t = fmt.Sprintf("%v %v", t, v)
		}
	}
	return fmt.Sprintf("(%v)", t)
}

func (s *SExp) address() uintptr {
	return uintptr(unsafe.Pointer(s))
}

func (s *SExp) len(visited_nodes memo) (c int) {
	c = len(*s)
	if visited_nodes.Memorise(s) {
		for _, v := range *s {
			if v, ok := v.(SExp); ok {
				if visited_nodes.Memorise(&v) {
					c += v.len(visited_nodes) - 1
				}
			}
		}
	}
	return
}

func (s SExp) Len() int {
	return s.len(make(memo))
}

func (s *SExp) depth(visited_nodes memo) (c int) {
	if visited_nodes.Memorise(s) {
		for _, v := range *s {
			if v, ok := v.(SExp); ok {
				if c == 0 {
					c = 1
				}
				if visited_nodes.Memorise(&v) {
					r := v.depth(visited_nodes)
					if r >= c {
						c = r + 1
					}
				}
			}
		}
	}
	return
}

func (s SExp) Depth() (c int) {
	return s.depth(make(memo))
}

func (s *SExp) bounds(visited_nodes memo) (l, d int) {
	l = len(*s)
	if visited_nodes.Memorise(s) {
		for _, v := range *s {
			if v, ok := v.(SExp); ok {
				if d == 0 {
					d = 1
				}
				if visited_nodes.Memorise(&v) {
					nl, nd := v.bounds(visited_nodes)
					if nd >= d {
						d = nd + 1
					}
					l += nl -1
				}
			}
		}
	}
	return
}

//	Bounds calculates both the Length and Depth of the SExp in a single pass
func (s SExp) Bounds() (l, d int) {
	return s.bounds(make(memo))
}

func (s SExp) Reverse() {
	end := len(s) - 1
	for i := 0; i < end; i++ {
		if c, ok := s[i].(SExp); ok {
			c.Reverse()
		}
		if c, ok := s[end].(SExp); ok {
			c.Reverse()
		}
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s *SExp) flatten(visited_nodes memo) (n SExp) {
	l := s.Len()
	n = make(SExp, l, l)
	for i, j := 0, 0; i < len(n); i++ {
		v := (*s)[j]
		switch v := v.(type) {
		case SExp:		if visited_nodes.Memorise(&v) {
							r := v.flatten(visited_nodes)
							copy(n[i:], r)
							i += len(r) - 1
						} else {
							n[i] = v
						}
		default:		n[i] = v
		}
		j++
	}
	return
}

func (s *SExp) Flatten() {
	*s = s.flatten(make(memo))
}

func (s SExp) Equal(o interface{}) (r bool) {
	return reflect.DeepEqual(s, o.(SExp))
}

func (s SExp) Car() (h interface{}) {
	if len(s) > 0 {
		h = s[0]
	}
	return
}

func (s SExp) Caar() (h interface{}) {
	car := s.Car()
	if car, ok := car.(SExp); ok {
		h = car.Car()
	}
	return
}

func (s SExp) Cdr() (t SExp) {
	switch len(s) {
	case 0:		fallthrough
	case 1:		break
	case 2:		if v, ok := s[1].(SExp); ok {
					t = v
				} else {
					t = s[1:]
				}
	default:	t = s[1:]
	}
	return
}

func (s SExp) Cddr() (t SExp) {
	if t = s.Cdr(); t != nil {
		t = t.Cdr()
	}
	return
}

func (s *SExp) Rplaca(v interface{}) {
	switch len(*s) {
	case 0:		*s = SExp{ v }
	case 1:		(*s)[0] = v
	default:	*s = Cons(v, (*s)[1:])
	}
}

func (s *SExp) Rplacd(v interface{}) {
	if len(*s) == 0 {
		*s = Cons(nil, v)
	} else {
		(*s)[1] = v
		*s = (*s)[:2]
	}
}