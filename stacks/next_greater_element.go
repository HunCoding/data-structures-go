package main

//func nextGreaterElement(arr []int) []int {
//	n := len(arr)
//	res := make([]int, n)
//	stack := Stack{make([]int, 0)}
//	for i := n - 1; i >= 0; i-- {
//		if !stack.isEmpty() {
//			for !stack.isEmpty() && stack.peek() <= arr[i] {
//				stack.pop()
//			}
//		}
//
//		if stack.isEmpty() {
//			res[i] = -1
//		} else {
//			res[i] = stack.peek()
//		}
//
//		stack.push(arr[i])
//	}
//	return res
//}
