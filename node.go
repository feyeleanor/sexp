package sexp

type ListNode interface {
	Equatable
	MoveTo(int) ListNode
	Content() interface{}
	Link(int, ListNode) bool
	Set(int, interface{}) bool
}

const(
	PREVIOUS_NODE = -1
	CURRENT_NODE = 0
	NEXT_NODE = 1
)

func NextNode(l ListNode) (r ListNode) {
	if l != nil {
		r = l.MoveTo(NEXT_NODE)
	}
	return
}

func PreviousNode(l ListNode) ListNode {
	return l.MoveTo(PREVIOUS_NODE)
}

func LastElement(l ListNode) (r ListNode) {
	if r = l; r != nil {
		for {
			if n := NextNode(r); n != nil {
				r = n
			} else {
				break
			}
		}
	}
	return
}

func FirstElement(l ListNode) (r ListNode) {
	if r = l; r != nil {
		for {
			if n := PreviousNode(r); n != nil {
				r = n
			} else {
				break
			}
		}
	}
	return
}