// https://leetcode.com/problems/balanced-binary-tree
// explanation: https://www.youtube.com/watch?v=t1XPK-ONKwM

package tree

import "testing"

func dfsBalancedTree(node *TreeNode) (depth int, isBalanced bool) {
	if node == nil {
		return 0, true
	}

	leftDepth, leftBalanced := dfsBalancedTree(node.Left)
	rightDepth, rightBalanced := dfsBalancedTree(node.Right)

	maxDepth := leftDepth
	if leftDepth < rightDepth {
		maxDepth = rightDepth
	}

	diff := leftDepth - rightDepth

	if leftBalanced && rightBalanced && (-1 <= diff && diff <= 1) {
		return maxDepth + 1, true
	}
	return 0, false
}

func isBalanced(root *TreeNode) (isBalanced bool) {
	_, isBalanced = dfsBalancedTree(root)
	return
}

func TestIsBalanced(t *testing.T) {
	node := func(val int, left, right *TreeNode) *TreeNode {
		return &TreeNode{Val: val, Left: left, Right: right}
	}

	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			name: "empty tree",
			root: nil,
			want: true,
		},
		{
			name: "single node",
			root: node(1, nil, nil),
			want: true,
		},
		{
			name: "example balanced",
			root: node(3,
				node(9, nil, nil),
				node(20, node(15, nil, nil), node(7, nil, nil)),
			),
			want: true,
		},
		{
			name: "example unbalanced",
			root: node(1,
				node(2,
					node(3, node(4, nil, nil), node(4, nil, nil)),
					node(3, nil, nil),
				),
				node(2, nil, nil),
			),
			want: false,
		},
		{
			name: "perfect tree",
			root: node(1,
				node(2, node(4, nil, nil), node(5, nil, nil)),
				node(3, node(6, nil, nil), node(7, nil, nil)),
			),
			want: true,
		},
		{
			name: "left skewed two nodes",
			root: node(1,
				node(2, nil, nil),
				nil,
			),
			want: true,
		},
		{
			name: "left skewed three nodes",
			root: node(1,
				node(2,
					node(3, nil, nil),
					nil,
				),
				nil,
			),
			want: false,
		},
		{
			name: "right skewed three nodes",
			root: node(1,
				nil,
				node(2,
					nil,
					node(3, nil, nil),
				),
			),
			want: false,
		},
		{
			name: "difference one at root still balanced",
			root: node(1,
				node(2, node(4, nil, nil), nil),
				node(3, nil, nil),
			),
			want: true,
		},
		{
			name: "root balanced but subtree not balanced",
			root: node(1,
				node(2,
					node(3,
						node(4, nil, nil),
						nil,
					),
					nil,
				),
				node(2, nil, nil),
			),
			want: false,
		},
		{
			name: "balanced with missing internal nodes",
			root: node(1,
				node(2, nil, node(4, nil, nil)),
				node(3, node(5, nil, nil), nil),
			),
			want: true,
		},
		{
			name: "unbalanced deep on right subtree",
			root: node(1,
				node(2, nil, nil),
				node(3,
					nil,
					node(4,
						nil,
						node(5, nil, nil),
					),
				),
			),
			want: false,
		},
		{
			name: "negative values balanced",
			root: node(-1,
				node(-2, node(-4, nil, nil), node(-5, nil, nil)),
				node(-3, nil, nil),
			),
			want: true,
		},
		{
			name: "zero values unbalanced",
			root: node(0,
				node(0,
					node(0,
						node(0, nil, nil),
						nil,
					),
					nil,
				),
				node(0, nil, nil),
			),
			want: false,
		},
		{
			name: "balanced larger tree",
			root: node(8,
				node(4,
					node(2, nil, nil),
					node(6, nil, nil),
				),
				node(12,
					node(10, nil, nil),
					node(14, nil, nil),
				),
			),
			want: true,
		},
		{
			name: "unbalanced larger tree",
			root: node(8,
				node(4,
					node(2,
						node(1, nil, nil),
						nil,
					),
					node(6, nil, nil),
				),
				node(12, nil, nil),
			),
			want: false,
		},
		{
			name: "balanced chain at lower subtree boundary",
			root: node(1,
				node(2,
					node(3, nil, nil),
					nil,
				),
				node(4, nil, nil),
			),
			want: true,
		},
		{
			name: "duplicate values balanced",
			root: node(1,
				node(1, node(1, nil, nil), nil),
				node(1, nil, node(1, nil, nil)),
			),
			want: true,
		},
		{
			name: "duplicate values unbalanced",
			root: node(1,
				node(1,
					node(1,
						node(1, nil, nil),
						nil,
					),
					nil,
				),
				node(1, nil, nil),
			),
			want: false,
		},
		{
			name: "one subtree empty other height one",
			root: node(1,
				node(2, nil, nil),
				nil,
			),
			want: true,
		},
		{
			name: "one subtree empty other height two",
			root: node(1,
				node(2, node(3, nil, nil), nil),
				nil,
			),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isBalanced(tt.root)
			if got != tt.want {
				t.Fatalf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
