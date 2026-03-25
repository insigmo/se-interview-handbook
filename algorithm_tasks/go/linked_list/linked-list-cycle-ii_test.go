// https://leetcode.com/problems/linked-list-cycle-ii/description/

package linked_list

import (
	"strconv"
	"testing"
)

func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	flag := false

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			flag = true
			break
		}
	}
	if !flag {
		return nil
	}

	pointer := head
	for pointer != fast {
		pointer = pointer.Next
		fast = fast.Next
	}

	return pointer
}

func TestDetectCycle(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		pos  int
	}{
		{
			name: "example 1",
			vals: []int{3, 2, 0, -4},
			pos:  1,
		},
		{
			name: "example 2",
			vals: []int{1, 2},
			pos:  0,
		},
		{
			name: "example 3 no cycle",
			vals: []int{1},
			pos:  -1,
		},
		{
			name: "empty list",
			vals: []int{},
			pos:  -1,
		},
		{
			name: "single node self cycle",
			vals: []int{1},
			pos:  0,
		},
		{
			name: "cycle starts at head",
			vals: []int{1, 2, 3, 4},
			pos:  0,
		},
		{
			name: "cycle starts in middle",
			vals: []int{1, 2, 3, 4, 5, 6},
			pos:  2,
		},

		{
			name: "cycle starts in middle",
			vals: []int{1, 2, 3, 4, 5, 6, 7},
			pos:  0,
		},
		{
			name: "two nodes no cycle",
			vals: []int{1, 2},
			pos:  -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head, want := BuildListWithCycle(tt.vals, tt.pos)

			got := detectCycle(head)

			if got != want {
				gotVal := "<nil>"
				wantVal := "<nil>"

				if got != nil {
					gotVal = stringValue(got.Val)
				}
				if want != nil {
					wantVal = stringValue(want.Val)
				}

				t.Fatalf("detectCycle() = %v (val=%s), want %v (val=%s)", got, gotVal, want, wantVal)
			}
		})
	}
}

func stringValue(v int) string {
	return strconv.Itoa(v)
}
