package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func ConvertSliceToListNode(sl []int) *ListNode {
	if len(sl) == 0 {
		return nil
	}
	node := &ListNode{Val: sl[0]}
	node.Next = ConvertSliceToListNode(sl[1:])
	return node
}

func ConvertListNodeToSlice(node *ListNode) []int {
	var res []int
	if node == nil {
		return res
	}
	for node != nil {
		res = append(res, node.Val)
		node = node.Next
	}
	return res
}

func BuildListWithCycle(vals []int, pos int) (*ListNode, *ListNode) {
	if len(vals) == 0 {
		return nil, nil
	}

	nodes := make([]*ListNode, len(vals))
	for i, v := range vals {
		nodes[i] = &ListNode{Val: v}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}

	var cycleStart *ListNode
	if pos >= 0 {
		cycleStart = nodes[pos]
		nodes[len(nodes)-1].Next = cycleStart
	}

	return nodes[0], cycleStart
}
