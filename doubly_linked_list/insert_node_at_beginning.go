package main

import "fmt"

func main() {
	dl := &DoublyLinkedList{}
	dl.insertNodeAtBeginning(1)
	dl.insertNodeAtBeginning(2)
	dl.insertNodeAtBeginning(3)
	dl.printListBackward()
}

func (dl *DoublyLinkedList) insertNodeAtBeginning(data int) {
	newNode := &ListNode{
		Data: data,
	}

	if dl.Length == 0 {
		dl.Tail = newNode
	} else {
		dl.Head.Previous = newNode
	}

	newNode.Next = dl.Head
	dl.Head = newNode
	dl.Length++
}

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

func (ll *DoublyLinkedList) printList() {
	current := ll.Head

	for current != nil {
		fmt.Print(current.Data, " -> ")
		current = current.Next
	}

	fmt.Printf("null\n")
}

func (ll *DoublyLinkedList) printListBackward() {
	current := ll.Tail

	for current != nil {
		fmt.Print(current.Data, " -> ")
		current = current.Previous
	}

	fmt.Printf("null\n")
}
