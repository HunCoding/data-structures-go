package main

import "fmt"

func (ll *LinkedList) printList() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data, " --> ")
		current = current.next
	}
	fmt.Println("null")
}

type LinkedList struct {
	head *Node
}

type Node struct {
	data string
	next *Node
}

func main() {
	nodes := &LinkedList{
		head: &Node{
			data: "The",
			next: &Node{
				data: "Huncoding",
				next: &Node{
					data: "programming",
					next: &Node{
						data: "channel",
						next: nil,
					},
				},
			},
		},
	}
	nodes.insertAtGivenPosition("youtube", 4)
	nodes.printList()
}
