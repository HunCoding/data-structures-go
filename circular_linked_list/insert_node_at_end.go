package circular_linked_list

func (cl *CircularLinkedList) insertNodeAtEnd(data int) {
	newNode := &ListNode{Data: data}

	if cl.Last == nil || cl.Length == 0 {
		cl.Last = newNode
		cl.Last.Next = cl.Last
	} else {
		newNode.Next, cl.Last.Next = cl.Last.Next, newNode
		cl.Last = newNode
	}

	cl.Length++
}
