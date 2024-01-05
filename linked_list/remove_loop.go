package linked_list

func (ll *LinkedList) removeLoop() {
	fastPointer := ll.head
	slowPointer := ll.head

	for fastPointer != nil && fastPointer.next != nil {
		fastPointer = fastPointer.next.next
		slowPointer = slowPointer.next

		if slowPointer == fastPointer {
			temp := ll.head
			for slowPointer.next != temp.next {
				temp = temp.next
				slowPointer = slowPointer.next
			}

			slowPointer.next = nil
		}
	}
}
