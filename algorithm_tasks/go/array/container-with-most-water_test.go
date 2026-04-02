// https://leetcode.com/problems/container-with-most-water/description/

package array

import "testing"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	v := 0
	minValue := 0

	for left < right {
		leftHeight := height[left]
		rightHeight := height[right]

		if leftHeight > rightHeight {
			minValue = rightHeight
		} else {
			minValue = leftHeight
		}
		currentVolume := (right - left) * minValue

		if v < currentVolume {
			v = currentVolume
		}

		if leftHeight > rightHeight {
			right--
		} else {
			left++
		}
	}
	return v
}

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height   []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{3, 7}, 3},
		{[]int{1000, 1000}, 1000},
		{[]int{5, 5, 5, 5}, 15},
		{[]int{1, 2, 3, 4, 5}, 6},
		{[]int{5, 4, 3, 2, 1}, 6},
		{[]int{0, 2, 0}, 0},
		{[]int{1, 0, 1}, 2},
		{[]int{1, 100, 1}, 2},
		{[]int{2, 3, 10, 5, 7, 8, 9}, 36},
		{[]int{0, 0}, 0},
		{[]int{1, 2, 1}, 2},
		{[]int{4, 3, 2, 1, 4}, 16},
	}

	for _, tt := range tests {
		result := maxArea(tt.height)
		if result != tt.expected {
			t.Errorf("maxArea(%v) = %d, want %d", tt.height, result, tt.expected)
		}
	}
}
