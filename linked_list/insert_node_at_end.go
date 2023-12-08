package main

func (ll *LinkedList) insertAtEnd(data string) {
	newNode := &Node{data: data, next: nil}
	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}
