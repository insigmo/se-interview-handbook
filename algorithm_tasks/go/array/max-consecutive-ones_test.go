// https://leetcode.com/problems/max-consecutive-ones

package array

import "testing"

func findMaxConsecutiveOnes(nums []int) int {
	counter := 0
	maxCounter := 0
	for _, i := range nums {
		if i == 1 {
			counter++
		} else {
			counter = 0
		}

		if maxCounter < counter {
			maxCounter++
		}
	}
	return maxCounter
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{1, 1, 0, 1, 1, 1},
			want: 3,
		},
		{
			name: "example 2",
			nums: []int{1, 0, 1, 1, 0, 1},
			want: 2,
		},
		{
			name: "single one",
			nums: []int{1},
			want: 1,
		},
		{
			name: "all ones",
			nums: []int{1, 1, 1, 1, 1},
			want: 5,
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0, 0},
			want: 0,
		},
		{
			name: "ones at beginning",
			nums: []int{1, 1, 1, 0, 0, 1},
			want: 3,
		},
		{
			name: "ones at end",
			nums: []int{0, 1, 0, 1, 1, 1},
			want: 3,
		},
		{
			name: "alternating starts with one",
			nums: []int{1, 0, 1, 0, 1, 0, 1},
			want: 1,
		},
		{
			name: "alternating starts with zero",
			nums: []int{0, 1, 0, 1, 0, 1, 0},
			want: 1,
		},
		{
			name: "long streak in middle",
			nums: []int{0, 1, 1, 1, 1, 0, 1, 1},
			want: 4,
		},
		{
			name: "multiple equal streaks",
			nums: []int{1, 1, 0, 1, 1, 0, 1, 1},
			want: 2,
		},
		{
			name: "zero breaks long run",
			nums: []int{1, 1, 1, 0, 1, 1, 1, 1},
			want: 4,
		},
		{
			name: "leading and trailing zeros",
			nums: []int{0, 0, 1, 1, 1, 0, 0},
			want: 3,
		},
		{
			name: "many zeros between ones",
			nums: []int{1, 0, 0, 0, 1, 0, 0, 1},
			want: 1,
		},
		{
			name: "max run length two",
			nums: []int{0, 1, 1, 0, 1, 0, 1, 1, 0},
			want: 2,
		},
		{
			name: "max run at very end",
			nums: []int{0, 0, 1, 0, 1, 1, 1, 1},
			want: 4,
		},
		{
			name: "max run at very beginning",
			nums: []int{1, 1, 1, 1, 0, 1, 0, 0},
			want: 4,
		},
		{
			name: "two long runs pick larger",
			nums: []int{1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 1},
			want: 5,
		},
		{
			name: "minimal break splits runs",
			nums: []int{1, 1, 1, 0, 1, 1, 1},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxConsecutiveOnes(tt.nums)
			if got != tt.want {
				t.Fatalf("findMaxConsecutiveOnes(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
