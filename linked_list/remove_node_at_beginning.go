package linked_list

func (ll *LinkedList) removeNodeAtBeginning() {
	if ll.head == nil {
		return
	}

	ll.head = ll.head.next
	return
}
