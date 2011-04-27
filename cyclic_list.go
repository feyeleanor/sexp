package sexp

import "fmt"
import "reflect"

func CList(items... interface{}) (c *CycList) {
	n := &Node{}
	c = &CycList{ n }
	tails := len(items) - 1
	for i, v := range items {
		if v == nil {
			n.Head = &Node{}
		} else {
			n.Head = v
		}
		if i < tails {
			n.Tail = &Node{}
		}
		n = n.Tail
	}
	return
}

type CycList struct {
	*Node
}

func (c *CycList) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *CycList:		r = reflect.DeepEqual(c, o)
	case CycList:		r = reflect.DeepEqual(*c, o)
	default:			r = c.Node.Equal(o)
	}
	return 
}

func (c *CycList) Each(f func(interface{})) {
	visited_nodes := make(memo)
	for n := c.Node; n != nil; n = n.Tail {
		if !visited_nodes.Memorise(n) {
			break
		}
		f(n.Head)
	}
}

func (c *CycList) _string(visited_nodes memo) (t string) {
	if !c.IsNil() {
		for n := c.Node; n != nil; n = n.Tail {
			visited_nodes.Memorise(n)
			if len(t) > 0 {
				t += " "
			}
			switch h := n.Head.(type) {
			case nil:				t += "nil"
			case *CycList:			switch {
									case h.Node == c.Node:
										t += "(...)"
									case visited_nodes.Find(h.Node) != nil:
										t += printAddress(h.Node)
									default:
										if term := h._string(visited_nodes); term == "()" {
											t += "nil"
										} else {
											t += term
										}
									}
			case *Node:				if term := (&CycList{ h })._string(visited_nodes); term == "()" {
										t += "nil"
									} else {
										t += term
									}
			case Addressable:		t += printAddress(h)
			default:				t += fmt.Sprintf("%v", h)
			}
			if visited_nodes.Find(n.Tail) != nil {
				if c.Node == n.Tail {
					t += " ..."
				} else {
					t += " " + printAddress(n.Tail) + " head = " + printAddress(c.Node)
				}
				break
			}
		}
		visited_nodes.Forget(c)
	}
	return "(" + t + ")"
}

func (c *CycList) String() (t string) {
	return c._string(make(memo))
}

func (c *CycList) Len() (i int) {
	if !c.IsNil() {
		visited_nodes := make(memo)
		for n := c.Node; n != nil; n = n.Tail {
			if visited_nodes.Memorise(n) {
				i++
			} else {
				panic(i)
			}
		}
	}
	return
}

func (c *CycList) depth(visited_nodes memo) (d int) {
	if visited_nodes.Memorise(c) {
		for n := c.Node; n != nil; n = n.Tail {
			if v, ok := n.Head.(CyclicNested); ok {
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

func (c *CycList) Depth() int {
	return c.depth(make(memo)) - 1
}

func (c *CycList) Reverse() {
	if c.Tail != nil {
		var n	*Node
		current := &Node{ Head: c.Head, Tail: c.Tail}
		for ; current != nil; {
			next := current.Tail
			current.Tail = n
			n = current
			current = next				
		}
		*c = CycList{ n }
	}
}

func (c *CycList) Flatten() {
	
}

func (c *CycList) At(i int) (r interface{}) {
	var n	*Node
	for n = c.Node; i > 0 && n.Tail != nil; i-- {
		n = n.Tail
	}
	if i == 0 {
		r = n.Head
	}
	return
}

func (c *CycList) Set(i int, v interface{}) {
	var n	*Node
	for n = c.Node; i > 0 && n.Tail != nil; i-- {
		n = n.Tail
	}
	if n != nil {
		n.Head = v
	}
}

func (c *CycList) Link(to, from int) (ok bool) {
	var target, source	*Node

	for n, i := c.Node, 0; n != nil; n = n.Tail {
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

func (c *CycList) End() (n *Node) {
	visited_nodes := make(memo)
	for n = c.Node; visited_nodes.Memorise(n) && n.Tail != nil; n = n.Tail {}
	return
}