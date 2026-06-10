// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

package array

import "testing"

func removeDuplicates(nums []int) int {
	left, length := 1, len(nums)

	if length == 0 {
		return 0
	}

	for i := 1; i < length; i++ {
		if nums[i] != nums[i-1] {
			nums[left] = nums[i]
			left++
		}
	}
	return left
}

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"empty", []int{}, 0},
		{"one", []int{1}, 1},
		{"all unique", []int{1, 2, 3}, 3},
		{"all duplicates", []int{1, 1, 1}, 1},
		{"simple example", []int{1, 1, 2}, 2},
		{"mixed", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
		{"negatives", []int{-3, -3, -2, -1, -1, 0, 0, 0, 5}, 5},
		{"large block", []int{1, 2, 2, 2, 3, 4, 4, 5}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.nums)
			if got != tt.want {
				t.Fatalf("removeDuplicates(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
