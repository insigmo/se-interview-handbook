// https://leetcode.com/problems/search-in-rotated-sorted-array-ii/description/

package binary_search

import "testing"

// rename func to "search"
// 1, 0, 1, 1, 1
func searchInRotatedSortedArrayII(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	for left <= right {

		mid := (left + right) / 2
		midValue := nums[mid]

		if target == midValue {
			return true
		}

		for left < mid && nums[left] == midValue {
			if left < len(nums) {
				left++
			}
		}

		for mid < right && nums[right] == midValue {
			if right > 0 {
				right--
			}
		}
		if nums[left] <= midValue {
			if nums[left] <= target && target < midValue {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if midValue < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}

// LeetCode 81: rotated sorted array search with duplicates. [web:38]
func TestSearchInRotatedSortedArrayII(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{
			name:   "example 1 target present",
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 0,
			want:   true,
		},
		{
			name:   "example 2 target absent",
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 3,
			want:   false,
		},
		{
			name:   "single element present",
			nums:   []int{1},
			target: 1,
			want:   true,
		},
		{
			name:   "single element absent",
			nums:   []int{1},
			target: 0,
			want:   false,
		},
		{
			name:   "two elements rotated target first",
			nums:   []int{3, 1},
			target: 3,
			want:   true,
		},
		{
			name:   "two elements rotated target second",
			nums:   []int{3, 1},
			target: 1,
			want:   true,
		},
		{
			name:   "two equal elements target present",
			nums:   []int{1, 1},
			target: 1,
			want:   true,
		},
		{
			name:   "two equal elements target absent",
			nums:   []int{1, 1},
			target: 2,
			want:   false,
		},
		{
			name:   "not effectively rotated with duplicates",
			nums:   []int{1, 1, 2, 2, 3, 3, 4, 4},
			target: 3,
			want:   true,
		},
		{
			name:   "not effectively rotated target absent",
			nums:   []int{1, 1, 2, 2, 3, 3, 4, 4},
			target: 5,
			want:   false,
		},
		{
			name:   "all elements same target present",
			nums:   []int{7, 7, 7, 7, 7},
			target: 7,
			want:   true,
		},
		{
			name:   "all elements same target absent",
			nums:   []int{7, 7, 7, 7, 7},
			target: 3,
			want:   false,
		},
		{
			name:   "duplicates around pivot target is pivot value",
			nums:   []int{2, 2, 2, 3, 4, 2},
			target: 2,
			want:   true,
		},
		{
			name:   "duplicates around pivot target in unique segment",
			nums:   []int{2, 2, 2, 3, 4, 2},
			target: 3,
			want:   true,
		},
		{
			name:   "duplicates around pivot target absent",
			nums:   []int{2, 2, 2, 3, 4, 2},
			target: 5,
			want:   false,
		},
		{
			name:   "ambiguous halves due to duplicates target present",
			nums:   []int{1, 0, 1, 1, 1},
			target: 0,
			want:   true,
		},
		{
			name:   "ambiguous halves due to duplicates target absent",
			nums:   []int{1, 0, 1, 1, 1},
			target: 2,
			want:   false,
		},
		{
			name:   "classic tricky case target present",
			nums:   []int{1, 1, 3, 1},
			target: 3,
			want:   true,
		},
		{
			name:   "classic tricky case target absent",
			nums:   []int{1, 1, 3, 1},
			target: 2,
			want:   false,
		},
		{
			name:   "pivot near end target present",
			nums:   []int{2, 2, 3, 4, 5, 1, 2},
			target: 1,
			want:   true,
		},
		{
			name:   "pivot near beginning target present",
			nums:   []int{5, 1, 2, 2, 3, 4, 5},
			target: 4,
			want:   true,
		},
		{
			name:   "target appears many times",
			nums:   []int{4, 4, 5, 6, 6, 7, 0, 1, 2, 4, 4},
			target: 4,
			want:   true,
		},
		{
			name:   "target absent among many duplicates",
			nums:   []int{4, 4, 5, 6, 6, 7, 0, 1, 2, 4, 4},
			target: 3,
			want:   false,
		},
		{
			name:   "negative values target present",
			nums:   []int{-1, 0, 0, 1, 2, -4, -3, -2, -2},
			target: -3,
			want:   true,
		},
		{
			name:   "negative values target absent",
			nums:   []int{-1, 0, 0, 1, 2, -4, -3, -2, -2},
			target: 3,
			want:   false,
		},
		{
			name:   "min and max constraint style values present low",
			nums:   []int{0, 1, 9999, 10000, -10000, -10000, -5},
			target: -10000,
			want:   true,
		},
		{
			name:   "min and max constraint style values present high",
			nums:   []int{0, 1, 9999, 10000, -10000, -10000, -5},
			target: 10000,
			want:   true,
		},
		{
			name:   "min and max constraint style values absent",
			nums:   []int{0, 1, 9999, 10000, -10000, -10000, -5},
			target: 42,
			want:   false,
		},
		{
			name:   "all duplicates except one target unique",
			nums:   []int{2, 2, 2, 2, 3, 2, 2},
			target: 3,
			want:   true,
		},
		{
			name:   "all duplicates except one target duplicate value",
			nums:   []int{2, 2, 2, 2, 3, 2, 2},
			target: 2,
			want:   true,
		},
		{
			name:   "all duplicates except one target absent",
			nums:   []int{2, 2, 2, 2, 3, 2, 2},
			target: 1,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchInRotatedSortedArrayII(tt.nums, tt.target)
			if got != tt.want {
				t.Fatalf("search(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
