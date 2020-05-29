/**
leetcode-141
给定一个链表，判断链表中是否有环。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 
如果 pos 是 -1，则在该链表中没有环。


 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 //哈希表也是个思路,每遍历一个节点,先检查是否已存入 地址和ture ,map[*ListNode]bool
 //如果存在环，则快慢指针会必然相遇
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
        return false
	}
	slow := head
    fast := head.Next  //人为先制造不相等
    for slow != fast {
        if (fast == nil || fast.Next == nil) {
            return false
        }
        slow = slow.Next;
        fast = fast.Next.Next;
    }
    return true
}