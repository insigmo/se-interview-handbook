// https://leetcode.com/problems/find-the-distance-value-between-two-arrays/

package binary_search_test

import (
	"math"
	"slices"
	"testing"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	slices.Sort(arr2)
	length := len(arr2)

	for _, num := range arr1 {
		ind, _ := slices.BinarySearch(arr2, num)
		minDist := math.MaxFloat64

		if ind > 0 {
			minDist = min(minDist, math.Abs(float64(num-arr2[ind-1])))
		}
		if ind < length {
			minDist = min(minDist, math.Abs(float64(num-arr2[ind])))
		}
		if minDist > float64(d) {
			res += 1
		}
	}

	return res

}

func TestFindTheDistanceValue(t *testing.T) {
	tests := []struct {
		name     string
		arr1     []int
		arr2     []int
		d        int
		expected int
	}{
		// --- из условия задачи ---
		{
			name:     "example 1",
			arr1:     []int{4, 5, 8},
			arr2:     []int{10, 9, 1, 8},
			d:        2,
			expected: 2,
		},
		{
			name:     "example 2",
			arr1:     []int{1, 4, 2, 3},
			arr2:     []int{-4, -3, 6, 10, 20, 30},
			d:        3,
			expected: 2,
		},
		{
			name:     "example 3",
			arr1:     []int{2, 1, 100, 3},
			arr2:     []int{-5, -2, 10, -3, 7},
			d:        6,
			expected: 1,
		},

		// --- граничные случаи ---
		{
			name:     "d = 0, нет совпадений",
			arr1:     []int{1, 3, 5},
			arr2:     []int{2, 4, 6},
			d:        0,
			expected: 3,
		},
		{
			name:     "d = 0, все совпадают",
			arr1:     []int{1, 2, 3},
			arr2:     []int{1, 2, 3},
			d:        0,
			expected: 0,
		},
		{
			name:     "d = 100, все в диапазоне",
			arr1:     []int{0},
			arr2:     []int{50},
			d:        100,
			expected: 0,
		},
		{
			name:     "один элемент в каждом, не попадает",
			arr1:     []int{1000},
			arr2:     []int{-1000},
			d:        1,
			expected: 1,
		},
		{
			name:     "один элемент в каждом, попадает",
			arr1:     []int{5},
			arr2:     []int{5},
			d:        0,
			expected: 0,
		},
		{
			name:     "отрицательные значения",
			arr1:     []int{-10, -5, 0},
			arr2:     []int{-100, -50},
			d:        3,
			expected: 3,
		},
		{
			name:     "arr1 один элемент, arr2 большой",
			arr1:     []int{500},
			arr2:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			d:        10,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findTheDistanceValue(tt.arr1, tt.arr2, tt.d)
			if got != tt.expected {
				t.Errorf("findTheDistanceValue(%v, %v, %d) = %d, want %d",
					tt.arr1, tt.arr2, tt.d, got, tt.expected)
			}
		})
	}
}
