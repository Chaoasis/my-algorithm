package list

type ListNode struct {
	Val  int
	Next *ListNode
}

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	f := func(head *ListNode, n int) *ListNode {
		pre := head
		for i := 0; i < n; i++ {
			pre = pre.Next
		}
		for pre != nil {
			pre = pre.Next
			head = head.Next
		}
		return head
	}
	//这里因为是获得倒数第几个节点,所以增加头节点不会影响到倒数节点的顺序
	dump := &ListNode{}
	dump.Next = head
	node := f(dump, n+1)
	node.Next = node.Next.Next
	return dump.Next
}

// 合并K个升序链表
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：链表数组如下：
// [
// 1->4->5,
// 1->3->4,
// 2->6
// ]
// 将它们合并到一个有序链表中得到。
// 1->1->2->3->4->4->5->6
func mergeKLists(lists []*ListNode) *ListNode {
	dump := &ListNode{}
	head := dump
	m := make(map[int]*ListNode)
	count := 0
	for _, e := range lists {
		if e == nil {
			break
		}
		kv, ok := m[count]
		if !ok {
			m[count], kv = e, e
		}
		if kv.Val > e.Val {
			m[count] = e
		}
	}
	if m[count] != nil {
		dump.Next = m[count]
		dump = dump.Next
		m[count] = dump.Next
		delete(m, count)
	}
	return head.Next
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
