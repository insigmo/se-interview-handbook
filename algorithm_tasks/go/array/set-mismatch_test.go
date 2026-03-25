// https://leetcode.com/problems/set-mismatch

package array

import (
	"reflect"
	"testing"
)

// d - duplicate
// m - missign
//
// d - m = x -> d = x + m
//
// d^2 - m^2 = y -> (x+m)^2 - m^2 = y ->
// -> x^2 + 2*x*m + m^2 - m^2 = y ->
// -> x^2 + 2*x*m = y ->
// -> 2*x*m = y - x^2 ->
// -> m = (y - x^2) / 2*x
//
// we have 2 formulas for "d" and "m"
// d = x + m
// m = (y - x^2) / 2*x
//
// for example
//
// [1 2 2 4]
// [1 2 3 4]
//
//	 d = 2
//		m = 3
//
// x = d - m -> x = 2 - 3 -> -1
// y = d^2 - m^2 -> 2^ - 3^2 = -5
// d = x + m -> -1 + m
// m = (y - x^2) / 2*x -> m = (-5 - -1^2) / 2*-1 -> -6 / -2 = 3
// d = -1 + m -> -1 + 3 = 2
// we found that d = 2 and m = 3
// result is []int{d, m}

func findErrorNums(nums []int) []int {
	length := len(nums)
	x := 0 // duplicate - missing
	y := 0 // duplicate^2 - missing^2

	// find duplicates
	for i := 1; i < length+1; i++ {
		x += nums[i-1] - i
		y += nums[i-1]*nums[i-1] - i*i
	}

	missing := (y - x*x) / (2 * x)
	duplicate := missing + x
	return []int{duplicate, missing}
}

func TestFindErrorNums(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example 1",
			nums: []int{1, 2, 2, 4},
			want: []int{2, 3},
		},
		{
			name: "example 2 minimum length",
			nums: []int{1, 1},
			want: []int{1, 2},
		},
		{
			name: "duplicate is last value missing in middle",
			nums: []int{1, 2, 3, 4, 4},
			want: []int{4, 5},
		},
		{
			name: "duplicate is first value missing second",
			nums: []int{1, 1, 3, 4, 5},
			want: []int{1, 2},
		},
		{
			name: "missing first duplicate in middle",
			nums: []int{2, 2},
			want: []int{2, 1},
		},
		{
			name: "missing last duplicate in middle",
			nums: []int{1, 2, 3, 3},
			want: []int{3, 4},
		},
		{
			name: "unsorted valid input",
			nums: []int{3, 1, 2, 5, 3},
			want: []int{3, 4},
		},
		{
			name: "duplicate near beginning",
			nums: []int{2, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: []int{2, 1},
		},
		{
			name: "duplicate near end",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9},
			want: []int{9, 10},
		},
		{
			name: "missing in center duplicate at end",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 8, 10},
			want: []int{8, 9},
		},
		{
			name: "duplicate repeated value one missing middle",
			nums: []int{1, 5, 3, 2, 2, 6, 7, 8, 9, 10},
			want: []int{2, 4},
		},
		{
			name: "small odd length case",
			nums: []int{1, 3, 3},
			want: []int{3, 2},
		},
		{
			name: "small even length unsorted",
			nums: []int{4, 2, 1, 4},
			want: []int{4, 3},
		},
		{
			name: "duplicate in middle missing first",
			nums: []int{2, 3, 4, 5, 5, 6, 7, 8, 9, 10},
			want: []int{5, 1},
		},
		{
			name: "duplicate in middle missing last",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 7},
			want: []int{7, 8},
		},
		{
			name: "longer shuffled input",
			nums: []int{8, 7, 3, 5, 4, 6, 2, 8, 9, 10, 1},
			want: []int{8, 11},
		},
		{
			name: "another shuffled input",
			nums: []int{6, 4, 2, 1, 3, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			want: []int{6, 5},
		},
		{
			name: "duplicate high value missing low value",
			nums: []int{4, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			want: []int{4, 1},
		},
		{
			name: "duplicate low value missing high value",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 14},
			want: []int{14, 15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findErrorNums(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("findErrorNums(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
