package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	carry := 0

	for l1 != nil || l2 != nil {
		x := func() int {
			if l1 != nil {
				return l1.Val
			}

			return 0
		}()

		y := func() int {
			if l2 != nil {
				return l2.Val
			}

			return 0
		}()

		sum := carry + x + y
		carry = sum / 10
		tail.Next = &ListNode{Val: sum % 10}
		tail = tail.Next
		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}
