package main

func (ll *LinkedList) insertAtBeginning(data string) {
	newNode := &Node{data: data, next: ll.head}
	ll.head = newNode
}
