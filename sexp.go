package sexp

import "reflect"

type SExp []interface{}

func Cons(l interface{}, r... interface{}) (s SExp) {
	length := len(r) + 1
	if length < 2 {
		length = 2
	}
	s = make(SExp, length, length)
	s[0] = l
	if len(r) > 0 {
		copy(s[1:], r)
	}
	return
}

func (s SExp) Reverse() {
	end := len(s) - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
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
	case 0:		*s = Cons(v)
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