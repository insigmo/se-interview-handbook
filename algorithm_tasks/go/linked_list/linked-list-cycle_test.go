// https://leetcode.com/problems/linked-list-cycle/

package linked_list

import (
	"testing"
)

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}

	}

	return false
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		mutate func(head *ListNode)
		want   bool
	}{
		{
			name:  "nil list",
			input: nil,
			want:  false,
		},
		{
			name:  "one node no cycle",
			input: []int{1},
			want:  false,
		},
		{
			name:  "one node self cycle",
			input: []int{1},
			mutate: func(head *ListNode) {
				head.Next = head
			},
			want: true,
		},
		{
			name:  "two nodes no cycle",
			input: []int{1, 2},
			want:  false,
		},
		{
			name:  "two nodes cycle",
			input: []int{1, 2},
			mutate: func(head *ListNode) {
				head.Next.Next = head
			},
			want: true,
		},
		{
			name:  "long list no cycle",
			input: []int{1, 1, 2, 3, 4, 4, 5, 6},
			want:  false,
		},
		{
			name:  "cycle to middle node",
			input: []int{1, 1, 2, 3, 4, 4, 5, 6},
			mutate: func(head *ListNode) {
				head.Next.Next.Next = head.Next
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head *ListNode
			if tt.input != nil {
				head = ConvertSliceToListNode(tt.input)
			}

			if tt.mutate != nil {
				tt.mutate(head)
			}

			got := hasCycle(head)
			if got != tt.want {
				t.Fatalf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
