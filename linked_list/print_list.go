package linked_list

import "fmt"

func (ll *LinkedList) printList() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}

	fmt.Printf("null\n")
}
