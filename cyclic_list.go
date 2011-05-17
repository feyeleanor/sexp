package sexp

import "fmt"
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
	ListHeader
}

//	Iterate over all elements of the list once
func (c CycList) Each(f func(interface{})) {
	if c.NotNil() {
		f(c.start.Head)
		for n := c.start.Tail; n != c.start; n = n.Tail {
			f(n.Head)
		}
	}
}

//	Iterate over all elements of the list indefinitely
//	The only way to terminate iteration is by raising a panic() in the applied function
func (c CycList) Cycle(f func(interface{})) {
	if c.NotNil() {
		for n := c.start; ; n = n.Tail {
			f(n.Head)
		}
	}
}

// Return the value stored at the given offset from the start of the list
func (c CycList) At(i int) interface{} {
	return c.ListHeader.At(i % c.length)
}

// Set the value stored at the given offset from the start of the list
func (c CycList) Set(i int, v interface{}) {
	c.ListHeader.Set(i % c.length, v)
}

func (c *CycList) Advance() {
	if c.NotNil() {
		start := c.start.Tail
		end := c.end.Tail
		c.start = start
		c.end = end
	}
	return
}

func (c *CycList) Rotate(i int) {
	if c.NotNil() {
		if i %= c.length; i > 0 {
			c.end = c.start.MoveTo(i - 1)
			c.start = c.end.Tail
		}
	}
}

func (c *CycList) Append(v interface{}) {
	if c.IsNil() {
		c.start = &Node{ Head: v }
		c.start.Tail = c.start
		c.end = c.start
	} else {
		c.end.Tail = &Node{ Head: v, Tail: c.start }
		c.end = c.end.Tail
	}
	c.length++
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
	case c.IsNil():				r = o.IsNil()
	case c.Len() == o.Len():	r = true
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

//	Determines if another object is equivalent to the CycList
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
	if c.NotNil() {
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
	if c.NotNil() {
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
	if c.NotNil() {
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
	if c.NotNil() {
		n := &Node{ Head: c.start.Head, Tail: c.start.Tail }
		for ; n != c.start; n = n.Tail {
			switch v := n.Head.(type) {
			case *LinearList:			v.Flatten()
										e := v.start.End()
										e.Tail = n.Tail
										n.Head = v.start.Head
										n.Tail = v.start.Tail
										c.length += v.length - 1
			case Flattenable:			v.Flatten()
			}
		}
	}
}