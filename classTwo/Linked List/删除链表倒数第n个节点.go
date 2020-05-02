/**
leetcode-19
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 //要找到倒数第n+1个节点，利用假头确实方便些
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{}
	result.Next = head
	
	slow, fast:= result, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil{
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	
	return result.Next
}
