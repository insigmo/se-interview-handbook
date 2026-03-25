package binary_search

import (
	"testing"
)

// 4567012
func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	minValue := nums[0]

	for left <= right {
		mid := (left + right) / 2
		midValue := nums[mid]

		if minValue > midValue {
			minValue = midValue
		}
		if minValue > nums[left] {
			minValue = nums[left]
		}
		if nums[left] <= midValue {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return minValue
}

func TestFindMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			name: "example 2",
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			name: "example 3 no effective rotation",
			nums: []int{11, 13, 15, 17},
			want: 11,
		},
		{
			name: "single element",
			nums: []int{42},
			want: 42,
		},
		{
			name: "two elements rotated",
			nums: []int{2, 1},
			want: 1,
		},
		{
			name: "two elements not changed",
			nums: []int{1, 2},
			want: 1,
		},
		{
			name: "minimum in middle",
			nums: []int{5, 6, 7, 1, 2, 3, 4},
			want: 1,
		},
		{
			name: "minimum near end",
			nums: []int{2, 3, 4, 5, 6, 7, 1},
			want: 1,
		},
		{
			name: "minimum near beginning",
			nums: []int{7, 1, 2, 3, 4, 5, 6},
			want: 1,
		},
		{
			name: "all negative values",
			nums: []int{-3, -2, -1, -10, -9, -8, -7},
			want: -10,
		},
		{
			name: "mixed negative and positive",
			nums: []int{0, 2, 4, -5, -3, -1},
			want: -5,
		},
		{
			name: "minimum is zero",
			nums: []int{1, 2, 3, 4, 5, 0},
			want: 0,
		},
		{
			name: "already sorted larger",
			nums: []int{-10, -3, 0, 5, 9, 12, 20},
			want: -10,
		},
		{
			name: "rotation by one",
			nums: []int{9, 1, 2, 3, 4, 5, 6, 7, 8},
			want: 1,
		},
		{
			name: "rotation by len minus one",
			nums: []int{2, 3, 4, 5, 6, 7, 8, 9, 1},
			want: 1,
		},
		{
			name: "extreme constraint values",
			nums: []int{0, 100, 200, 5000, -5000, -100, -1},
			want: -5000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMin(tt.nums)
			if got != tt.want {
				t.Fatalf("findMin(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
