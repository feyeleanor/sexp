package sexp

import "fmt"

func SList(n... interface{}) *Slice {
	s := Slice(n)
	return &s
}

type Slice	[]interface{}

func (s Slice) At(i int) interface{} {
	return s[i]
}

func (s Slice) Set(i int, v interface{}) {
	s[i] = v
}

func (s Slice) Clear(i int) {
	s[i] = nil
}

func (s Slice) Each(f func(interface{})) {
	for _, v := range s {
		f(v)
	}
}

func (s Slice) String() (t string) {
	for _, v := range s {
		if len(t) > 0 {
			t += " "
		}
		t += fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("(%v)", t)
}

func (s Slice) Len() int {
	return len(s)
}

func (s Slice) Cap() int {
	return cap(s)
}

func (s Slice) BlockCopy(destination, source, count int) {
	end := source + count
	if end > len(s) {
		end = len(s)
	}
	copy(s[destination:], s[source:end])
}

func (s Slice) BlockClear(start, count int) {
	n := make(Slice, count, count)
	copy(s[start:], n)
}

func (s Slice) Overwrite(offset int, source Slice) {
	copy(s[offset:], source)
}

func (s *Slice) Reallocate(length, capacity int) {
	switch {
	case length > capacity:		s.Reallocate(capacity, capacity)
	case capacity != cap(*s):	x := make(Slice, length, capacity)
								copy(x, *s)
								*s = x
	default:					*s = (*s)[:length]
	}
}

func (s Slice) Depth() (c int) {
	for _, v := range s {
		if v, ok := v.(Nested); ok {
			if r := v.Depth() + 1; r > c {
				c = r
			}
		}
	}
	return
}

func (s Slice) Reverse() {
	end := s.Len() - 1
	for i := 0; i < end; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func (s *Slice) Append(v interface{}) {
	*s = append(*s, v)
}

func (s *Slice) AppendSlice(o Slice) {
	*s = append(*s, o...)
}

func (s *Slice) Prepend(v interface{}) {
	l := s.Len() + 1
	n := make(Slice, l, l)
	n[0] = v
	copy(n[1:], *s)
	*s = n
}

func (s *Slice) PrependSlice(o Slice) {
	l := s.Len() + o.Len()
	n := make(Slice, l, l)
	copy(n, o)
	copy(n[o.Len():], *s)
	*s = n
}

func (s Slice) Repeat(count int) Slice {
	length := len(s) * count
	capacity := cap(s)
	if capacity < length {
		capacity = length
	}
	destination := make(Slice, length, capacity)
	for start, end := 0, len(s); count > 0; count-- {
		copy(destination[start:end], s)
		start = end
		end += len(s)
	}
	return destination
}

func (s *Slice) Flatten() {
	if s != nil {
		n := make(Slice, 0, 0)
		for _, v := range *s {
			switch v := v.(type) {
			case *Slice:			v.Flatten()
									n = append(n, (*v)...)
			case Slice:				(&v).Flatten()
									n = append(n, v...)
			case *[]interface{}:	n = append(n, (*v)...)
			case []interface{}:		n = append(n, v...)
			case Flattenable:		v.Flatten()
									n = append(n, v)
			default:				n = append(n, v)
			}
		}
		*s = n
	}
}

func (s Slice) equal(o Slice) (r bool) {
	switch {
	case s == nil:				r = o == nil
	case s.Len() == o.Len():	r = true
								for i, v := range s {
									switch v := v.(type) {
									case Equatable:		r = v.Equal(o[i])
									default:			r = v == o[i]
									}
									if !r {
										return
									}
								}
	}
	return
}

func (s Slice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *Slice:			r = o != nil && s.equal(*o)
	case Slice:				r = s.equal(o)
	case *[]interface{}:	r = o != nil && s.equal(([]interface{})(*o))
	case []interface{}:		r = s.equal(([]interface{})(o))
	}
	return
}

func (s Slice) Car() (h interface{}) {
	if s.Len() > 0 {
		h = s[0]
	}
	return
}

func (s Slice) Caar() (h interface{}) {
	switch car := s.Car().(type) {
	case *Slice:		h = car.Car()
	case Slice:			h = car.Car()
	}
	return
}

func (s Slice) Cdr() (t Slice) {
	switch s.Len() {
	case 0:		fallthrough
	case 1:		break
	case 2:		switch v := s[1].(type) {
				case *Slice:		t = *v
				case Slice:			t = v
				default:			t = s[1:]
				}
	default:	t = s[1:]
	}
	return
}

func (s Slice) Cddr() Slice {
	return s.Cdr().Cdr()
}

func (s *Slice) Rplaca(v interface{}) {
	switch {
	case s == nil:			*s = *SList(v)
	case s.Len() == 0:		*s = append(*s, v)
	default:				(*s)[0] = v
	}
}

func (s *Slice) Rplacd(v interface{}) {
	if s == nil {
		*s = *SList(v)
	} else {
		ReplaceSlice := func(v Slice) {
			if l := v.Len(); l >= cap(*s) {
				l++
				n := make([]interface{}, l, l)
				copy(n, (*s)[:1])
				copy(n[1:], v)
				*s = n
			} else {
				copy((*s)[1:], v)
			}
		}

		switch v := v.(type) {
		case *Slice:		ReplaceSlice(*v)
		case Slice:			ReplaceSlice(v)
		case nil:			*s = (*s)[:1]
		default:			(*s)[1] = v
							*s = (*s)[:2]
		}
	}
}