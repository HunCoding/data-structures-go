package main

func (ll *LinkedList) insertAtGivenPosition(data string, position int) {
	newNode := &Node{data: data}
	if position == 1 {
		newNode.next = ll.head
		ll.head = newNode
		return
	}

	previous := ll.head
	count := 1

	for count < position-1 {
		previous = previous.next
		count++
	}

	current := previous.next
	newNode.next = current
	previous.next = newNode
}
