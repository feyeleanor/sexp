package sexp

import "fmt"
import "reflect"
import "unsafe"

func ConsNil() *ConsCell {
	return &ConsCell{}
}

func Cons(head interface{}, tail *ConsCell) (c *ConsCell) {
	if tail == nil {
		c = &ConsCell{ Head: head }
	} else {
		c = &ConsCell{ head, &ConsCell{ Head: tail } }
	}
	return
}

func List(items... interface{}) (c *ConsCell) {
	c = &ConsCell{}
	n := c
	tails := len(items) - 1
	for i, v := range items {
		if v == nil {
			n.Head = ConsNil()
		} else {
			n.Head = v
		}
		if i < tails {
			n.Tail = &ConsCell{}
		}
		n = n.Tail
	}
	return
}

type ConsCell struct {
	Head		interface{}
	Tail		*ConsCell
}

func (c *ConsCell) Cons(x interface{}) *ConsCell {
	return &ConsCell{ Head: x, Tail: c }
}

func (c ConsCell) IsNil() bool {
	return (c.Head == nil) && (c.Tail == nil)
}

func (c *ConsCell) Addr() uintptr {
	return uintptr(unsafe.Pointer(c))
}

func (c *ConsCell) Each(f func(interface{})) {
	visited_nodes := make(memo)
	for n := c; n != nil; n = n.Tail {
		if !visited_nodes.Memorise(n) {
			break
		}
		f(n.Head)
	}
}

func (c *ConsCell) _string(visited_nodes memo) (t string) {
	if !c.IsNil() {
		for n := c; n != nil; n = n.Tail {
			visited_nodes.Memorise(n)
			if len(t) > 0 {
				t += " "
			}
			switch h := n.Head.(type) {
			case nil:				t += "nil"
			case *ConsCell:			switch {
									case h == c:
										t += "(...)"
									case visited_nodes.Find(h) != nil:
										t += printAddress(h)
									default:
										if term := h._string(visited_nodes); term == "()" {
											t += "nil"
										} else {
											t += term
										}
									}
			case Addressable:		t += printAddress(h)
			default:				t += fmt.Sprintf("%v", h)
			}
			if visited_nodes.Find(n.Tail) != nil {
				if c == n.Tail {
					t += " ..."
				} else {
					t += " " + printAddress(n.Tail)
				}
				break
			}
		}
		visited_nodes.Forget(c)
	}
	return "(" + t + ")"
}

func (c *ConsCell) String() (t string) {
	return c._string(make(memo))
}

func (c *ConsCell) Len() (i int, recursive bool) {
	if !c.IsNil() {
		visited_nodes := make(memo)
		for n := c; n != nil; n = n.Tail {
			if visited_nodes.Memorise(n) {
				i++
			} else {
				recursive = true
				break
			}
		}
	}
	return
}

func (c *ConsCell) depth(visited_nodes memo) (d int) {
	if visited_nodes.Memorise(c) {
		for n := c; n != nil; n = n.Tail {
			if v, ok := n.Head.(Nested); ok {
				if r := v.depth(visited_nodes); r > d {
					d = r
				}
			}
		}
		visited_nodes.Forget(c)
	}
	d++
	return
}

func (c *ConsCell) Depth() int {
	return c.depth(make(memo)) - 1
}

func (c ConsCell) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *ConsCell:			r = reflect.DeepEqual(c, *o)
	case ConsCell:			r = reflect.DeepEqual(c, o)
	}
	return 
}

func (c *ConsCell) Reverse() {
	if c.Tail != nil {
		var n	*ConsCell
		current := &ConsCell{ Head: c.Head, Tail: c.Tail}
		for ; current != nil; {
			next := current.Tail
			current.Tail = n
			n = current
			current = next
		}
		*c = *n
	}
}

func (c *ConsCell) Flatten() {
	
}

func (c *ConsCell) At(i int) (n interface{}) {
	return
}

func (c *ConsCell) Set(i int, n interface{}) {
	
}

func (c *ConsCell) Link(to, from int) (ok bool) {
	var target, source	*ConsCell

	for n, i := c, 0; n != nil; n = n.Tail {
		if i == to		{ target = n }
		if i == from	{ source = n }
		i++
	}
	if source != nil && target != nil {
		source.Tail = target
		ok = true
	}
	return
}

func (c *ConsCell) End() (n *ConsCell) {
	visited_nodes := make(memo)
	for n = c; visited_nodes.Memorise(n) && n.Tail != nil; n = n.Tail {}
	return
}