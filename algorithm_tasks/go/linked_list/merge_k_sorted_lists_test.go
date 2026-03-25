// https://leetcode.com/problems/merge-k-sorted-lists/

package linked_list

import (
	"slices"
	"testing"
)

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		var mergeList []*ListNode
		listLength := len(lists)
		for i := 0; i < listLength; i += 2 {
			l1 := lists[i]
			var l2 *ListNode

			if i+1 < listLength {
				l2 = lists[i+1]
			}
			mergeList = append(mergeList, mergeTwoLists(l1, l2))
		}
		lists = mergeList
	}

	return lists[0]
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	tail := dummy

	for list1 != nil && list2 != nil {
		if list1.Val == 0 || list2.Val == 0 {
			break
		}
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}
	return dummy.Next
}

func TestMergeKLists(t *testing.T) {
	buildLists := func(data [][]int) []*ListNode {
		if data == nil {
			return nil
		}

		lists := make([]*ListNode, 0, len(data))
		for _, arr := range data {
			lists = append(lists, ConvertSliceToListNode(arr))
		}
		return lists
	}

	tests := []struct {
		name  string
		input [][]int
		build func() []*ListNode
		want  []int
	}{
		{
			name:  "nil outer slice",
			input: nil,
			want:  nil,
		},
		{
			name:  "empty outer slice",
			input: [][]int{},
			want:  nil,
		},
		{
			name: "single nil list",
			build: func() []*ListNode {
				return []*ListNode{nil}
			},
			want: nil,
		},
		{
			name:  "single list",
			input: [][]int{{1, 2, 3}},
			want:  []int{1, 2, 3},
		},
		{
			name: "nil lists mixed with values",
			build: func() []*ListNode {
				return []*ListNode{
					nil,
					ConvertSliceToListNode([]int{1, 3, 5}),
					nil,
					ConvertSliceToListNode([]int{2, 4}),
				}
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "leetcode example",
			input: [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			want:  []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:  "duplicates",
			input: [][]int{{1, 1, 1}, {1, 1}, {1}},
			want:  []int{1, 1, 1, 1, 1, 1},
		},
		{
			name: "all nil lists",
			build: func() []*ListNode {
				return []*ListNode{nil, nil, nil}
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			var lists []*ListNode
			if tt.build != nil {
				lists = tt.build()
			} else {
				lists = buildLists(tt.input)
			}

			got := ConvertListNodeToSlice(mergeKLists(lists))
			if !slices.Equal(got, tt.want) {
				t.Fatalf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
