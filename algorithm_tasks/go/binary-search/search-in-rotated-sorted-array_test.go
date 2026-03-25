// https://leetcode.com/problems/search-in-rotated-sorted-array/

package binary_search

import "testing"

// rename func to "search"
func searchInRotatedArray(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		midValue := nums[mid]

		if target == midValue {
			return mid
		}
		if nums[left] <= midValue {
			if target > midValue || target < nums[left] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if target < midValue || target > nums[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		// Примеры из условия
		{
			name:   "example_found",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			name:   "example_not_found",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			name:   "single_not_found",
			nums:   []int{1},
			target: 0,
			want:   -1,
		},

		// Один элемент
		{
			name:   "single_found",
			nums:   []int{1},
			target: 1,
			want:   0,
		},

		// Без ротации
		{
			name:   "not_rotated_found_first",
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			target: 1,
			want:   0,
		},
		{
			name:   "not_rotated_found_last",
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			target: 7,
			want:   6,
		},
		{
			name:   "not_rotated_found_middle",
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			target: 4,
			want:   3,
		},
		{
			name:   "not_rotated_missing",
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			target: 8,
			want:   -1,
		},

		// Ротация на 1
		{
			name:   "rotated_by_one_found_pivot",
			nums:   []int{7, 1, 2, 3, 4, 5, 6},
			target: 1,
			want:   1,
		},
		{
			name:   "rotated_by_one_found_first",
			nums:   []int{7, 1, 2, 3, 4, 5, 6},
			target: 7,
			want:   0,
		},

		// Ротация почти на полный размер
		{
			name:   "rotated_near_full_found_last",
			nums:   []int{2, 3, 4, 5, 6, 7, 1},
			target: 1,
			want:   6,
		},
		{
			name:   "rotated_near_full_found_middle",
			nums:   []int{2, 3, 4, 5, 6, 7, 1},
			target: 5,
			want:   3,
		},

		// Два элемента
		{
			name:   "two_elements_rotated_found_first",
			nums:   []int{2, 1},
			target: 2,
			want:   0,
		},
		{
			name:   "two_elements_rotated_found_second",
			nums:   []int{2, 1},
			target: 1,
			want:   1,
		},
		{
			name:   "two_elements_rotated_missing",
			nums:   []int{2, 1},
			target: 3,
			want:   -1,
		},

		// Точка поворота внутри массива
		{
			name:   "find_pivot_element",
			nums:   []int{6, 7, 8, 1, 2, 3, 4, 5},
			target: 1,
			want:   3,
		},
		{
			name:   "find_left_sorted_half",
			nums:   []int{6, 7, 8, 1, 2, 3, 4, 5},
			target: 7,
			want:   1,
		},
		{
			name:   "find_right_sorted_half",
			nums:   []int{6, 7, 8, 1, 2, 3, 4, 5},
			target: 4,
			want:   6,
		},
		{
			name:   "missing_between_values",
			nums:   []int{6, 7, 8, 1, 2, 3, 4, 5},
			target: 9,
			want:   -1,
		},

		// Отрицательные значения
		{
			name:   "negative_values_found",
			nums:   []int{0, 2, 4, -10, -5, -3},
			target: -5,
			want:   4,
		},
		{
			name:   "negative_values_missing",
			nums:   []int{0, 2, 4, -10, -5, -3},
			target: 1,
			want:   -1,
		},

		// Границы ограничений
		{
			name:   "min_constraint_found",
			nums:   []int{0, 5000, 10000, -10000, -5000},
			target: -10000,
			want:   3,
		},
		{
			name:   "max_constraint_found",
			nums:   []int{0, 5000, 10000, -10000, -5000},
			target: 10000,
			want:   2,
		},

		// Цель рядом с точкой разрыва
		{
			name:   "target_before_pivot",
			nums:   []int{30, 40, 50, 5, 10, 20},
			target: 50,
			want:   2,
		},
		{
			name:   "target_at_pivot",
			nums:   []int{30, 40, 50, 5, 10, 20},
			target: 5,
			want:   3,
		},
		{
			name:   "target_after_pivot",
			nums:   []int{30, 40, 50, 5, 10, 20},
			target: 20,
			want:   5,
		},
		{
			name:   "target_missing_around_pivot",
			nums:   []int{30, 40, 50, 5, 10, 20},
			target: 35,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchInRotatedArray(tt.nums, tt.target)
			if got != tt.want {
				t.Fatalf("search(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
