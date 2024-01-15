package stacks

type StackLinkedList struct {
	Top    *StackNode
	Length int
}

type StackNode struct {
	Data int
	Next *StackNode
}

func (sl *StackLinkedList) push(data int) {
	newNode := &StackNode{Data: data}

	newNode.Next = sl.Top
	sl.Top = newNode
	sl.Length++
}

func (sl *StackLinkedList) pop() int {
	result := sl.Top.Data
	sl.Top = sl.Top.Next
	sl.Length--

	return result
}
