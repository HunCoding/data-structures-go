package doubly_linked_list

import "fmt"

func (dl *DoublyLinkedList) printList() {
	current := dl.Head

	for current != nil {
		fmt.Print(current.Data, " -> ")
		current = current.Next
	}

	fmt.Printf("null\n")
}

func (dl *DoublyLinkedList) printListBackward() {
	current := dl.Tail

	for current != nil {
		fmt.Print(current.Data, " -> ")
		current = current.Previous
	}

	fmt.Printf("null\n")
}
