package sexp

import "fmt"
import "reflect"
import "strings"

type cachedNode struct {
	ListNode
	index	int
}

func (c cachedNode) Update(i int, node ListNode) {
	c.index = i
	c.ListNode = node
}

func (c cachedNode) Clear() {
	c.index = 0
	c.ListNode = nil
}

func (c cachedNode) ClosestNode(i int) (node ListNode, offset int) {
	if i > c.index {
		return c.ListNode, c.index
	}
	return
}


type ListHeader struct {
	nodeType	reflect.Type
	start 		ListNode
	end			ListNode
	cache		cachedNode
	length		int
}


func NewListHeader(n ListNode) ListHeader {
	t := reflect.TypeOf(n)
	if t.Kind() != reflect.Ptr {
		t = reflect.PtrTo(t)
	}
	return ListHeader{ nodeType: t }
}

func (l ListHeader) newListNode() ListNode {
	return reflect.New(l.nodeType.Elem()).Interface().(ListNode)
}

func (l ListHeader) NewListNode(value interface{}) (n ListNode) {
	n = l.newListNode()
	n.Store(CURRENT_NODE, value)
	return
}

func (l ListHeader) EnforceBounds(start, end *int) (ok bool) {
	if *start < 0 {
		*start = 0
	}

	if *end > l.length - 1 {
		*end = l.length - 1
	}

	if *end >= *start {
		ok = true
	}
	return
}

func (l *ListHeader) Clear() {
	l.start = nil
	l.end = nil
	l.length = 0
	l.cache.Clear()
}

func (l ListHeader) String() (t string) {
	terms := []string{}
	l.Each(func(term interface{}) {
		terms = append(terms, fmt.Sprintf("%v", term))
	})
	if l.length > 0 && l.start == NextNode(l.end) {
		terms = append(terms, "...")
	}
	t = strings.Join(terms, " ")
	t = strings.Replace(t, "()", "nil", -1)
	t = strings.Replace(t, "<nil>", "nil", -1)
	return "(" + t + ")"
}

func (l ListHeader) Len() (c int) {
	return l.length
}

func (l ListHeader) Depth() (d int) {
	l.Each(func(v interface{}) {
		if v, ok := v.(Nested); ok {
			if r := v.Depth() + 1; r > d {
				d = r
			}
		}
	})
	return
}

func (l ListHeader) Start() ListNode {
	return l.start
}

func (l ListHeader) End() ListNode {
	return l.end
}

func (l ListHeader) Clone() (r *ListHeader) {
	r = &ListHeader{ nodeType: l.nodeType }
	l.Each(func(v interface{}) { r.Append(v) })
	return
}

func (l *ListHeader) Expand(i, n int) {
	if i > -1 && i <= l.length {
		switch {
		case l == nil:					fallthrough
		case i == l.length:				for ; n > 0; n-- {
											l.Append(l.newListNode())
										}

		case i == 0:					l.length = n
										for ; n > 0; n-- {
											x := l.newListNode()
											x.Link(NEXT_NODE, l.start)
											l.start = x
										}

		default:						x1 := l.findNode(i - 1)
										x2 := l.findNode(i)
										l.length += n
										for ; n > 0; n-- {
											x1.Link(NEXT_NODE, l.newListNode())
											x1 = NextNode(x1)
										}
										x1.Link(NEXT_NODE, x2)
		}
	}
}

func (l ListHeader) Each(f func(interface{})) {
	n := l.start
	for i := l.length; i > 0; i-- {
		f(n.Content())
		n = NextNode(n)
	}
}

func (l ListHeader) equal(o ListHeader) (r bool) {
	if l.length == o.length {
		r = true
		n := l.start
		x := o.start
		for i := l.length; r && i > 0; i-- {
			if r = n != nil && n.Equal(x); r {
				n = NextNode(n)
				x = NextNode(x)
			}
		}
	}
	return
}

func (l ListHeader) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *ListHeader:	r = o != nil && l.equal(*o)
	case ListHeader:	r = l.equal(o)
	}
	return
}

func (l *ListHeader) eachNode(f func(int, ListNode)) {
	n := l.start
	for i := 0; i < l.length; i++ {
		f(i, n)
		n = NextNode(n)
	}
}

func (l ListHeader) findNode(i int) (n ListNode) {
	switch {
	case i == 0:				n = l.start
	case i == l.length - 1:		n = l.end
	default:					start, offset := l.cache.ClosestNode(i)
								if start == nil {
									start = l.start
								}
								if n = start.MoveTo(i - offset); n != nil {
									l.cache.Update(i, n)
								}
	}
	return
}

func (l ListHeader) At(i int) (r interface{}) {
	if n := l.findNode(i); n != nil {
		r = n.Content()
	}
	return
}

func (l ListHeader) Set(i int, v interface{}) {
	if n := l.findNode(i); n != nil {
		n.Store(CURRENT_NODE, v)
	}
}

func (l *ListHeader) Append(v interface{}) {
	switch {
	case l.start == nil:	l.start = l.NewListNode(v)
							l.end = l.start

	default:				tail := NextNode(l.end)
							l.end.Link(NEXT_NODE, l.NewListNode(v))
							l.end = NextNode(l.end)
							l.end.Link(NEXT_NODE, tail)
	}
	l.length++
}

func (l *ListHeader) AppendSlice(s Slice) {
	length := s.Len()
	if length > 0 {
		l.Append(s[0])
		if length > 1 {
			tail := NextNode(l.end)
			for _, v := range s[1:] {
				l.end.Link(NEXT_NODE, l.NewListNode(v))
				l.end = NextNode(l.end)
			}
			l.end.Link(NEXT_NODE, tail)
			l.length += length - 1
		}
	}
}

//	Iterates through the list reducing the nesting of each element which can be flattened.
//	Elements which are themselves LinearLists will be inlined as part of the containing list and their contained list destroyed.
func (l *ListHeader) Flatten() {
	l.eachNode(func(i int, n ListNode) {
		value := n.Content()
		if h, ok := value.(Flattenable); ok {
			h.Flatten()
		}

		if h, ok := value.(Linkable); ok {
			switch length := h.Len(); {
			case length == 0:		n.Store(CURRENT_NODE, nil)

			case length == 1:		n.Store(CURRENT_NODE, h.Start().Content())

			default:				l.length += length - 1
									h.End().Link(NEXT_NODE, NextNode(n))
									n.Link(CURRENT_NODE, h.Start())
									if n == l.start {
										l.start = h.Start()
									}

									if n == l.end {
										l.end = h.End()
									}
			}
		} else {
			n.Store(CURRENT_NODE, value)
		}
	})
}

func (l ListHeader) Compact() *Slice {
	s := make(Slice, l.Len(), l.Len())
	i := 0
	l.Each(func(v interface{}) {
		s[i] = v
		i++
	})
	return &s
}

func (l *ListHeader) reverseLinks() (r ListNode) {
	if l != nil {
		current := l.start
		l.end = current

		for i := l.length; i > 0; i-- {
			next := NextNode(current)
			current.Link(NEXT_NODE, r)
			r = current
			current = next				
		}
	}
	return
}

//	Reverses the order in which elements of a List are traversed
func (l *ListHeader) Reverse() {
	l.start = l.reverseLinks()
}

func (l ListHeader) Head() (r interface{}) {
	if l.start != nil {
		l.start.Content()
	}
	return
}

func (l *ListHeader) Tail() {
	if n := l.start; n != nil {
		l.start = NextNode(n)
		n.Link(NEXT_NODE, nil)
		l.length--
	}
}