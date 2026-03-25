// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array

package array

import (
	"reflect"
	"testing"
)

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	seen := make([]byte, n+1)
	for _, v := range nums {
		seen[v] = 1
	}

	res := make([]int, 0, n/2)
	for i := 1; i <= n; i++ {
		if seen[i] == 0 {
			res = append(res, i)
		}
	}

	return res
}

// LeetCode 448: return all numbers in the range [1, n] that do not appear in nums. [web:145][web:155]
func TestFindDisappearedNumbers(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example 1",
			nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			want: []int{5, 6},
		},
		{
			name: "example 2",
			nums: []int{1, 1},
			want: []int{2},
		},
		{
			name: "single element present",
			nums: []int{1},
			want: []int{},
		},
		{
			name: "all numbers present sorted",
			nums: []int{1, 2, 3, 4, 5},
			want: []int{},
		},
		{
			name: "all numbers present unsorted",
			nums: []int{5, 3, 1, 2, 4},
			want: []int{},
		},
		{
			name: "all same value missing many",
			nums: []int{2, 2, 2, 2},
			want: []int{1, 3, 4},
		},
		{
			name: "missing first number",
			nums: []int{2, 2, 3, 4},
			want: []int{1},
		},
		{
			name: "missing last number",
			nums: []int{1, 2, 3, 3},
			want: []int{4},
		},
		{
			name: "missing first and last",
			nums: []int{2, 2, 3, 3},
			want: []int{1, 4},
		},
		{
			name: "multiple missing in middle",
			nums: []int{1, 1, 4, 4},
			want: []int{2, 3},
		},
		{
			name: "alternating duplicates",
			nums: []int{1, 3, 1, 3, 5, 5},
			want: []int{2, 4, 6},
		},
		{
			name: "duplicates create consecutive missing tail",
			nums: []int{1, 2, 3, 4, 4, 4},
			want: []int{5, 6},
		},
		{
			name: "duplicates create consecutive missing head",
			nums: []int{3, 3, 3},
			want: []int{1, 2},
		},
		{
			name: "repeated first value",
			nums: []int{1, 1, 1, 1, 5},
			want: []int{2, 3, 4},
		},
		{
			name: "repeated last value",
			nums: []int{5, 5, 5, 2, 1},
			want: []int{3, 4},
		},
		{
			name: "single missing in larger array",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 9, 9},
			want: []int{8},
		},
		{
			name: "several missing scattered",
			nums: []int{10, 2, 5, 10, 9, 1, 1, 4, 7, 2},
			want: []int{3, 6, 8},
		},
		{
			name: "many duplicates but one of each edge present",
			nums: []int{1, 6, 6, 6, 6, 6},
			want: []int{2, 3, 4, 5},
		},
		{
			name: "two numbers only one missing second",
			nums: []int{1, 1},
			want: []int{2},
		},
		{
			name: "two numbers only one missing first",
			nums: []int{2, 2},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDisappearedNumbers(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("findDisappearedNumbers(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
