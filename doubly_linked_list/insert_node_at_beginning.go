package doubly_linked_list

func (dl *DoublyLinkedList) insertNodeAtBeginning(data int) {
	newNode := &ListNode{
		Data: data,
	}

	if dl.Length == 0 {
		dl.Tail = newNode
	} else {
		dl.Head.Previous = newNode
	}

	newNode.Next = dl.Head
	dl.Head = newNode
	dl.Length++
}
