// https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number

package array

import (
	"reflect"
	"testing"
)

func smallerNumbersThanCurrent(nums []int) []int {
	length := len(nums)
	res := make([]int, length, len(nums))

	for i, n := range nums {
		for _, v := range nums {
			if v < n {
				res[i]++
			}
		}
	}
	return res
}

func TestSmallerNumbersThanCurrent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example 1",
			nums: []int{8, 1, 2, 2, 3},
			want: []int{4, 0, 1, 1, 3},
		},
		{
			name: "example 2",
			nums: []int{6, 5, 4, 8},
			want: []int{2, 1, 0, 3},
		},
		{
			name: "example 3 all equal",
			nums: []int{7, 7, 7, 7},
			want: []int{0, 0, 0, 0},
		},
		{
			name: "minimum length distinct",
			nums: []int{0, 1},
			want: []int{0, 1},
		},
		{
			name: "minimum length equal",
			nums: []int{5, 5},
			want: []int{0, 0},
		},
		{
			name: "strictly increasing",
			nums: []int{0, 1, 2, 3, 4},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "strictly decreasing",
			nums: []int{4, 3, 2, 1, 0},
			want: []int{4, 3, 2, 1, 0},
		},
		{
			name: "contains zeros and duplicates",
			nums: []int{0, 0, 1, 1, 2},
			want: []int{0, 0, 2, 2, 4},
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0, 0},
			want: []int{0, 0, 0, 0},
		},
		{
			name: "all maximum values",
			nums: []int{100, 100, 100},
			want: []int{0, 0, 0},
		},
		{
			name: "mixed with repeated minimum",
			nums: []int{1, 0, 0, 2, 3},
			want: []int{2, 0, 0, 3, 4},
		},
		{
			name: "mixed with repeated maximum",
			nums: []int{5, 2, 5, 1},
			want: []int{2, 1, 2, 0},
		},
		{
			name: "one value repeated many times",
			nums: []int{3, 3, 3, 1, 2},
			want: []int{2, 2, 2, 0, 1},
		},
		{
			name: "duplicates on both sides",
			nums: []int{2, 5, 2, 8, 5, 6},
			want: []int{0, 2, 0, 5, 2, 4},
		},
		{
			name: "values span full constraint style range",
			nums: []int{100, 0, 50, 100, 0},
			want: []int{3, 0, 2, 3, 0},
		},
		{
			name: "single smallest in middle",
			nums: []int{9, 8, 0, 7, 6},
			want: []int{4, 3, 0, 2, 1},
		},
		{
			name: "single largest in middle",
			nums: []int{1, 2, 9, 2, 1},
			want: []int{0, 2, 4, 2, 0},
		},
		{
			name: "alternating duplicates",
			nums: []int{4, 1, 4, 1, 4, 1},
			want: []int{3, 0, 3, 0, 3, 0},
		},
		{
			name: "many repeated buckets",
			nums: []int{0, 2, 2, 2, 4, 4, 6},
			want: []int{0, 1, 1, 1, 4, 4, 6},
		},
		{
			name: "unsorted random looking case",
			nums: []int{8, 3, 5, 3, 9, 1},
			want: []int{4, 1, 3, 1, 5, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := smallerNumbersThanCurrent(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("smallerNumbersThanCurrent(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
