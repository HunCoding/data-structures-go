package main

func main() {
	list3 := &ListNode{
		Data:     3,
		Next:     nil,
		Previous: nil,
	}
	list2 := &ListNode{
		Data:     2,
		Next:     list3,
		Previous: nil,
	}
	list1 := &ListNode{
		Data:     1,
		Next:     list2,
		Previous: nil,
	}

	list3.Previous = list2
	list2.Previous = list1

	ll := &DoublyLinkedList{
		Head: list1,
		Tail: list3,
	}

	ll.printListBackward()
}
