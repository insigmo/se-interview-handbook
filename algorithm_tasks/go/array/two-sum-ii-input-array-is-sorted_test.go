// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted

package array

import (
	"reflect"
	"testing"
)

// rename to twoSum
func twoSumII(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{}
}

func TestTwoSumII(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			name:    "example 1",
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		{
			name:    "example 2",
			numbers: []int{2, 3, 4},
			target:  6,
			want:    []int{1, 3},
		},
		{
			name:    "example 3",
			numbers: []int{-1, 0},
			target:  -1,
			want:    []int{1, 2},
		},
		{
			name:    "minimum length positive",
			numbers: []int{1, 2},
			target:  3,
			want:    []int{1, 2},
		},
		{
			name:    "minimum length with duplicates",
			numbers: []int{3, 3},
			target:  6,
			want:    []int{1, 2},
		},
		{
			name:    "pair at both ends",
			numbers: []int{1, 2, 3, 4, 9},
			target:  10,
			want:    []int{1, 5},
		},
		{
			name:    "pair in the middle",
			numbers: []int{1, 3, 4, 6, 8, 10},
			target:  10,
			want:    []int{3, 4},
		},
		{
			name:    "all negative values",
			numbers: []int{-10, -8, -3, -1},
			target:  -11,
			want:    []int{1, 4},
		},
		{
			name:    "negative and positive mix",
			numbers: []int{-7, -3, 2, 3, 11},
			target:  0,
			want:    []int{2, 4},
		},
		{
			name:    "includes zero target zero",
			numbers: []int{-4, -1, 0, 2, 5},
			target:  1,
			want:    []int{1, 5},
		},
		{
			name:    "two zeros",
			numbers: []int{-3, 0, 0, 4},
			target:  0,
			want:    []int{2, 3},
		},
		{
			name:    "duplicate values valid pair",
			numbers: []int{1, 2, 2, 4, 9},
			target:  4,
			want:    []int{2, 3},
		},
		{
			name:    "repeated low values unique solution",
			numbers: []int{1, 1, 3, 5},
			target:  2,
			want:    []int{1, 2},
		},
		{
			name:    "repeated high values unique solution",
			numbers: []int{1, 4, 6, 6, 10},
			target:  12,
			want:    []int{3, 4},
		},
		{
			name:    "target uses first two elements",
			numbers: []int{2, 4, 7, 11, 15},
			target:  6,
			want:    []int{1, 2},
		},
		{
			name:    "target uses last two elements",
			numbers: []int{1, 2, 5, 7, 9},
			target:  16,
			want:    []int{4, 5},
		},
		{
			name:    "large magnitude within constraints",
			numbers: []int{-1000, -20, 0, 20, 1000},
			target:  0,
			want:    []int{1, 5},
		},
		{
			name:    "solution near center with negatives",
			numbers: []int{-9, -4, -1, 3, 8, 12},
			target:  4,
			want:    []int{2, 5},
		},
		{
			name:    "monotonic with one clear answer",
			numbers: []int{1, 5, 6, 8, 20, 25},
			target:  26,
			want:    []int{1, 6},
		},
		{
			name:    "another sorted duplicate case",
			numbers: []int{2, 2, 3, 7, 11},
			target:  4,
			want:    []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSumII(tt.numbers, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("twoSum(%v, %d) = %v, want %v", tt.numbers, tt.target, got, tt.want)
			}
		})
	}
}
