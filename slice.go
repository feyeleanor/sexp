package sexp

import "fmt"
import "reflect"
import "unsafe"


func SList(n... interface{}) Slice {
 	return append(make(Slice, 0, len(n)), n...)
}

type Slice []interface{}

func (s Slice) String() (t string) {
	for _, v := range s {
		if len(t) == 0 {
			t = fmt.Sprintf("%v", v)
		} else {
			t = fmt.Sprintf("%v %v", t, v)
		}
	}
	return fmt.Sprintf("(%v)", t)
}

func (s Slice) Addr() uintptr {
	return uintptr(unsafe.Pointer(&s))
}

func (s Slice) Len() int {
	return len(s)
}

func (s Slice) depth(visited_nodes memo) (c int) {
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

func (s Slice) Depth() (c int) {
	return s.depth(make(memo)) - 1
}

func (s Slice) Reverse() {
	end := len(s) - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s Slice) flatten(visited_nodes memo) (n Slice) {
	for _, v := range s {
		switch v := v.(type) {
		case Slice:		if visited_nodes.Memorise(&v) {
							n = append(n, v.flatten(visited_nodes)...)
						} else {
							n = append(n, v)
						}
		default:		n = append(n, v)
		}
	}
	return
}

func (s *Slice) Flatten() {
	*s = s.flatten(make(memo))
}

func (s Slice) Equal(o interface{}) (r bool) {
	return reflect.DeepEqual(s, o)
}

func (s Slice) Car() (h interface{}) {
	if len(s) > 0 {
		h = s[0]
	}
	return
}

func (s Slice) Caar() (h interface{}) {
	car := s.Car()
	if car, ok := car.(Slice); ok {
		h = car.Car()
	}
	return
}

func (s Slice) Cdr() (t Slice) {
	switch len(s) {
	case 0:		fallthrough
	case 1:		break
	case 2:		if v, ok := s[1].(Slice); ok {
					t = v
				} else {
					t = s[1:]
				}
	default:	t = s[1:]
	}
	return
}

func (s Slice) Cddr() (t Slice) {
	if t = s.Cdr(); t != nil {
		t = t.Cdr()
	}
	return
}

func (s *Slice) Rplaca(v interface{}) {
	switch len(*s) {
	case 0:		*s = Slice{ v }
	default:	(*s)[0] = v
	}
}

func (s *Slice) Rplacd(v interface{}) {
	if len(*s) == 0 {
		*s = Slice{ v }
	} else {
		switch v := v.(type) {
		case Slice:			if len(v) >= cap(*s) {
								n := make(Slice, len(v) + 1, len(v) + 1)
								n[0] = (*s)[0]
								copy(n[1:], v)
								*s = n
							} else {
								copy((*s)[1:], v)
							}
		case nil:			*s = (*s)[:1]
		default:			(*s)[1] = v
							*s = (*s)[:2]
		}
	}
}