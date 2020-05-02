/**
leetcode-876

 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
        if (fast == nil || fast.Next == nil) {
            break
        }
        slow = slow.Next;
        fast = fast.Next.Next;
	}
	return slow
}