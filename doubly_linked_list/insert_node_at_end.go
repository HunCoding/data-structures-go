package doubly_linked_list

func (dl *DoublyLinkedList) insertNodeAtEnd(data int) {
	newNode := &ListNode{
		Data: data,
	}

	if dl.Length == 0 {
		dl.Head = newNode
	} else {
		dl.Tail.Next = newNode
		newNode.Previous = dl.Tail
	}

	dl.Tail = newNode
	dl.Length++
}
