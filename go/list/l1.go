package list

type ListNode struct {
	Val  int
	Next *ListNode
}


func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	smallHead := small
	large := &ListNode{}
	largeHead := large
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	large.Next = nil
	small.Next = largeHead.Next
	return smallHead.Next
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := new(ListNode)
	l3 := l
	var i, v1, v2, sum int
	for l1 != nil || l2 != nil {
		l3.Next = &ListNode{0, nil}
		l3 = l3.Next
		if l1 == nil {
			v1 = 0
		} else {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			v2 = 0
		} else {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum = v1 + v2 + i
		if sum > 9 {
			l3.Val = sum % 10
			i = 1
		} else {
			l3.Val = sum
			i = 0
		}

	}
	return l.Next
}
