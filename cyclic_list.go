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
	c.AppendSlice((Slice)(items))
	return
}

type CycList struct {
	ListHeader
}

func (c CycList) Clone() (r *CycList) {
	r = &CycList{ *c.ListHeader.Clone() }
	if !r.IsNil() {
		r.end.Tail = r.start
	}
	return
}

//	Iterate over all elements of the list indefinitely
//	The only way to terminate iteration is by raising a panic() in the applied function
func (c CycList) Cycle(f func(interface{})) {
	if !c.IsNil() {
		for n := c.start; ; n = n.Tail {
			f(n.Head)
		}
	}
}

// Return the value stored at the given offset from the start of the list
func (c CycList) At(i int) interface{} {
	i %= c.length
	if i < 0 {
		i = c.length + i
	}
	return c.ListHeader.At(i)
}

// Set the value stored at the given offset from the start of the list
func (c CycList) Set(i int, v interface{}) {
	i %= c.length
	if i < 0 {
		i = c.length + i
	}
	c.ListHeader.Set(i, v)
}

func (c *CycList) Advance() {
	if !c.IsNil() {
		start := c.start.Tail
		end := c.end.Tail
		c.start = start
		c.end = end
	}
	return
}

func (c *CycList) Rotate(i int) {
	if !c.IsNil() {
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

func (c *CycList) AppendSlice(s Slice) {
	c.ListHeader.AppendSlice(s)
	if c.end != nil {
		c.end.Tail = c.start
	}
}

//	Determines if another object is equivalent to the CycList
//	Two CycLists are identical if they both have the same number of nodes, and the head of each node is the same
func (c CycList) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *CycList:		r = o != nil && c.ListHeader.Equal((*o).ListHeader)
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

//	Iterates through the list reducing the nesting of each element which can be flattened.
//	Elements which are themselves LinearLists will be inlined as part of the containing list and their contained list destroyed.
func (c *CycList) Flatten() {
	c.eachConsCell(func(n *ConsCell) {
		if h, ok := n.Head.(Flattenable); ok {
			h.Flatten()
		}

		if h, ok := n.Head.(Linkable); ok {
			start := h.Start()
			end := h.End()

			length := h.Len()
			switch {
			case start == nil:		fallthrough
			case length == 0:		n.Head = nil

			case length == 1:		n.Head = start.Head

			case n == c.end:		end.Tail = c.start
									c.end = end
									n.Head = start.Head
									n.Tail = start.Tail
									c.length += length - 1

			default:				end.Tail = n.Tail
									n.Head = start.Head
									n.Tail = start.Tail
									c.length += length - 1
			}
		}
	})
}


func (c *CycList) Tail() {
	c.ListHeader.Tail()
	c.end = c.start
}