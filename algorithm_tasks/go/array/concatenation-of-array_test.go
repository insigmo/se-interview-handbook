// https://leetcode.com/problems/concatenation-of-array

package array

import (
	"reflect"
	"testing"
)

func getConcatenation(nums []int) []int {
	var ans []int
	ans = append(ans, nums...)
	return append(ans, nums...)
}

func TestGetConcatenation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example 1",
			nums: []int{1, 2, 1},
			want: []int{1, 2, 1, 1, 2, 1},
		},
		{
			name: "example 2",
			nums: []int{1, 3, 2, 1},
			want: []int{1, 3, 2, 1, 1, 3, 2, 1},
		},
		{
			name: "single element",
			nums: []int{7},
			want: []int{7, 7},
		},
		{
			name: "two elements",
			nums: []int{4, 5},
			want: []int{4, 5, 4, 5},
		},
		{
			name: "all same values",
			nums: []int{9, 9, 9},
			want: []int{9, 9, 9, 9, 9, 9},
		},
		{
			name: "strictly increasing",
			nums: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5},
		},
		{
			name: "contains minimum constraint value",
			nums: []int{1, 1000},
			want: []int{1, 1000, 1, 1000},
		},
		{
			name: "mixed repeated values",
			nums: []int{2, 1, 2, 3},
			want: []int{2, 1, 2, 3, 2, 1, 2, 3},
		},
		{
			name: "palindrome like input",
			nums: []int{1, 2, 3, 2, 1},
			want: []int{1, 2, 3, 2, 1, 1, 2, 3, 2, 1},
		},
		{
			name: "larger input",
			nums: []int{8, 6, 7, 5, 3, 0, 9},
			want: []int{8, 6, 7, 5, 3, 0, 9, 8, 6, 7, 5, 3, 0, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getConcatenation(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("getConcatenation(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
