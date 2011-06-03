package sexp

import "fmt"

func Cons(items... interface{}) (c *ConsCell) {
	var n *ConsCell
	for i, v := range items {
		if i == 0 {
			c = &ConsCell{ Head: v }
			n = c
		} else {
			n.Tail = &ConsCell{ Head: v }
			n = n.Tail
		}
	}
	return
}

type ConsCell struct {
	Head		interface{}
	Tail		*ConsCell
}

func (c ConsCell) Content() interface{} {
	return c.Head
}

func (c *ConsCell) End() (r *ConsCell) {
	if c != nil {
		for r = c; r.Tail != nil; r = r.Tail {}
	}
	return
}

func (c *ConsCell) MoveTo(i int) (l ListNode) {
	switch {
	case i < 0:				break
	case i == 0:			l = c
	default:				n := c
							for ; i > 0 && n != nil; i-- {
								n = n.Tail
							}
							if n != nil {
								l = n
							}
	}
	return
}

func (c *ConsCell) Link(i int, l ListNode) (b bool) {
	if l == nil {
		switch i {
		case PREVIOUS_NODE:		b = false

		case CURRENT_NODE:		c = nil
								b = true

		case NEXT_NODE:			c.Tail = nil
								b = true
		}
	} else {
		switch i {
		case PREVIOUS_NODE:		if n, ok := l.(*ConsCell); ok {
									n.Tail = c
								} else {
									l.Link(NEXT_NODE, c)
								}
								b = true

		case CURRENT_NODE:		if n, ok := l.(*ConsCell); ok {
									*c = *n
									b = true
								} else {
									if t, ok := NextNode(l).(*ConsCell); ok {
										c.Head = l.Content()
										c.Tail = t
										b = true
									}
								}

		case NEXT_NODE:			if n, ok := l.(*ConsCell); ok {
									c.Tail = n
									b = true
								}
		}
	}
	return
}

func (c *ConsCell) Store(i int, v interface{}) bool {
	switch {
	case i == CURRENT_NODE:		if c == nil {
									c = &ConsCell{}
								}
	case i > CURRENT_NODE:		if c == nil {
									c = &ConsCell{}
									i--
								}
								for ; i > 0; i-- {
									c.Tail = &ConsCell{ Head: nil }
									c = c.Tail
								}

	case i < CURRENT_NODE:		for ; i < 0; i++ {
									*c = ConsCell{ Head: nil, Tail: c }
								}
	}
	c.Head = v
	return true
}

func (c *ConsCell) Append(x interface{}) {
	c.Store(NEXT_NODE, x)
}

func (c *ConsCell) Prepend(x interface{}) {
	*c = ConsCell{ Head: x, Tail: c }
}

func (c ConsCell) equal(o ConsCell) (r bool) {
	defer func() {
		if x := recover(); x != nil {
			r = false
		}
	}()
	if v, ok := c.Head.(Equatable); ok {
		r = v.Equal(o.Head)
	} else {
		r = c.Head == o.Head
	}
	return
}

func (c *ConsCell) Equal(o interface{}) (r bool) {
	if c == nil {
		r = o == nil
	} else {
		switch o := o.(type) {
		case *ConsCell:			r = o != nil && c.equal(*o)
		case ConsCell:			r = c.equal(o)
		default:				r = c.equal(ConsCell{ Head: o })
		}
	}
	return
}

func (c ConsCell) String() (t string) {
	return fmt.Sprint(c.Head)
}

func (c *ConsCell) Car() (v interface{}) {
	if c != nil {
		v = c.Head
	}
	return
}

func (c *ConsCell) Cdr() (v *ConsCell) {
	if c != nil {
		v = c.Tail
	}
	return
}

func (c *ConsCell) Rplaca(i interface{}) {
	if c != nil {
		c.Head = i
	} else {
		*c = ConsCell{ Head: i }
	}
}

func (c *ConsCell) Rplacd(next *ConsCell) {
	if c != nil {
		c.Tail = next
	} else {
		*c = *next
	}
}