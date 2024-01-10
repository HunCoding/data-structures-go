package circular_linked_list

func (cl *CircularLinkedList) removeNodeAtEnd() {
	if cl.Last == nil || cl.Length == 0 {
		return
	}

	if cl.Last == cl.Last.Next {
		cl.Last = nil
	} else {
		first := cl.Last.Next
		for first.Next != cl.Last {
			first = first.Next
		}

		first.Next, cl.Last = cl.Last.Next, first
	}

	cl.Length--
}
