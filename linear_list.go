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
	l.AppendSlice(items)
	return
}

type LinearList struct {
	ListHeader
}

func (l LinearList) Clone() *LinearList {
	return &LinearList{ *l.ListHeader.Clone() }
}

//	Determines if another object is equivalent to the LinearList
//	Two CycLists are identical if they both have the same number of nodes, and the head of each node is the same
func (l LinearList) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *LinearList:	r = l.ListHeader.Equal(o.ListHeader)
	case LinearList:	r = l.ListHeader.Equal(o.ListHeader)
	default:			r = l.start.Equal(o)
	}
	return 
}

//	Removes all elements in the range from the list.
func (l *LinearList) Delete(from, to int) {
	if l.NotNil() && from >= 0 && to < l.length && from <= to {
		last_element_index := l.length - 1
		switch {
		case from == 0:
			switch {
			case to == 0:
				l.start = l.start.Tail
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
			s.Tail = e.Tail
			l.length -= 1

		case from == last_element_index:
			l.end = l.start.MoveTo(from - 1)
			l.end.Tail = nil
			l.length -= 1

		case to == last_element_index:
			l.end = l.start.MoveTo(from - 1)
			l.end.Tail = nil
			l.length = from

		default:
			e := l.start.MoveTo(from - 1)
			e.Tail = e.MoveTo(to - from + 2)
			l.length -= to - from + 1
		}
	}
}

//	Removes the elements in the range from the current list and returns a new list containing them.
func (l *LinearList) Cut(from, to int) (r LinearList) {
	if l.NotNil() && from >= 0 && to < l.length && from <= to {
		last_element_index := l.length - 1
		switch {
		case from == 0:
			switch {
			case to == 0:
				r.start = l.start
				r.end = r.start
				r.length = 1

				l.start = l.start.Tail
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
				l.start = r.end.Tail
				r.end.Tail = nil
				r.length = to + 1
				l.length -= r.length
			}

		case from == to:
			s := l.start.MoveTo(from - 1)
			r.start = s.Tail
			r.end = r.start
			r.length = 1
			s.Tail = r.end.Tail
			r.end.Tail = nil
			l.length -= 1

		case from == last_element_index:
			l.end = l.start.MoveTo(from - 1)
			l.end.Tail = nil
			l.length -= 1

		case to == last_element_index:
			l.end = l.start.MoveTo(from - 1)
			r.start = l.end.Tail
			r.end = r.start
			l.end.Tail = nil
			r.length = to - from + 1
			l.length = from

		default:
			e := l.start.MoveTo(from - 1)
			r.start = e.Tail
			r.end = r.start.MoveTo(to - from)
			if r.end != nil {
				e.Tail = r.end.Tail
			} else {
				e.Tail = nil
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
	case o == nil:					fallthrough
	case o.IsNil():					return false

	case l.IsNil():					l.start = o.start
									l.end = o.end

	case i == 0:					o.end.Tail = l.start
									l.start = o.start

	case i == l.length:				l.end.Tail = o.start
									l.end = o.end

	default:						n := l.start.MoveTo(i - 1)
									o.end.Tail = n.Tail
									n.Tail = o.start
	}
	l.length += o.length
	o.Clear()
	return true
}

func (l LinearList) Car() (r interface{}) {
	if l.NotNil() {
		r = l.start.Head
	}
	return
}

func (l LinearList) Cdr() (r LinearList) {
	if l.NotNil() {
		r.start = l.start.Tail
		r.end = l.end
		r.length = l.length - 1
	}
	return
}

func (l *LinearList) Rplaca(i interface{}) {
	if l.IsNil() {
		*l = *(List(i))
	} else {
		l.start.Head = i
	}
}

func (l *LinearList) Rplacd(tail *LinearList) {
	if l.IsNil() {
		l.start = tail.start
		l.end = tail.end
		l.length = tail.length
	} else {
		l.start.Tail = tail.start
		l.end = tail.end
		l.length = tail.length + 1
	}
}