/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	length := 0
	p := head
	for p != nil {
		p = p.Next
		length++
	}
	q := mergeSort(head, length)
	// fmt.Println(listToSlice(q))
	return q
}

func mergeSort(head *ListNode, length int) *ListNode {
	if head.Next == nil {
		return head
	}
	q := head
	prev := head
	for i := 0; i < length/2; i++ {
		prev = q
		q = q.Next
	}
	prev.Next = nil

	p := mergeSort(head, length/2)
	q = mergeSort(q, length-length/2)
	// fmt.Printf("p: %v\n", listToSlice(p))
	// fmt.Printf("q: %v\n", listToSlice(q))
	// fmt.Println("==========")
	var t *ListNode
	if p.Val > q.Val {
		p, q = q, p
	}
	t = p
	for p.Next != nil && q != nil {
		// fmt.Printf("Here1: p=%v q=%v\n", listToSlice(p), listToSlice(q))
		for p.Next != nil && p.Next.Val < q.Val {
			p = p.Next
		}
		if p.Next != nil && p.Next.Val >= q.Val {
			tmp := p.Next
			p.Next = q
			q = q.Next
			p = p.Next
			p.Next = tmp
		}
	}
	// fmt.Println("here2")
	for q != nil {
		// fmt.Printf("here3: p=%v, q=%v\n", listToSlice(p), listToSlice(q))
		tmp := p.Next
		p.Next = q
		q = q.Next
		p = p.Next
		p.Next = tmp
	}
	// fmt.Println(listToSlice(t))
	// fmt.Printf("t: %v\n", listToSlice(t))
	// fmt.Println("==========")
	return t
}
