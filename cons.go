package sexp

import "fmt"
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

func (c *ConsCell) _string(visited_nodes memo) (t string) {
	if c.Head != nil || c.Tail != nil {
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
										t += "cons(" + printAddress(h) + ")"
									default:
										if term := h._string(visited_nodes); term == "()" {
											t += "nil"
										} else {
											t += term
										}
									}
			case Addressable:		t += "blob(" + printAddress(h) + ")"
			default:				t += fmt.Sprintf("%v", h)
			}
			if visited_nodes.Find(n.Tail) != nil {
				if c == n.Tail {
					t += " ..."
				} else {
					t += " cons(" + printAddress(n.Tail) + ")"
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