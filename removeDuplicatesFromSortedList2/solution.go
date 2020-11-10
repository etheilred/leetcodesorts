/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	root := &ListNode{0, head}
	p := root
	for p != nil && p.Next != nil {
		if p.Next.Next != nil && p.Next.Val == p.Next.Next.Val {
			for p.Next.Next != nil && p.Next.Val == p.Next.Next.Val {
				p.Next.Next = p.Next.Next.Next
			}
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return root.Next
}