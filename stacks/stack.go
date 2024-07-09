package main

type StackLinkedList struct {
	Top    *StackNode
	Length int
}

type StackNode struct {
	Data *ListNode
	Next *StackNode
}

func (sl *StackLinkedList) isEmpty() bool {
	return sl.Length == 0
}

func (sl *StackLinkedList) push(data *ListNode) {
	newNode := &StackNode{Data: data}

	newNode.Next = sl.Top
	sl.Top = newNode
	sl.Length++
}

func (sl *StackLinkedList) pop() *ListNode {
	result := sl.Top.Data
	sl.Top = sl.Top.Next
	sl.Length--

	return result
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	for i := 0; i < len(s)-1; i++ {
		for j := i; j < len(s)-1; j++ {

		}
	}

	return 0
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxP := 0
	left, right := 0, 1

	for right < len(prices) {
		if prices[left] < prices[right] {
			maxP = max(maxP, prices[right]-prices[left])
		} else {
			left = right
		}

		right += 1
	}

	return maxP
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func romanToInt(s string) int {
	/*
		I can be placed before V (5) and X (10) to make 4 and 9.
		X can be placed before L (50) and C (100) to make 40 and 90.
		C can be placed before D (500) and M (1000) to make 400 and 900.
	*/
	romanValues := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return romanValues[s]
	}

	total := 0
	for i := 1; i < len(s)-1; i++ {
		if value, ok := romanValues[string(s[i-1]+s[i])]; ok {
			total += value
			continue
		}

		total += romanValues[string(s[i-1])]
	}

	return total
}

func reverseNumber(num int) int {
	sign := 1
	if num < 0 {
		sign = -1
		num *= -1
	}

	res := 0
	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10
	}

	res *= sign

	return res
}
