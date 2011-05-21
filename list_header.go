package sexp

import "fmt"
import "strings"

type ListHeader struct {
	start 	*Node
	end		*Node
	length	int
}

func (l *ListHeader) Clear() {
	l.start = nil
	l.end = nil
	l.length = 0
}

func (l ListHeader) IsNil() bool {
	return l.start == nil && l.end == nil && l.length == 0
}

func (l ListHeader) NotNil() bool {
	return l.start != nil || l.end != nil || l.length > 0
}

//	Produces a human-readable representation for the CycList
func (l ListHeader) String() (t string) {
	if l.NotNil() {
		terms := []string{}
		l.Each(func(term interface{}) {
			terms = append(terms, fmt.Sprintf("%v", term))
		})
		if l.start == l.end.Tail {
			terms = append(terms, "...")
		}
		t = strings.Join(terms, " ")
		t = strings.Replace(t, "()", "nil", -1)
		t = strings.Replace(t, "<nil>", "nil", -1)
	}
	return "(" + t + ")"
}

func (l ListHeader) Len() (c int) {
	if l.NotNil() {
		c = l.length
	}
	return
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

func (l ListHeader) Start() *Node {
	return l.start
}

func (l ListHeader) End() *Node {
	return l.end
}

func (l ListHeader) Clone() (r *ListHeader) {
	r = &ListHeader{}
	l.Each(func(v interface{}) { r.Append(v) })
	return
}

func (l ListHeader) Each(f func(interface{})) {
	if l.NotNil() {
		n := l.start
		for i := 0; i < l.length; i++ {
			f(n.Head)
			n = n.Tail
		}
	}
}

func (l ListHeader) Equal(o ListHeader) (r bool) {
	switch {
	case l.IsNil():				r = o.IsNil()
	case l.length == o.length:	r = true
								n := l.start
								x := o.start
								for i := l.length; r && i > 0; i-- {
									if r = n.Equal(x); r {
										n = n.Tail
										x = x.Tail
									}
								}
	}
	return
}

func (l *ListHeader) eachNode(f func(*Node)) {
	if l.NotNil() {
		n := l.start
		for i := 0; i < l.length; i++ {
			f(n)
			n = n.Tail
		}
	}
}

func (l ListHeader) At(i int) (r interface{}) {
	if l.NotNil() {
		if n := l.start.MoveTo(i); n != nil {
			r = n.Head
		}
	}
	return
}

func (l ListHeader) Set(i int, v interface{}) {
	if l.NotNil() {
		if n := l.start.MoveTo(i); n != nil {
			n.Head = v
		}
	}
}

func (l *ListHeader) Append(v interface{}) {
	if l.IsNil() {
		l.start = &Node{ Head: v }
		l.end = l.start
		l.length = 1
	} else {
		l.end.Tail = &Node{ Head: v }
		l.end = l.end.Tail
		l.length++
	}
}

func (l *ListHeader) AppendSlice(s []interface{}) {
	if len(s) > 0 {
		if l.IsNil() {
			l.Append(s[0])
			s = s[1:]
		}
		for _, v := range s {
			l.end.Tail = &Node{ Head: v }
			l.end = l.end.Tail
		}
		l.length += len(s)
	}
}

//	Iterates through the list reducing the nesting of each element which can be flattened.
//	Elements which are themselves LinearLists will be inlined as part of the containing list and their contained list destroyed.
func (l *ListHeader) Flatten() {
	l.eachNode(func(n *Node) {
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

			case n == l.end:		l.end = end
									n.Head = start.Head
									n.Tail = start.Tail
									l.length += length - 1

			default:				end.Tail = n.Tail
									n.Head = start.Head
									n.Tail = start.Tail
									l.length += length - 1
			}
			h.Clear()
		}
	})
}

func (l *ListHeader) reverseLinks() (r *Node) {
	if l.NotNil() {
		current := l.start
		l.end = current

		for i := l.length; i > 0; i-- {
			next := current.Tail
			current.Tail = r
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
	if l.NotNil() {
		r = l.start.Head
	}
	return
}

func (l *ListHeader) Tail() {
	if l.NotNil() {
		n := l.start
		l.start = l.start.Tail
		n.Tail = nil
		l.length--
	}
}