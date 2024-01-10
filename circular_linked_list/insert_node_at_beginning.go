package circular_linked_list

func (cl *CircularLinkedList) insertNodeAtBeginning(data int) {
	newNode := &ListNode{Data: data}

	if cl.Length == 0 || cl.Last == nil {
		cl.Last = newNode
	} else {
		newNode.Next = cl.Last.Next
	}

	cl.Last.Next = newNode
	cl.Length++
}
