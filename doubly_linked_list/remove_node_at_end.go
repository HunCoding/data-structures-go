package doubly_linked_list

func (dl *DoublyLinkedList) removeNodeAtEnd() {
	if dl.Length == 0 {
		return
	}

	if dl.Head == dl.Tail {
		dl.Tail = nil
		dl.Head = nil
		return
	}

	dl.Tail.Previous.Next = nil
	dl.Tail.Previous, dl.Tail = nil, dl.Tail.Previous

	dl.Length--
}
