/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 //leetcode-206. 反转链表
 //递归
 func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
        return head
    }
	first := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return first
}
//迭代
func reverseList(head *ListNode) *ListNode {
    // corner
    if head == nil || head.Next == nil {
        return head
	}
    // revise
    var pre *ListNode
    for head != nil {
        temp := head.Next
        head.Next = pre
        pre = head
        head = temp
    }
    return pre
}