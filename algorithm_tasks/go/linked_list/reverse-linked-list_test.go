package linked_list

import (
	"slices"
	"testing"
)

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for prev != current {
		if current == nil {
			break
		}

		nextNode := current.Next
		current.Next = prev

		prev = current
		current = nextNode
	}
	return prev
}

func TestReverseLinkedList(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "пример 1: [1,2,3,4,5]",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "пример 2: [1,2]",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "пустой список",
			input:    []int{},
			expected: nil,
		},
		{
			name:     "один элемент",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "два одинаковых элемента",
			input:    []int{7, 7},
			expected: []int{7, 7},
		},
		{
			name:     "отрицательные значения",
			input:    []int{-5000, -1, 0, 1, 5000},
			expected: []int{5000, 1, 0, -1, -5000},
		},
		{
			name:     "уже обратный порядок",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "палиндром (должен остаться таким же по значениям)",
			input:    []int{1, 2, 3, 2, 1},
			expected: []int{1, 2, 3, 2, 1},
		},
		{
			name:     "граничный максимум: значения -5000 и 5000",
			input:    []int{-5000, 5000},
			expected: []int{5000, -5000},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			head := ConvertSliceToListNode(tc.input)
			got := ConvertListNodeToSlice(reverseList(head))
			if !slices.Equal(got, tc.expected) {
				t.Errorf("\n\tgot:\t\t%v\n\texpected:\t%v", got, tc.expected)
			}
		})
	}
}
