package sexp

import "fmt"
//import "reflect"
import "strings"

/*
	A CycList is a circular list structure.
	Each node in the list may point to exactly one other node in the list.
	No node may be pointed to by more than one other node in the list.
	There are no nil links between nodes in the list.
*/

//	A declarative method for building CycLists
func Loop(items... interface{}) (c *CycList) {
	c = new(CycList)
	c.AppendSlice(items)
	return
}

type CycList struct {
	start 	*Node
	end		*Node
	length	int
}

//	The empty list is represented by a CycList containing a nil pointer to a Node
func (c CycList) IsNil() bool {
	return c.start == nil || c.length == 0
}

//	Return the number of chained elements in the list
func (c CycList) Len() (i int) {
	return c.length
}

//	Iterate over all elements of the list
//	The only way to terminate iteration is by raising a panic() in the applied function
func (c CycList) Each(f func(interface{})) {
	if !c.IsNil() {
		for n := c.start; ; n = n.Tail {
			f(n.Head)
		}
	}
}

// Return the value stored at the given offset from the start of the list
func (c CycList) At(i int) (r interface{}, ok bool) {
	if !c.IsNil() {
		var n		*Node
		i = i % c.length
		for n = c.start; i > 0 && n != c.end; i-- {
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
		i = i % c.length
		for n = c.start; i > 0 && n.Tail != c.start; i-- {
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
		n.start = c.start.Tail
		n.end = c.end.Tail
		n.length = c.length
	}
	return
}

//	
func (c *CycList) Append(v interface{}) {
	if c.IsNil() {
		c.start = &Node{ Head: v }
		c.start.Tail = c.start
		c.end = c.start
		c.length = 1
	} else {
		c.end.Tail = &Node{ Head: v, Tail: c.start }
		c.end = c.end.Tail
		c.length++
	}
}

func (c *CycList) AppendSlice(s []interface{}) {
	if len(s) > 0 {
		if c.IsNil() {
			c.Append(s[0])
			s = s[1:]
		}
		for _, v := range s {
			c.end.Tail = &Node{ Head: v, Tail: c.start }
			c.end = c.end.Tail
		}
		c.length += len(s)
	}
}

func (c CycList) equal(o CycList) (r bool) {
	switch {
	case c.IsNil():
		r = o.IsNil()
	case c.Len() == o.Len():
		r = true
		n := c.start
		x := o.start
		for i := 0; r && i < c.Len(); i++ {
			if r = n.Equal(x); r {
				n = n.Tail
				x = x.Tail
			}
		}
	}
	return
}

//	Determines if another object is identical to the CycList
//	Two CycLists are identical if they both have the same number of nodes, and the head of each node is the same
func (c CycList) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *CycList:		r = c.equal(*o)
	case CycList:		r = c.equal(o)
	default:			r = c.start.Equal(o)
	}
	return 
}

//	Produces a human-readable representation for the CycList
func (c CycList) String() (t string) {
	if !c.IsNil() {
		terms := []string{ fmt.Sprintf("%v", c.start.Head) }
		for n := c.start.Tail; n != c.start; n = n.Tail {
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
		if v, ok := c.start.Head.(Nested); ok {
			if r := v.Depth() + 1; r > d {
				d = r
			}
		}
		for n := c.start.Tail; n != c.start; n = n.Tail {
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
		var result	*Node

		current := c.start
		c.end = current

		for ; current != nil; {
			next := current.Tail
			current.Tail = result
			result = current
			current = next
			if current == c.start {
				break
			}
		}
		c.start.Tail = result
		c.start = result
		c.end.Tail = c.start
	}
}

//	Flatten reduces all Flattenable items in the CycList to their flattest form.
//	In the case of LinearList items these will be spliced inline into the CycList.
func (c *CycList) Flatten() {
	if !c.IsNil() {
		n := &Node{ Head: c.start.Head, Tail: c.start.Tail }
		for ; n != c.start; n = n.Tail {
			switch v := n.Head.(type) {
			case LinearList:			v.Flatten()
										e := v.node.End()
										e.Tail = n.Tail
										n.Head = v.node.Head
										n.Tail = v.node.Tail
										c.length += v.length - 1
			case Flattenable:			v.Flatten()
			}
		}
	}
}