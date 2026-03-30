// https://leetcode.com/problems/path-sum
//
// Объяснение.
// Я иду по дереву сверху вниз и на каждом узле прибавляю его значение к текущей сумме.
// Когда дохожу до конца пути, то есть до листа, я просто смотрю: получилась нужная сумма или нет.
// Если в одной ветке не получилось, я пробую другую.
// Если хотя бы в одной ветке сумма совпала, значит ответ true
//
// explanation:
// I go through the tree from top to bottom, and at each node I add its value to the current sum.
// When I reach the end of a path, meaning a leaf node, I simply check whether the sum matches the target or not.
// If one branch doesn’t work, I try another one.
// If the sum matches in at least one branch, then the answer is true.

package tree

import "testing"

func dfsPathSum(node *TreeNode, currentSum int, target int) bool {
	if node == nil {
		return false
	}

	currentSum += node.Val
	if node.Left == nil && node.Right == nil {
		return currentSum == target
	}

	return dfsPathSum(node.Left, currentSum, target) || dfsPathSum(node.Right, currentSum, target)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfsPathSum(root, 0, targetSum)
}

func TestHasPathSum(t *testing.T) {
	node := func(val int, left, right *TreeNode) *TreeNode {
		return &TreeNode{Val: val, Left: left, Right: right}
	}

	tests := []struct {
		name      string
		root      *TreeNode
		targetSum int
		want      bool
	}{
		{
			name: "example one",
			root: node(5,
				node(4,
					node(11, node(7, nil, nil), node(2, nil, nil)),
					nil,
				),
				node(8,
					node(13, nil, nil),
					node(4, nil, node(1, nil, nil)),
				),
			),
			targetSum: 22,
			want:      true,
		},
		{
			name: "example two",
			root: node(1,
				node(2, nil, nil),
				node(3, nil, nil),
			),
			targetSum: 5,
			want:      false,
		},
		{
			name:      "example three empty tree",
			root:      nil,
			targetSum: 0,
			want:      false,
		},
		{
			name:      "single node matches",
			root:      node(1, nil, nil),
			targetSum: 1,
			want:      true,
		},
		{
			name:      "single node does not match",
			root:      node(1, nil, nil),
			targetSum: 2,
			want:      false,
		},
		{
			name: "root to non leaf prefix must not count",
			root: node(1,
				node(2,
					node(3, nil, nil),
					nil,
				),
				nil,
			),
			targetSum: 3,
			want:      false,
		},
		{
			name: "left skewed valid path",
			root: node(1,
				node(2,
					node(3,
						node(4, nil, nil),
						nil,
					),
					nil,
				),
				nil,
			),
			targetSum: 10,
			want:      true,
		},
		{
			name: "left skewed invalid path",
			root: node(1,
				node(2,
					node(3,
						node(4, nil, nil),
						nil,
					),
					nil,
				),
				nil,
			),
			targetSum: 9,
			want:      false,
		},
		{
			name: "right skewed valid path",
			root: node(5,
				nil,
				node(4,
					nil,
					node(3,
						nil,
						node(2, nil, nil),
					),
				),
			),
			targetSum: 14,
			want:      true,
		},
		{
			name: "negative values valid path",
			root: node(-2,
				nil,
				node(-3, nil, nil),
			),
			targetSum: -5,
			want:      true,
		},
		{
			name: "negative values no valid path",
			root: node(-2,
				node(-3, nil, nil),
				node(-4, nil, nil),
			),
			targetSum: -6,
			want:      true,
		},
		{
			name: "mixed positive and negative valid",
			root: node(1,
				node(-2,
					node(1, node(-1, nil, nil), nil),
					node(3, nil, nil),
				),
				node(-3,
					node(-2, nil, nil),
					nil,
				),
			),
			targetSum: -1,
			want:      true,
		},
		{
			name: "mixed positive and negative invalid",
			root: node(1,
				node(-2,
					node(1, node(-1, nil, nil), nil),
					node(3, nil, nil),
				),
				node(-3,
					node(-2, nil, nil),
					nil,
				),
			),
			targetSum: 100,
			want:      false,
		},
		{
			name: "zero target with zero path",
			root: node(0,
				node(1, nil, nil),
				node(0, nil, nil),
			),
			targetSum: 0,
			want:      true,
		},
		{
			name: "zero target but only internal zero prefix",
			root: node(0,
				node(1, nil, nil),
				nil,
			),
			targetSum: 0,
			want:      false,
		},
		{
			name: "duplicate values valid",
			root: node(1,
				node(1,
					node(1, nil, nil),
					node(1, nil, nil),
				),
				node(1, nil, nil),
			),
			targetSum: 3,
			want:      true,
		},
		{
			name: "duplicate values invalid",
			root: node(1,
				node(1,
					node(1, nil, nil),
					node(1, nil, nil),
				),
				node(1, nil, nil),
			),
			targetSum: 5,
			want:      false,
		},
		{
			name: "leaf on one side only",
			root: node(7,
				node(3, nil, nil),
				node(4, node(1, nil, nil), nil),
			),
			targetSum: 11,
			want:      false,
		},
		{
			name: "target achieved on deeper right leaf",
			root: node(7,
				node(3, nil, nil),
				node(4, node(1, nil, nil), node(2, nil, nil)),
			),
			targetSum: 13,
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasPathSum(tt.root, tt.targetSum)
			if got != tt.want {
				t.Fatalf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
