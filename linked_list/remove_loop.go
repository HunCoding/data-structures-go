package linked_list

func (ll *LinkedList) removeLoop() {
	fastPointer := ll.head
	slowPointer := ll.head

	for fastPointer != nil && fastPointer.next != nil {
		fastPointer = fastPointer.next.next
		slowPointer = slowPointer.next

		if fastPointer == slowPointer {
			temp := ll.head

			for temp.next != slowPointer.next {
				temp = temp.next
				slowPointer = slowPointer.next
			}

			slowPointer.next = nil
		}
	}
}
