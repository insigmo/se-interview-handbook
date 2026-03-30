package tree

import "testing"

func pathSum(root *TreeNode, targetSum int) (answer [][]int) {
	// abs by bit wise for getting max slice capacity
	mask := targetSum >> 1
	length := (targetSum ^ mask) - mask

	answer = make([][]int, 0, length)

	if root == nil {
		return answer
	}
	dfsPathSumII(root, targetSum, []int{}, &answer)
	return
}

func dfsPathSumII(node *TreeNode, remain int, nums []int, answer *[][]int) {
	// Если узел пустой, дальше идти некуда.
	if node == nil {
		return
	}

	// Добавляем текущее значение в слайс чисел.
	nums = append(nums, node.Val)

	// Проверяем что права и левая ноды, это не листья, то есть не конец дерева
	// Также проверяем что последнее значение равно значению что осталось от target - node.Val
	// Если получаем true, то тогда добавляем копию в основной результат, а не тот же самый слайс,
	// Если будем добавлять тот же, то слайсы передадутся по ссылке и данные менять постоянно
	if node.Left == nil && node.Right == nil && node.Val == remain {
		*answer = append(*answer, append(make([]int, 0, len(nums)+1), nums...))
		return
	}

	// Если мы не нашли результат, то проверяем отдельно не лист ли правый или левый элемент
	// После чего, отнимаем от таргета текущее значение. remain назвал потому что это не совсем таргет,
	// а остаток от разности
	if node.Left != nil {
		dfsPathSumII(node.Left, remain-node.Val, nums, answer)
	}
	if node.Right != nil {
		dfsPathSumII(node.Right, remain-node.Val, nums, answer)
	}
}

func TestPathSum(t *testing.T) {
	node := func(val int, left, right *TreeNode) *TreeNode {
		return &TreeNode{Val: val, Left: left, Right: right}
	}

	equalPath := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	assertPaths := func(t *testing.T, got, want [][]int, target int) {
		t.Helper()

		for _, path := range got {
			sum := 0
			for _, v := range path {
				sum += v
			}
			if sum != target {
				t.Fatalf("path %v has sum %d, want %d", path, sum, target)
			}
		}

		if len(got) != len(want) {
			t.Fatalf("got %v, want %v, target %d", got, want, target)
		}

		used := make([]bool, len(got))
		for _, wp := range want {
			found := false
			for i, gp := range got {
				if used[i] {
					continue
				}
				if equalPath(gp, wp) {
					used[i] = true
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("missing path %v in %v", wp, got)
			}
		}
	}

	tests := []struct {
		name      string
		root      *TreeNode
		targetSum int
		want      [][]int
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
					node(4, node(5, nil, nil), node(1, nil, nil)),
				),
			),
			targetSum: 22,
			want:      [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
		{
			name: "example two",
			root: node(1,
				node(2, nil, nil),
				node(3, nil, nil),
			),
			targetSum: 5,
			want:      [][]int{},
		},
		{
			name: "example three",
			root: node(1,
				node(2, nil, nil),
				nil,
			),
			targetSum: 0,
			want:      [][]int{},
		},
		{
			name:      "empty tree",
			root:      nil,
			targetSum: 0,
			want:      [][]int{},
		},
		{
			name:      "single node matches",
			root:      node(7, nil, nil),
			targetSum: 7,
			want:      [][]int{{7}},
		},
		{
			name:      "single node does not match",
			root:      node(7, nil, nil),
			targetSum: 8,
			want:      [][]int{},
		},
		{
			name:      "single negative node matches",
			root:      node(-7, nil, nil),
			targetSum: -7,
			want:      [][]int{{-7}},
		},
		{
			name: "prefix sum on non leaf must not count",
			root: node(1,
				node(2,
					node(3, nil, nil),
					nil,
				),
				nil,
			),
			targetSum: 3,
			want:      [][]int{},
		},
		{
			name: "left skewed valid",
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
			want:      [][]int{{1, 2, 3, 4}},
		},
		{
			name: "right skewed valid",
			root: node(1,
				nil,
				node(2,
					nil,
					node(3,
						nil,
						node(4, nil, nil),
					),
				),
			),
			targetSum: 10,
			want:      [][]int{{1, 2, 3, 4}},
		},
		{
			name: "multiple valid paths with same values",
			root: node(1,
				node(2, node(3, nil, nil), nil),
				node(2, nil, node(3, nil, nil)),
			),
			targetSum: 6,
			want:      [][]int{{1, 2, 3}, {1, 2, 3}},
		},
		{
			name: "zero values duplicate paths",
			root: node(0,
				node(0, nil, nil),
				node(0, nil, nil),
			),
			targetSum: 0,
			want:      [][]int{{0, 0}, {0, 0}},
		},
		{
			name: "mixed positive and negative",
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
			want:      [][]int{{1, -2, 1, -1}},
		},
		{
			name: "negative path on right",
			root: node(-2,
				nil,
				node(-3, nil, nil),
			),
			targetSum: -5,
			want:      [][]int{{-2, -3}},
		},
		{
			name: "two valid paths different lengths",
			root: node(10,
				node(5,
					node(3, node(3, nil, nil), node(-2, nil, nil)),
					node(2, nil, node(1, nil, nil)),
				),
				node(-3,
					nil,
					node(11, nil, nil),
				),
			),
			targetSum: 18,
			want:      [][]int{{10, 5, 2, 1}, {10, -3, 11}},
		},
		{
			name: "single valid path among close sums",
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
			targetSum: 26,
			want:      [][]int{{5, 8, 13}},
		},
		{
			name: "leaf requirement blocks internal match",
			root: node(5,
				node(4,
					node(8, nil, nil),
					nil,
				),
				nil,
			),
			targetSum: 9,
			want:      [][]int{},
		},
		{
			name: "balanced tree one valid path",
			root: node(3,
				node(9, nil, nil),
				node(20, node(15, nil, nil), node(7, nil, nil)),
			),
			targetSum: 38,
			want:      [][]int{{3, 20, 15}},
		},
		{
			name: "duplicate values many valid branches",
			root: node(1,
				node(1, node(1, nil, nil), node(1, nil, nil)),
				node(1, node(1, nil, nil), node(1, nil, nil)),
			),
			targetSum: 3,
			want:      [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		},
		{
			name: "target impossible",
			root: node(2,
				node(3, node(4, nil, nil), node(5, nil, nil)),
				node(6, node(7, nil, nil), node(8, nil, nil)),
			),
			targetSum: 100,
			want:      [][]int{},
		},
		{
			name: "path with zeros and negatives",
			root: node(0,
				node(-1, node(1, nil, nil), nil),
				node(1, node(-1, nil, nil), nil),
			),
			targetSum: 0,
			want:      [][]int{{0, -1, 1}, {0, 1, -1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pathSum(tt.root, tt.targetSum)
			assertPaths(t, got, tt.want, tt.targetSum)
		})
	}
}
