package doubly_linked_list

func (dl *DoublyLinkedList) removeNodeAtBeginning() {
	if dl.Length == 0 {
		return
	}

	temp := dl.Head
	if dl.Head == dl.Tail {
		dl.Tail = nil
	} else {
		dl.Head.Next.Previous = nil
	}

	dl.Head = dl.Head.Next
	temp.Next = nil
}
