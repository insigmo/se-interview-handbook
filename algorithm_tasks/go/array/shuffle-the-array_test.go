// https://leetcode.com/problems/shuffle-the-array

package array

import (
	"reflect"
	"testing"
)

func shuffle(nums []int, n int) []int {
	res := make([]int, 2*n)
	for i := 0; i < n; i++ {
		res[2*i] = nums[i]
		res[2*i+1] = nums[i+n]
	}
	return res
}

func TestShuffle(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		n    int
		want []int
	}{
		{
			name: "example 1",
			nums: []int{2, 5, 1, 3, 4, 7},
			n:    3,
			want: []int{2, 3, 5, 4, 1, 7},
		},
		{
			name: "example 2",
			nums: []int{1, 2, 3, 4, 4, 3, 2, 1},
			n:    4,
			want: []int{1, 4, 2, 3, 3, 2, 4, 1},
		},
		{
			name: "example 3",
			nums: []int{1, 1, 2, 2},
			n:    2,
			want: []int{1, 2, 1, 2},
		},
		{
			name: "minimum n",
			nums: []int{9, 8},
			n:    1,
			want: []int{9, 8},
		},
		{
			name: "two pairs increasing",
			nums: []int{1, 2, 3, 4},
			n:    2,
			want: []int{1, 3, 2, 4},
		},
		{
			name: "all same values",
			nums: []int{5, 5, 5, 5, 5, 5},
			n:    3,
			want: []int{5, 5, 5, 5, 5, 5},
		},
		{
			name: "strictly increasing halves",
			nums: []int{1, 2, 3, 4, 5, 6},
			n:    3,
			want: []int{1, 4, 2, 5, 3, 6},
		},
		{
			name: "repeated values across halves",
			nums: []int{2, 2, 2, 3, 3, 3},
			n:    3,
			want: []int{2, 3, 2, 3, 2, 3},
		},
		{
			name: "alternating result pattern",
			nums: []int{1, 1, 1, 2, 2, 2},
			n:    3,
			want: []int{1, 2, 1, 2, 1, 2},
		},
		{
			name: "larger example",
			nums: []int{10, 20, 30, 40, 50, 60, 70, 80},
			n:    4,
			want: []int{10, 50, 20, 60, 30, 70, 40, 80},
		},
		{
			name: "contains lower and upper constraint style values",
			nums: []int{1, 1000, 500, 999},
			n:    2,
			want: []int{1, 500, 1000, 999},
		},
		{
			name: "palindrome like halves",
			nums: []int{7, 8, 8, 7},
			n:    2,
			want: []int{7, 8, 8, 7},
		},
		{
			name: "odd looking values but valid layout",
			nums: []int{4, 9, 1, 6, 7, 3},
			n:    3,
			want: []int{4, 6, 9, 7, 1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shuffle(tt.nums, tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("shuffle(%v, %d) = %v, want %v", tt.nums, tt.n, got, tt.want)
			}
		})
	}
}
