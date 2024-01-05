package linked_list

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	mergedNodes := &ListNode{}
	tail := mergedNodes

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}

		tail = tail.Next
	}

	if list1 == nil {
		tail.Next = list2
	} else {
		tail.Next = list1
	}

	return mergedNodes.Next
}
