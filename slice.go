package sexp

import "fmt"

func SList(n... interface{}) Slice {
	return Slice{ &n }
}

type Slice struct {
	nodes	*[]interface{}
}

func (s Slice) IsNil() bool {
	return s.nodes == nil || s.Len() > 0
}

func (s Slice) NotNil() bool {
	return s.nodes != nil && s.Len() > 0
}

func (s Slice) At(i int) interface{} {
	return (*s.nodes)[i]
}

func (s Slice) Set(i int, v interface{}) {
	(*s.nodes)[i] = v
}

func (s Slice) Each(f func(interface{})) {
	for _, v := range *(s.nodes) {
		f(v)
	}
}

func (s Slice) String() (t string) {
	if s.NotNil() {
		for _, v := range *s.nodes {
			if len(t) > 0 {
				t += " "
			}
			t += fmt.Sprintf("%v", v)
		}
	}
	return fmt.Sprintf("(%v)", t)
}

func (s Slice) Len() int {
	return len(*s.nodes)
}

func (s Slice) Depth() (c int) {
	if s.NotNil() {
		for _, v := range *s.nodes {
			if v, ok := v.(Nested); ok {
				if r := v.Depth() + 1; r > c {
					c = r
				}
			}
		}
	}
	return
}

func (s Slice) Reverse() {
	if s.NotNil() {
		end := len(*s.nodes) - 1
		for i := 0; i < end; i++ {
			(*s.nodes)[i], (*s.nodes)[end] = (*s.nodes)[end], (*s.nodes)[i]
			end--
		}
	}
}

func (s Slice) Flatten() {
	if s.NotNil() {
		n := make([]interface{}, 0, cap((*s.nodes)))
		for _, v := range *s.nodes {
			switch v := v.(type) {
			case Slice:			v.Flatten()
								n = append(n, (*v.nodes)...)
			case Flattenable:	v.Flatten()
								n = append(n, v)
			default:			n = append(n, v)
			}
		}
		(*s.nodes) = n
	}
}

func (s Slice) equal(o Slice) (r bool) {
	switch {
	case s.IsNil():
		r = o.IsNil()
	case s.Len() == o.Len():
		r = true
		for i := 0; r && i < s.Len(); i++ {
			n := (*s.nodes)[i]
			x := (*o.nodes)[i]
			fmt.Printf("Slice::equal() n = '%v', x = '%v'\n", n, x)
			if n, ok := n.(Equatable); ok {
				r = n.Equal(x)
			} else {
				r = n == x
			}
		}
	}
	return
}

func (s Slice) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *Slice:		r = s.equal(*o)
	case Slice:			r = s.equal(o)
	}
	return
}

func (s Slice) Car() (h interface{}) {
	if s.NotNil() {
		h = (*s.nodes)[0]
	}
	return
}

func (s Slice) Caar() (h interface{}) {
	if car, ok := s.Car().(Slice); ok {
		h = car.Car()
	}
	return
}

func (s Slice) Cdr() (t Slice) {
	if s.NotNil() {
		switch len(*s.nodes) {
		case 1:		break
		case 2:		if v, ok := (*s.nodes)[1].(Slice); ok {
						t = v
					} else {
						x := (*s.nodes)[1:]
						t.nodes = &x
					}
		default:	x := (*s.nodes)[1:]
					t.nodes = &x
		}
	}
	return
}

func (s Slice) Cddr() (t Slice) {
	return s.Cdr().Cdr()
}

func (s *Slice) Rplaca(v interface{}) {
	if s.IsNil() {
		*s = SList(v)
	} else {
		(*s.nodes)[0] = v
	}
}

func (s *Slice) Rplacd(v interface{}) {
	if s.IsNil() {
		*s = SList(v)
	} else {
		switch v := v.(type) {
		case Slice:			if v.Len() >= cap(*s.nodes) {
								n := make([]interface{}, v.Len() + 1, v.Len() + 1)
								n[0] = (*s.nodes)[0]
								copy(n[1:], *v.nodes)
								*s.nodes = n
							} else {
								copy((*s.nodes)[1:], *v.nodes)
							}
		case nil:			(*s.nodes) = (*s.nodes)[:1]
		default:			(*s.nodes)[1] = v
							(*s.nodes) = (*s.nodes)[:2]
		}
	}
}