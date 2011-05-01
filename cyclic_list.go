package sexp

import "fmt"
import "reflect"
import "strings"

/*
	A CycList is a circular list structure.
	Each node in the list may point to exactly one other node in the list.
	No node may be pointed to by more than one other node in the list.
	There are no nil links between nodes in the list.
*/


func Loop(items... interface{}) (c CycList) {
	c = CycList{}
	if len(items) > 0 {
		c.Node = &Node{}
		n := &Node{ Tail: c.Node }
		for i := len(items); i > 0; {
			i--
			n.Head = items[i]
			n = &Node{ Tail: n }
		}
		c.Head = items[0]
		c.Tail = n.Tail.Tail
	}
	return
}

type CycList struct {
	*Node
}

//	The empty list is represented by a CycList containing an nil pointer to a Node
func (c CycList) IsNil() bool {
	return c.Node == nil
}

//	Return the number of chained elements in the list
func (c CycList) Len() (i int) {
	if !c.IsNil() {
		for n := c.Node; n.Tail != c.Node; n = n.Tail {
			i++
		}
		i++
	}
	return
}

//	Iterate over all elements of the list until an exception is raised in the applied function
func (c CycList) Each(f func(interface{})) {
	if !c.IsNil() {
		for n := c.Node; ; n = n.Tail {
			f(n.Head)
		}
	}
}

// Return the value stored at the given offset from the start of the list
func (c CycList) At(i int) (r interface{}, ok bool) {
	if !c.IsNil() {
		var n		*Node
		for n = c.Node; i > 0 && n.Tail != c.Node; i-- {
			n = n.Tail
		}
		if i == 0 && n != nil {
			r, ok = n.Head, true
		}
	}
	return
}

// Set the value stored at the given offset from the start of the list
func (c CycList) Set(i int, v interface{}) {
	if !c.IsNil() {
		var n	*Node
		for n = c.Node; i > 0 && n.Tail != c.Node; i-- {
			n = n.Tail
		}
		if i == 0 && n != nil {
			n.Head = v
		}
	}
}

func (c CycList) Next() (n CycList) {
	if !c.IsNil() {
		n.Node = c.Tail
	}
	return
}

// Return a Cyclist with the last item of the current list as its start
func (c CycList) End() (n CycList) {
	if !c.IsNil() {
		for n.Node = c.Node; n.Tail != c.Node; n.Node = n.Tail {}
	}
	return
}

func (c CycList) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *CycList:		r = c.Equal(*o)
	case CycList:		r = reflect.DeepEqual(c, o)
	default:			r = c.Node.Equal(o)
	}
	return 
}

func (c CycList) String() (t string) {
	if !c.IsNil() {
		terms := []string{ fmt.Sprintf("%v", c.Head) }
		for n := c.Node.Tail; n != c.Node; n = n.Tail {
			terms = append(terms, fmt.Sprintf("%v", n.Head))
		}
		terms = append(terms, "...")
		t = strings.Join(terms, " ")
		t = strings.Replace(t, "()", "nil", -1)
		t = strings.Replace(t, "<nil>", "nil", -1)
	}
	return "(" + t + ")"
}

func (c CycList) Depth() (d int) {
	if !c.IsNil() {
		if v, ok := c.Head.(Nested); ok {
			if r := v.Depth() + 1; r > d {
				d = r
			}
		}
		for n := c.Tail; n != c.Node; n = n.Tail {
			if v, ok := n.Head.(Nested); ok {
				if r := v.Depth() + 1; r > d {
					d = r
				}
			}
		}
	}
	return
}

func (c *CycList) Reverse() {
	if !c.IsNil() {
		list_head := &Node{ Tail: c.Node }
		reverse_list := CycList{ list_head }
		for n := c.Tail; n != c.Node; n = n.Tail {
			list_head.Head = n.Head
			list_head = &Node{ Tail: list_head }
		}
		list_head.Head = c.Head
		reverse_list.Tail = list_head
		*c = reverse_list
	}
}





func (c *CycList) flatten(visited_nodes memo) (r CycList) {
	var n	*Node

	for n = c.Node; n != nil; n = n.Tail {
		//	iterate nodes and whenever a list node is met expand it
		switch v := n.Head.(type) {
		case LinearList:	v.Flatten()
							n.Head = v.Head
							t := v.End()
							t.Tail = n.Tail
							n.Tail = v.Tail
		case CycList:		v.flatten(visited_nodes)




//							n = append(n, v.flatten(visited_nodes)...)
//					} else {
//						n = append(n, v)
//					}
//		default:		n = append(n, v)
		}
	}
	return r
}

func (c *CycList) Flatten() {
	*c = c.flatten(make(memo))
}