package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	n := 0
	dummy := &ListNode{}
	dummy.Next = head
	for head != nil {
		head = head.Next
		n++
	}
	n = n / 2
	head = dummy
	for i := 0; i < n; i++ {
		head = head.Next
	}
	head.Next = head.Next.Next

	return dummy.Next
}
func main() {

}
