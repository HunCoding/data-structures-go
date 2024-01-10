package circular_linked_list

type CircularLinkedList struct {
	Length int
	Last   *ListNode
}

type ListNode struct {
	Data int
	Next *ListNode
}
