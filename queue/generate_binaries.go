package main

func (q *Queue) generateBinaryNumbers(n int) []string {
	result := make([]string, n)

	q.enqueue("1")
	for i := 0; i < n; i++ {
		result[i] = q.dequeue()
		n1 := result[i] + "0"
		n2 := result[i] + "1"
		q.enqueue(n1)
		q.enqueue(n2)
	}

	return result
}
