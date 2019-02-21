package add_two_numbers

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp1 := l1
	tmp2 := l2
	carry := 0

	head := new(ListNode)
	current := head

	for {
		v1 := 0
		v2 := 0

		if tmp1 != nil {
			v1 = tmp1.Val
			tmp1 = tmp1.Next
		}

		if tmp2 != nil {
			v2 = tmp2.Val
			tmp2 = tmp2.Next
		}

		sum := carry + v1 + v2
		carry = sum / 10
		current.Next = &ListNode{sum % 10, nil}
		current = current.Next

		if tmp1 == nil && tmp2 == nil {
			break
		}
	}

	if carry != 0 {
		current.Next = &ListNode{carry, nil}
	}

	return head.Next
}
