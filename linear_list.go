package sexp

/*
	A LinearList is a finitely-terminated list structure.
	Each node in the list may point to exactly one other node in the list.
	The terminating node does not point to any other node.
	No node may be pointed to by more than one other node in the list.
	There are no nil links between nodes in the list.
*/

func List(items... interface{}) (l *LinearList) {
	l = new(LinearList)
	l.AppendSlice((Slice)(items))
	return
}

type LinearList struct {
	ListHeader
}

func (l LinearList) End() ListNode {
	return l.end
}

func (l LinearList) Clone() *LinearList {
	return &LinearList{ *l.ListHeader.Clone() }
}

//	Determines if another object is equivalent to the LinearList
//	Two LinearLists are identical if they both have the same number of nodes, and the head of each node is the same
func (l LinearList) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *LinearList:	r = o != nil && l.ListHeader.Equal(o.ListHeader)
	case LinearList:	r = l.ListHeader.Equal(o.ListHeader)
	default:			r = l.start.Equal(o)
	}
	return 
}

//	Removes all elements in the range from the list.
func (l *LinearList) Delete(from, to int) {
	if l != nil && from >= 0 && to < l.length && from <= to {
		last_element_index := l.length - 1
		switch {
		case from == 0:
			switch {
			case to == 0:
				l.start = NextNode(l.start)
				l.length -= 1
			case to == last_element_index:
				l.start = nil
				l.end = nil
				l.length = 0
			default:
				l.start = l.start.MoveTo(to + 1)
				l.length -= to + 1
			}

		case from == to:
			s := l.start.MoveTo(from - 1)
			e := s.MoveTo(1)
			s.Link(NEXT_NODE, NextNode(e))
			l.length -= 1

		case from == last_element_index:
			l.end = l.start.MoveTo(from - 1)
			l.end.Link(NEXT_NODE, nil)
			l.length -= 1

		case to == last_element_index:
			l.end = l.start.MoveTo(from - 1)
			l.end.Link(NEXT_NODE, nil)
			l.length = from

		default:
			e := l.start.MoveTo(from - 1)
			e.Link(NEXT_NODE, e.MoveTo(to - from + 2))
			l.length -= to - from + 1
		}
	}
}

//	Removes the elements in the range from the current list and returns a new list containing them.
func (l *LinearList) Cut(from, to int) (r LinearList) {
	if l != nil && from >= 0 && to < l.length && from <= to {
		last_element_index := l.length - 1
		switch {
		case from == 0:
			switch {
			case to == 0:
				r.start = l.start
				r.end = r.start
				r.length = 1

				l.start = NextNode(l.start)
				l.length -= 1
			case to == last_element_index:
				r.start = l.start
				r.end = l.end
				r.length = l.length

				l.start = nil
				l.end = nil
				l.length = 0
			default:
				r.start = l.start
				r.end = r.start.MoveTo(to)
				l.start = NextNode(r.end)
				r.end.Link(NEXT_NODE, nil)
				r.length = to + 1
				l.length -= r.length
			}

		case from == to:
			s := l.start.MoveTo(from - 1)
			r.start = NextNode(s)
			r.end = r.start
			r.length = 1
			s.Link(NEXT_NODE, NextNode(r.end))
			r.end.Link(NEXT_NODE, nil)
			l.length -= 1

		case from == last_element_index:
			l.end = l.start.MoveTo(from - 1)
			l.end.Link(NEXT_NODE, nil)
			l.length -= 1

		case to == last_element_index:
			l.end = l.start.MoveTo(from - 1)
			r.start = NextNode(l.end)
			r.end = r.start
			l.end.Link(NEXT_NODE, nil)
			r.length = to - from + 1
			l.length = from

		default:
			e := l.start.MoveTo(from - 1)
			r.start = NextNode(e)
			r.end = r.start.MoveTo(to - from)
			if r.end != nil {
				e.Link(NEXT_NODE, NextNode(r.end))
			} else {
				e.Link(NEXT_NODE, nil)
			}
			r.length = to - from + 1
			l.length -= r.length
		}
	}
	return
}

//	Insert an item into the list at the given location.
func (l *LinearList) Insert(i int, o interface{}) {
	
}

//	Take all the elements from another list and insert them into this list, destroying the other list if successful.
func (l *LinearList) Absorb(i int, o *LinearList) (ok bool) {
	switch {
	case i < 0:						fallthrough
	case i > l.length:				fallthrough
	case o == nil:					return false

	case l == nil:					l.start = o.start
									l.end = o.end

	case i == 0:					o.end.Link(NEXT_NODE, l.start)
									l.start = o.start

	case i == l.length:				l.end.Link(NEXT_NODE, o.start)
									l.end = o.end

	default:						n := l.start.MoveTo(i - 1)
									o.end.Link(NEXT_NODE, NextNode(n))
									n.Link(NEXT_NODE, o.start)
	}
	l.length += o.length
	o.Clear()
	return true
}