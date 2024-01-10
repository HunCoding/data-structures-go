package circular_linked_list

import "fmt"

func (cl *CircularLinkedList) printList() {
	if cl.Length == 0 {
		return
	}

	first := cl.Last.Next
	for first != cl.Last {
		fmt.Print(first.Data, " -> ")
		first = first.Next
	}
	fmt.Printf("%d -> null\n", first.Data)
}
