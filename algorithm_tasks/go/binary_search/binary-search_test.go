// https://leetcode.com/problems/binary-search/description/

package binary_search

import "testing"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := 0
	for left <= right {
		mid = (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "example 1 found",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   4,
		},
		{
			name:   "example 2 not found",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			want:   -1,
		},
		{
			name:   "single element found",
			nums:   []int{5},
			target: 5,
			want:   0,
		},
		{
			name:   "single element not found",
			nums:   []int{5},
			target: 1,
			want:   -1,
		},
		{
			name: "target is first element",

			nums:   []int{1, 3, 5, 7, 9},
			target: 1,
			want:   0,
		},
		{
			name:   "target is last element",
			nums:   []int{1, 3, 5, 7, 9},
			target: 9,
			want:   4,
		},
		{
			name:   "target in middle",
			nums:   []int{1, 3, 5, 7, 9},
			target: 5,
			want:   2,
		},
		{
			name:   "target less than min",
			nums:   []int{1, 3, 5, 7, 9},
			target: 0,
			want:   -1,
		},
		{
			name:   "target greater than max",
			nums:   []int{1, 3, 5, 7, 9},
			target: 10,
			want:   -1,
		},
		{
			name:   "negative numbers",
			nums:   []int{-10, -5, -2, 0, 4, 8},
			target: -5,
			want:   1,
		},
		{
			name:   "negative target not found",
			nums:   []int{-10, -5, -2, 0, 4, 8},
			target: -3,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.nums, tt.target)
			if got != tt.want {
				t.Fatalf("search(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
