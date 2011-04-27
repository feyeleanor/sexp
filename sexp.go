package sexp

import "fmt"
import "reflect"
import "unsafe"


func SCons(n... interface{}) SEXP {
 	return append(make(SEXP, 0, len(n)), n...)
}

type SEXP []interface{}

func (s SEXP) String() (t string) {
	for _, v := range s {
		if len(t) == 0 {
			t = fmt.Sprintf("%v", v)
		} else {
			t = fmt.Sprintf("%v %v", t, v)
		}
	}
	return fmt.Sprintf("(%v)", t)
}

func (s SEXP) Addr() uintptr {
	return uintptr(unsafe.Pointer(&s))
}

func (s SEXP) Len() int {
	return len(s)
}

func (s SEXP) depth(visited_nodes memo) (c int) {
	if visited_nodes.Memorise(s) {
		for _, v := range s {
			if v, ok := v.(CyclicNested); ok {
				if r := v.depth(visited_nodes); r > c {
					c = r
				}
			}
		}
		visited_nodes.Forget(s)
	}
	c++
	return
}

func (s SEXP) Depth() (c int) {
	return s.depth(make(memo)) - 1
}

func (s SEXP) Reverse() {
	end := len(s) - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s SEXP) flatten(visited_nodes memo) (n SEXP) {
	for _, v := range s {
		switch v := v.(type) {
		case SEXP:		if visited_nodes.Memorise(&v) {
							n = append(n, v.flatten(visited_nodes)...)
						} else {
							n = append(n, v)
						}
		default:		n = append(n, v)
		}
	}
	return
}

func (s *SEXP) Flatten() {
	*s = s.flatten(make(memo))
}

func (s SEXP) Equal(o interface{}) (r bool) {
	return reflect.DeepEqual(s, o)
}

func (s SEXP) Car() (h interface{}) {
	if len(s) > 0 {
		h = s[0]
	}
	return
}

func (s SEXP) Caar() (h interface{}) {
	car := s.Car()
	if car, ok := car.(SEXP); ok {
		h = car.Car()
	}
	return
}

func (s SEXP) Cdr() (t SEXP) {
	switch len(s) {
	case 0:		fallthrough
	case 1:		break
	case 2:		if v, ok := s[1].(SEXP); ok {
					t = v
				} else {
					t = s[1:]
				}
	default:	t = s[1:]
	}
	return
}

func (s SEXP) Cddr() (t SEXP) {
	if t = s.Cdr(); t != nil {
		t = t.Cdr()
	}
	return
}

func (s *SEXP) Rplaca(v interface{}) {
	switch len(*s) {
	case 0:		*s = SEXP{ v }
	case 1:		(*s)[0] = v
	default:	*s = SCons(v, (*s)[1:])
	}
}

func (s *SEXP) Rplacd(v interface{}) {
	if len(*s) == 0 {
		*s = SCons(nil, v)
	} else {
		(*s)[1] = v
		*s = (*s)[:2]
	}
}