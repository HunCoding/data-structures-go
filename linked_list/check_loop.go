package linked_list

func (ll *LinkedList) checkLoop() bool {
	fastPointer := ll.head
	slowPointer := ll.head

	for fastPointer != nil && fastPointer.next != nil {
		fastPointer = fastPointer.next.next
		slowPointer = slowPointer.next
		if slowPointer == fastPointer {
			return true
		}
	}

	return false
}
