package main

func (ll *LinkedList) findLength() (count int) {
	current := ll.head
	for current != nil {
		count++
		current = current.next
	}

	return
}
