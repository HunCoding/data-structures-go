package doubly_linked_list

type DoublyLinkedList struct {
	Head   *ListNode
	Tail   *ListNode
	Length int
}

type ListNode struct {
	Next     *ListNode
	Previous *ListNode
	Data     int
}
