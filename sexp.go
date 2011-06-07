package sexp

// Insert n elements at position x in Indexable i.
func Expand(i Indexable, x, n int) (b bool) {
	var capacity, length		int

	//	ensure that for capacity constrained types there's enough underlying capacity
	//	for lists we insert additional nodes at the correct location and update the list length
	//	for linkable types we insert additional nodes at the correct location
	switch block := i.(type) {
	case Resizeable:			capacity = block.Cap()
								length = block.Len()
								end := l + n
								if end > capacity {
									block.Reallocate(end)
								}

	case FixedSize:				capacity = block.Cap()
								length = block.Len()
								if length + n > capacity {
									return
								}

	case Expandable:			block.Expand(x, n)
								b = true

	case Appendable:			for count := n; count > 0; count-- {
									block.Append(nil)
								}
								b = true
	}

	if !b {
		//	for block types we now have to create a "hole" where the new cells have been inserted
		switch block := i.(type) {
		case Blitter:			i.BlockCopy(x + n, x, n)
								i.BlockClear(x, n)

		case Indexable:			for end := x + n; x < end; x++ {
									i.Set(x + n, i.At(x))
									i.Set(x, nil)
								}
		}


		// make a hole
		for j := len0 - 1; j >= i; j-- {
			a[j+n] = a[j]
		}

		*p = a
	}
}