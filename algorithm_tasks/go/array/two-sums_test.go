// https://leetcode.com/problems/two-sum/

package array

import "testing"

func twoSum(nums []int, target int) []int {
	switch len(nums) {
	case 0:
		return []int{}
	case 1:
		return nums
	}

	numsMap := make(map[int]int)
	for i, v := range nums {
		numsMap[v] = i
	}

	for i, v := range nums {
		if j, ok := numsMap[target-v]; ok && i != j {
			return []int{i, j}
		}
	}
	return []int{}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
	}{
		{
			name:   "example 1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
		},
		{
			name:   "example 2",
			nums:   []int{3, 2, 4},
			target: 6,
		},
		{
			name:   "example 3 duplicates",
			nums:   []int{3, 3},
			target: 6,
		},
		{
			name:   "two elements only",
			nums:   []int{1, 9},
			target: 10,
		},
		{
			name:   "uses negative and positive",
			nums:   []int{-3, 4, 3, 90},
			target: 0,
		},
		{
			name:   "uses zero twice",
			nums:   []int{0, 4, 3, 0},
			target: 0,
		},
		{
			name:   "answer at the end",
			nums:   []int{5, 1, 9, 2, 8},
			target: 10,
		},
		{
			name:   "answer includes first element",
			nums:   []int{10, -2, 1, 3, 7},
			target: 8,
		},
		{
			name:   "duplicate values but unique pair by indices",
			nums:   []int{1, 5, 1, 7},
			target: 2,
		},
		{
			name:   "negative pair",
			nums:   []int{-5, -2, -3, -4},
			target: -8,
		},
		{
			name:   "mixed values unique answer",
			nums:   []int{8, 6, 7, 2, 15, -1},
			target: 14,
		},
		{
			name:   "large magnitude values",
			nums:   []int{1000000000, -1000000000, 5, 12},
			target: 0,
		},
		{
			name:   "pair not adjacent",
			nums:   []int{11, 15, 2, 7},
			target: 9,
		},
		{
			name:   "zero with negative",
			nums:   []int{0, -1, 2, 1},
			target: 1,
		},
		{
			name:   "same value needed from two positions",
			nums:   []int{4, 6, 5, 5},
			target: 10,
		},
		{
			name:   "many distractors",
			nums:   []int{13, 22, 1, 4, 8, 17, 9},
			target: 26,
		},
		{
			name:   "solution uses middle elements",
			nums:   []int{20, 1, 14, 6, 9},
			target: 20,
		},
		{
			name:   "small negative and positive",
			nums:   []int{-1, 0, 1},
			target: 0,
		},
		{
			name:   "repeated zeros and unique solution",
			nums:   []int{1, 0, 3, 0},
			target: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)

			if len(got) != 2 {
				t.Fatalf("twoSum(%v, %d) returned %v, want exactly 2 indices", tt.nums, tt.target, got)
			}

			i, j := got[0], got[1]

			if i < 0 || i >= len(tt.nums) || j < 0 || j >= len(tt.nums) {
				t.Fatalf("twoSum(%v, %d) returned out-of-range indices %v", tt.nums, tt.target, got)
			}

			if i == j {
				t.Fatalf("twoSum(%v, %d) returned the same index twice: %v", tt.nums, tt.target, got)
			}

			if tt.nums[i]+tt.nums[j] != tt.target {
				t.Fatalf(
					"twoSum(%v, %d) returned indices %v with values %d and %d, sum = %d",
					tt.nums,
					tt.target,
					got,
					tt.nums[i],
					tt.nums[j],
					tt.nums[i]+tt.nums[j],
				)
			}
		})
	}
}
