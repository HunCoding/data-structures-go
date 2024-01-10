package circular_linked_list

func (cl *CircularLinkedList) removeNodeAtBeginning() {
	if cl.Last == nil || cl.Length == 0 {
		return
	}

	temp := cl.Last.Next
	if cl.Last.Next == cl.Last {
		cl.Last = nil
	} else {
		cl.Last.Next = temp.Next
	}

	temp.Next = nil
	cl.Length--
}
