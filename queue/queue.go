package main

import (
	"fmt"
)

func main() {
	q := &Queue{}
	fmt.Println(q.generateBinaryNumbers(100))
}

type Queue struct {
	length int
	front  *ListNode
	rear   *ListNode
}

type ListNode struct {
	data string
	next *ListNode
}

func (ln *Queue) Length() int {
	return ln.length
}

func (ln *Queue) isEmpty() bool {
	return ln.Length() == 0
}

func (q *Queue) enqueue(data string) {
	temp := &ListNode{data: data}
	if q.isEmpty() {
		q.front = temp
	} else {
		q.rear.next = temp
	}

	q.rear = temp
	q.length++
}

func (q *Queue) dequeue() string {
	if q.isEmpty() {
		panic("is not possible to dequeue an empty queue")
	}

	result := q.front.data
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	q.length--
	return result
}
