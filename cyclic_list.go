package sexp

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

func (c CycList) Clone() (r *CycList) {
	r = &CycList{ *c.ListHeader.Clone() }
	if r.NotNil() {
		r.end.Tail = r.start
	}
	return
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
	c.ListHeader.Append(v)
	c.end.Tail = c.start
}

func (c *CycList) AppendSlice(s []interface{}) {
	c.ListHeader.AppendSlice(s)
	if c.end != nil {
		c.end.Tail = c.start
	}
}

//	Determines if another object is equivalent to the CycList
//	Two CycLists are identical if they both have the same number of nodes, and the head of each node is the same
func (c CycList) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *CycList:		r = c.ListHeader.Equal(o.ListHeader)
	case CycList:		r = c.ListHeader.Equal(o.ListHeader)
	default:			r = c.start.Equal(o)
	}
	return 
}

//	Reverses the order in which elements of a CycList are traversed
func (c *CycList) Reverse() {
	if r := c.reverseLinks(); r != nil {
		c.start.Tail = r
		c.start = r
		c.end.Tail = c.start
	}
}

//	Flatten reduces all Flattenable items in the CycList to their flattest form.
//	In the case of LinearList items these will be spliced inline into the CycList.
func (c *CycList) Flatten() {
	if c.NotNil() {
		n := &ConsCell{ Head: c.start.Head, Tail: c.start.Tail }
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

func (c *CycList) Tail() {
	c.ListHeader.Tail()
	c.end = c.start
}