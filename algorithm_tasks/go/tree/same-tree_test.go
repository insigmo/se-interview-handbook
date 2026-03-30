// https://leetcode.com/problems/same-tree/description/
// explanation: https://www.youtube.com/watch?v=jMBzan7hAPI

package tree

import "testing"

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func TestIsSameTree(t *testing.T) {
	node := func(val int, left, right *TreeNode) *TreeNode {
		return &TreeNode{Val: val, Left: left, Right: right}
	}

	tests := []struct {
		name string
		p    *TreeNode
		q    *TreeNode
		want bool
	}{
		{
			name: "both nil",
			p:    nil,
			q:    nil,
			want: true,
		},
		{
			name: "one nil one non nil",
			p:    node(1, nil, nil),
			q:    nil,
			want: false,
		},
		{
			name: "single same node",
			p:    node(1, nil, nil),
			q:    node(1, nil, nil),
			want: true,
		},
		{
			name: "single different value",
			p:    node(1, nil, nil),
			q:    node(2, nil, nil),
			want: false,
		},
		{
			name: "example one identical",
			p:    node(1, node(2, nil, nil), node(3, nil, nil)),
			q:    node(1, node(2, nil, nil), node(3, nil, nil)),
			want: true,
		},
		{
			name: "example two different structure",
			p:    node(1, node(2, nil, nil), nil),
			q:    node(1, nil, node(2, nil, nil)),
			want: false,
		},
		{
			name: "example three different values",
			p:    node(1, node(2, nil, nil), node(1, nil, nil)),
			q:    node(1, node(1, nil, nil), node(2, nil, nil)),
			want: false,
		},
		{
			name: "same values different deeper structure",
			p:    node(1, node(2, node(3, nil, nil), nil), nil),
			q:    node(1, node(2, nil, node(3, nil, nil)), nil),
			want: false,
		},
		{
			name: "identical left skewed",
			p:    node(1, node(2, node(3, node(4, nil, nil), nil), nil), nil),
			q:    node(1, node(2, node(3, node(4, nil, nil), nil), nil), nil),
			want: true,
		},
		{
			name: "identical right skewed",
			p:    node(1, nil, node(2, nil, node(3, nil, node(4, nil, nil)))),
			q:    node(1, nil, node(2, nil, node(3, nil, node(4, nil, nil)))),
			want: true,
		},
		{
			name: "right skewed different tail value",
			p:    node(1, nil, node(2, nil, node(3, nil, nil))),
			q:    node(1, nil, node(2, nil, node(4, nil, nil))),
			want: false,
		},
		{
			name: "negative values identical",
			p:    node(-1, node(-2, nil, nil), node(-3, nil, nil)),
			q:    node(-1, node(-2, nil, nil), node(-3, nil, nil)),
			want: true,
		},
		{
			name: "negative values different",
			p:    node(-1, node(-2, nil, nil), node(-3, nil, nil)),
			q:    node(-1, node(-2, nil, nil), node(-4, nil, nil)),
			want: false,
		},
		{
			name: "duplicate values identical",
			p:    node(1, node(1, node(1, nil, nil), node(1, nil, nil)), node(1, nil, nil)),
			q:    node(1, node(1, node(1, nil, nil), node(1, nil, nil)), node(1, nil, nil)),
			want: true,
		},
		{
			name: "duplicate values different structure",
			p:    node(1, node(1, node(1, nil, nil), nil), node(1, nil, nil)),
			q:    node(1, node(1, nil, node(1, nil, nil)), node(1, nil, nil)),
			want: false,
		},
		{
			name: "one missing subtree",
			p:    node(5, node(3, nil, nil), node(7, nil, nil)),
			q:    node(5, node(3, nil, nil), nil),
			want: false,
		},
		{
			name: "larger identical balanced tree",
			p: node(4,
				node(2, node(1, nil, nil), node(3, nil, nil)),
				node(6, node(5, nil, nil), node(7, nil, nil)),
			),
			q: node(4,
				node(2, node(1, nil, nil), node(3, nil, nil)),
				node(6, node(5, nil, nil), node(7, nil, nil)),
			),
			want: true,
		},
		{
			name: "larger balanced tree one deep mismatch",
			p: node(4,
				node(2, node(1, nil, nil), node(3, nil, nil)),
				node(6, node(5, nil, nil), node(7, nil, nil)),
			),
			q: node(4,
				node(2, node(1, nil, nil), node(3, nil, nil)),
				node(6, node(5, nil, nil), node(8, nil, nil)),
			),
			want: false,
		},
		{
			name: "same root different child side",
			p:    node(1, node(2, nil, nil), nil),
			q:    node(1, nil, node(2, nil, nil)),
			want: false,
		},
		{
			name: "zero values identical",
			p:    node(0, node(0, nil, nil), node(0, nil, nil)),
			q:    node(0, node(0, nil, nil), node(0, nil, nil)),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSameTree(tt.p, tt.q)
			if got != tt.want {
				t.Fatalf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
