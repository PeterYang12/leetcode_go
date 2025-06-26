package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	var nums []int
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		nums = append(nums, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}
	ans := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		ans = max(nums[n-1-i]+slow.Val, ans)
		slow = slow.Next
	}
	return ans
}
func pairSum2(head *ListNode) int {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//	reverse
	prev := (*ListNode)(nil)
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}
	ans := 0
	for prev != nil {
		ans = max(ans, prev.Val+head.Val)
		prev = prev.Next
		head = head.Next
	}
	return ans
}
func makeList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for i := 0; i < len(nums); i++ {
		curr.Next = &ListNode{nums[i], nil}
		curr = curr.Next
	}
	return dummy.Next
}
func main() {
	nums := []int{5, 4, 2, 1}
	head := makeList(nums)
	fmt.Println(pairSum2(head))
}
