package kaka

import (
	"math"
	"regexp"
	"strings"
	"unicode"
)

func Solution(str string) []string {
	if len(str)&'1' == 0 {
		str += "_"
	}
	var arr []string
	for i := 0; i < len(str); i += 2 {
		arr = append(arr, str[i:i+2])
	}
	return arr
}

func Solution2(str string) []string {
	return regexp.MustCompile(".{2}").FindAllString(str+"_", -1)
}

func SquareSum(numbers []int) int {
	var result int
	for i, i2 := range numbers {
		pow := math.Pow(float64(i2), float64(i))
		result += int(pow)
	}
	return result
}

func EachCons(arr []int, n int) [][]int {
	var z [][]int
	for i := 0; i < len(arr); i += n {
		x := []int(arr[i : i+n])
		z = append(z, x)
	}
	return z
}

func ToAlternatingCase(str string) string {
	f := func(r rune) rune {
		if unicode.IsUpper(r) {
			return unicode.ToLower(r)
		}
		return unicode.ToUpper(r)
	}
	return strings.Map(f, str)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func Ll(nums []int) *ListNode {
	head := &ListNode{}
	fnode := head
	for _, num := range nums {
		temp := ListNode{Val: num}
		head.Next = &temp
		head = &temp
	}
	return fnode.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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
