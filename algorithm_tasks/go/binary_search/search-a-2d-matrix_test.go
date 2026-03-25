// https://leetcode.com/problems/search-a-2d-matrix/

package binary_search

import "testing"

func searchMatrix(matrix [][]int, target int) bool {
	res := 0
	for _, row := range matrix {
		if row[0] <= target && row[len(row)-1] >= target {
			res = binarySearch(row, target)
			return res != -1
		}
	}
	return false
}

func binarySearch(row []int, target int) int {
	left := 0
	right := len(row) - 1

	for left <= right {
		mid := (left + right) / 2
		if row[mid] == target {
			return mid
		} else if row[mid] > target {
			right = mid - 1
		} else if row[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		// Примеры из условия
		{
			name:   "example_true",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 3,
			want:   true,
		},
		{
			name:   "example_false",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 13,
			want:   false,
		},

		// Минимальный размер матрицы
		{
			name:   "single_cell_found",
			matrix: [][]int{{1}},
			target: 1,
			want:   true,
		},
		{
			name:   "single_cell_not_found",
			matrix: [][]int{{1}},
			target: 0,
			want:   false,
		},

		// Одна строка
		{
			name:   "single_row_first",
			matrix: [][]int{{1, 3, 5, 7}},
			target: 1,
			want:   true,
		},
		{
			name:   "single_row_last",
			matrix: [][]int{{1, 3, 5, 7}},
			target: 7,
			want:   true,
		},
		{
			name:   "single_row_missing",
			matrix: [][]int{{1, 3, 5, 7}},
			target: 4,
			want:   false,
		},

		// Один столбец
		{
			name:   "single_col_first",
			matrix: [][]int{{1}, {3}, {5}, {7}},
			target: 1,
			want:   true,
		},
		{
			name:   "single_col_last",
			matrix: [][]int{{1}, {3}, {5}, {7}},
			target: 7,
			want:   true,
		},
		{
			name:   "single_col_missing",
			matrix: [][]int{{1}, {3}, {5}, {7}},
			target: 4,
			want:   false,
		},

		// Границы диапазона
		{
			name:   "first_element",
			matrix: [][]int{{1, 3, 5}, {7, 9, 11}},
			target: 1,
			want:   true,
		},
		{
			name:   "last_element",
			matrix: [][]int{{1, 3, 5}, {7, 9, 11}},
			target: 11,
			want:   true,
		},

		// Значение между строками, которого нет
		{
			name:   "between_rows_missing",
			matrix: [][]int{{1, 2, 3}, {10, 11, 12}},
			target: 5,
			want:   false,
		},

		// Отрицательные числа
		{
			name:   "negative_found",
			matrix: [][]int{{-10, -5, -1}, {0, 3, 8}},
			target: -5,
			want:   true,
		},
		{
			name:   "negative_missing",
			matrix: [][]int{{-10, -5, -1}, {0, 3, 8}},
			target: -2,
			want:   false,
		},

		// Ноль и переход через 0
		{
			name:   "zero_found",
			matrix: [][]int{{-3, -1, 0}, {2, 4, 6}},
			target: 0,
			want:   true,
		},
		{
			name:   "zero_missing",
			matrix: [][]int{{-3, -1, 1}, {2, 4, 6}},
			target: 0,
			want:   false,
		},

		// Значения около границ 10^4
		{
			name:   "min_constraint_value",
			matrix: [][]int{{-10000, -9999}, {-5, 0}, {9999, 10000}},
			target: -10000,
			want:   true,
		},
		{
			name:   "max_constraint_value",
			matrix: [][]int{{-10000, -9999}, {-5, 0}, {9999, 10000}},
			target: 10000,
			want:   true,
		},
		{
			name:   "constraint_value_missing",
			matrix: [][]int{{-10000, -9999}, {-5, 0}, {9999, 10000}},
			target: 5000,
			want:   false,
		},

		// Середина "виртуально расплющенного" массива
		{
			name: "middle_element_found",
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			target: 7,
			want:   true,
		},
		{
			name: "middle_element_missing",
			matrix: [][]int{
				{1, 2, 3, 4},
				{6, 7, 8, 9},
				{11, 12, 13, 14},
			},
			target: 5,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchMatrix(tt.matrix, tt.target)
			if got != tt.want {
				t.Fatalf("searchMatrix(%v, %d) = %v, want %v", tt.matrix, tt.target, got, tt.want)
			}
		})
	}
}
