package main

func (ll *LinkedList) removeNodeAtEnd() {
	if ll.head == nil || ll.head.next == nil {
		return
	}

	var previous *Node
	current := ll.head.next

	for current.next != nil {
		previous = current
		current = current.next
	}

	previous.next = nil
	current = nil
}
