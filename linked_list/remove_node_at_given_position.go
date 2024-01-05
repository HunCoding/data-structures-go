package linked_list

func (ll *LinkedList) removeNodeGivenPosition(position int) {
	if position <= 0 {
		return
	}

	if position == 1 {
		ll.head = ll.head.next
		return
	}

	count := 1
	previous := ll.head

	for count < position-1 {
		previous = previous.next
		count++
	}

	current := previous.next
	previous.next = current.next
	current = nil
}
