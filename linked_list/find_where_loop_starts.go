package linked_list

func (ll *LinkedList) findWhereLoopStarts() *Node {
	fastPointer := ll.head
	slowPointer := ll.head

	for fastPointer != nil && fastPointer.next != nil {
		fastPointer = fastPointer.next.next
		slowPointer = slowPointer.next

		if slowPointer == fastPointer {
			temp := ll.head
			for slowPointer != temp {
				temp = temp.next
				slowPointer = slowPointer.next
			}

			return temp
		}
	}

	return nil
}
