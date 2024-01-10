package linked_list

func (ll *LinkedList) findWhereLoopStarts() *Node {
	fastPointer := ll.head
	slowPointer := ll.head

	for fastPointer != nil && fastPointer.next != nil {
		fastPointer = fastPointer.next.next
		slowPointer = slowPointer.next

		if fastPointer == slowPointer {
			break
		}
	}

	if slowPointer != fastPointer {
		return nil
	}

	temp := ll.head
	for temp != slowPointer {
		temp = temp.next
		slowPointer = slowPointer.next
	}

	return temp
}
