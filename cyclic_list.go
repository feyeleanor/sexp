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

//	A declarative method for building CycLists
func Loop(items... interface{}) (c CycList) {
	c = CycList{}
	if len(items) > 0 {
		c.node = &Node{}
		n := &Node{ Tail: c.node }
		for i := len(items); i > 0; {
			i--
			n.Head = items[i]
			n = &Node{ Tail: n }
		}
		c.node.Head = items[0]
		c.node.Tail = n.Tail.Tail
	}
	return
}

type CycList struct {
	node *Node
}

//	The empty list is represented by a CycList containing a nil pointer to a Node
func (c CycList) IsNil() bool {
	return c.node == nil
}

//	Return the number of chained elements in the list
func (c CycList) Len() (i int) {
	if !c.IsNil() {
		for n := c.node; n.Tail != c.node; n = n.Tail {
			i++
		}
		i++
	}
	return
}

//	Iterate over all elements of the list
//	The only way to terminate iteration is by raising a panic() in the applied function
func (c CycList) Each(f func(interface{})) {
	if !c.IsNil() {
		for n := c.node; ; n = n.Tail {
			f(n.Head)
		}
	}
}

// Return the value stored at the given offset from the start of the list
func (c CycList) At(i int) (r interface{}, ok bool) {
	if !c.IsNil() {
		var n		*Node
		for n = c.node; i > 0 && n.Tail != c.node; i-- {
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
		for n = c.node; i > 0 && n.Tail != c.node; i-- {
			n = n.Tail
		}
		if i == 0 && n != nil {
			n.Head = v
		}
	}
}

//	Return a Cyclist with the next item in the current list as its start
func (c CycList) Next() (n CycList) {
	if !c.IsNil() {
		n.node = c.node.Tail
	}
	return
}

// Return a Cyclist with the last concrete item of the current list as its start
func (c CycList) End() (n CycList) {
	if !c.IsNil() {
		for n.node = c.node; n.node.Tail != c.node; n.node = n.node.Tail {}
	}
	return
}

//	
func (c *CycList) Append(v interface{}) {
	if c.IsNil() {
		c.node = &Node{ Head: v }
		c.node.Tail = c.node
	} else {
		var n	*Node
		for n = c.node; n.Tail != c.node; n = n.Tail {}
		n.Tail = &Node{ Head: v, Tail: c.node }
	}
}

//	Determines if another object is identical to the CycList
func (c CycList) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *CycList:		r = c.Equal(*o)
	case CycList:		r = reflect.DeepEqual(c, o)
	default:			r = c.node.Equal(o)
	}
	return 
}

//	Produces a human-readable representation for the CycList
func (c CycList) String() (t string) {
	if !c.IsNil() {
		terms := []string{ fmt.Sprintf("%v", c.node.Head) }
		for n := c.node.Tail; n != c.node; n = n.Tail {
			terms = append(terms, fmt.Sprintf("%v", n.Head))
		}
		terms = append(terms, "...")
		t = strings.Join(terms, " ")
		t = strings.Replace(t, "()", "nil", -1)
		t = strings.Replace(t, "<nil>", "nil", -1)
	}
	return "(" + t + ")"
}

//	Calculates the nesting of elements within the CycList
func (c CycList) Depth() (d int) {
	if !c.IsNil() {
		if v, ok := c.node.Head.(Nested); ok {
			if r := v.Depth() + 1; r > d {
				d = r
			}
		}
		for n := c.node.Tail; n != c.node; n = n.Tail {
			if v, ok := n.Head.(Nested); ok {
				if r := v.Depth() + 1; r > d {
					d = r
				}
			}
		}
	}
	return
}

//	Reverses the order in which elements of a CycList are traversed
func (c *CycList) Reverse() {
	if !c.IsNil() {
		var n	*Node
		current := &Node{ Head: c.node.Head, Tail: c.node.Tail}
		start := current
		for ; current != c.node; {
			next := current.Tail
			current.Tail = n
			n = current
			current = next
		}
		start.Tail = n
		*c = CycList{ node: n }
	}
}

//	Flatten reduces all Flattenable items in the CycList to their flattest form.
//	In the case of LinearList items these will be spliced inline into the CycList.
func (c *CycList) Flatten() {
	if !c.IsNil() {
		n := &Node{ Head: c.node.Head, Tail: c.node.Tail }
		for ; n != c.node; n = n.Tail {
			switch v := n.Head.(type) {
			case *LinearList:			v.Flatten()
										e := v.End()
										e.Tail = n.Tail
										n.Head = v.Head
										n.Tail = v.Tail
			case Flattenable:			v.Flatten()
			}
		}
	}
}