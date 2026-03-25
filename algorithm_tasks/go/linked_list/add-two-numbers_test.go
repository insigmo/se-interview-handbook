//

package linked_list

import (
	"slices"
	"testing"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	tail := dummy
	extra := 0

	for l1 != nil || l2 != nil {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		total := a + b + extra
		extra = total / 10

		tail.Next = &ListNode{Val: total % 10}
		tail = tail.Next
	}

	if extra == 1 {
		tail.Next = &ListNode{Val: extra}
	}

	return dummy.Next
}

func TestAddTwoNumbers(t *testing.T) {
	toList := func(nums []int) *ListNode {
		if nums == nil {
			return nil
		}
		return ConvertSliceToListNode(nums)
	}

	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		{
			name: "example",
			l1:   []int{1, 2, 3},
			l2:   []int{7, 9, 3},
			want: []int{8, 1, 7},
		},
		{
			name: "both nil",
			l1:   nil,
			l2:   nil,
			want: nil,
		},
		{
			name: "first nil",
			l1:   nil,
			l2:   []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "second nil",
			l1:   []int{4, 5, 6},
			l2:   nil,
			want: []int{4, 5, 6},
		},
		{
			name: "single digit no carry",
			l1:   []int{2},
			l2:   []int{3},
			want: []int{5},
		},
		{
			name: "single digit with carry",
			l1:   []int{9},
			l2:   []int{7},
			want: []int{6, 1},
		},
		{
			name: "different lengths",
			l1:   []int{9, 9},
			l2:   []int{1},
			want: []int{0, 0, 1},
		},
		{
			name: "carry through all digits",
			l1:   []int{9, 9, 9},
			l2:   []int{1},
			want: []int{0, 0, 0, 1},
		},
		{
			name: "zeros",
			l1:   []int{0},
			l2:   []int{0},
			want: []int{0},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			var (
				got      []int
				panicked any
			)

			func() {
				defer func() {
					panicked = recover()
				}()

				n1 := toList(tt.l1)
				n2 := toList(tt.l2)
				got = ConvertListNodeToSlice(addTwoNumbers(n1, n2))
			}()

			if panicked != nil {
				t.Fatalf("addTwoNumbers() panicked: %v", panicked)
			}

			if !slices.Equal(got, tt.want) {
				t.Fatalf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
