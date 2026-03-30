// https://leetcode.com/problems/symmetric-tree
// explanation: https://www.youtube.com/watch?v=Mao9uzxwvmc

package tree

import "testing"

func dfsSymmetricTree(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && dfsSymmetricTree(left.Left, right.Right) && dfsSymmetricTree(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfsSymmetricTree(root.Left, root.Right)
}

func TestIsSymmetric(t *testing.T) {
	node := func(val int, left, right *TreeNode) *TreeNode {
		return &TreeNode{Val: val, Left: left, Right: right}
	}

	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			name: "single node",
			root: node(1, nil, nil),
			want: true,
		},
		{
			name: "example symmetric",
			root: node(1,
				node(2, node(3, nil, nil), node(4, nil, nil)),
				node(2, node(4, nil, nil), node(3, nil, nil)),
			),
			want: true,
		},
		{
			name: "example asymmetric by structure",
			root: node(1,
				node(2, nil, node(3, nil, nil)),
				node(2, nil, node(3, nil, nil)),
			),
			want: false,
		},
		{
			name: "two equal children",
			root: node(1,
				node(2, nil, nil),
				node(2, nil, nil),
			),
			want: true,
		},
		{
			name: "different child values",
			root: node(1,
				node(2, nil, nil),
				node(3, nil, nil),
			),
			want: false,
		},
		{
			name: "only left child",
			root: node(1,
				node(2, nil, nil),
				nil,
			),
			want: false,
		},
		{
			name: "only right child",
			root: node(1,
				nil,
				node(2, nil, nil),
			),
			want: false,
		},
		{
			name: "symmetric with negative values",
			root: node(0,
				node(-1, node(-2, nil, nil), node(-3, nil, nil)),
				node(-1, node(-3, nil, nil), node(-2, nil, nil)),
			),
			want: true,
		},
		{
			name: "asymmetric with negative values",
			root: node(0,
				node(-1, node(-2, nil, nil), node(-3, nil, nil)),
				node(-1, node(-2, nil, nil), node(-3, nil, nil)),
			),
			want: false,
		},
		{
			name: "symmetric deeper with missing mirrored nodes",
			root: node(1,
				node(2, node(3, nil, nil), nil),
				node(2, nil, node(3, nil, nil)),
			),
			want: true,
		},
		{
			name: "asymmetric deeper because missing node on one side",
			root: node(1,
				node(2, node(3, nil, nil), nil),
				node(2, node(3, nil, nil), nil),
			),
			want: false,
		},
		{
			name: "symmetric all duplicate values",
			root: node(1,
				node(1, node(1, nil, nil), node(1, nil, nil)),
				node(1, node(1, nil, nil), node(1, nil, nil)),
			),
			want: true,
		},
		{
			name: "asymmetric all duplicate values by structure",
			root: node(1,
				node(1, node(1, nil, nil), nil),
				node(1, node(1, nil, nil), nil),
			),
			want: false,
		},
		{
			name: "large symmetric balanced tree",
			root: node(8,
				node(6,
					node(5, node(4, nil, nil), node(7, nil, nil)),
					node(9, nil, nil),
				),
				node(6,
					node(9, nil, nil),
					node(5, node(7, nil, nil), node(4, nil, nil)),
				),
			),
			want: true,
		},
		{
			name: "large asymmetric balanced tree by one deep value",
			root: node(8,
				node(6,
					node(5, node(4, nil, nil), node(7, nil, nil)),
					node(9, nil, nil),
				),
				node(6,
					node(9, nil, nil),
					node(5, node(7, nil, nil), node(10, nil, nil)),
				),
			),
			want: false,
		},
		{
			name: "zigzag symmetric",
			root: node(1,
				node(2, nil, node(3, node(4, nil, nil), nil)),
				node(2, node(3, nil, node(4, nil, nil)), nil),
			),
			want: true,
		},
		{
			name: "zigzag asymmetric",
			root: node(1,
				node(2, nil, node(3, node(4, nil, nil), nil)),
				node(2, node(3, node(4, nil, nil), nil), nil),
			),
			want: false,
		},
		{
			name: "zero values symmetric",
			root: node(0,
				node(0, node(0, nil, nil), nil),
				node(0, nil, node(0, nil, nil)),
			),
			want: true,
		},
		{
			name: "zero values asymmetric",
			root: node(0,
				node(0, node(0, nil, nil), nil),
				node(0, node(0, nil, nil), nil),
			),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSymmetric(tt.root)
			if got != tt.want {
				t.Fatalf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
