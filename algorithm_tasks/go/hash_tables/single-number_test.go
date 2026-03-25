// https://leetcode.com/problems/single-number/description/

package hash_tables

import "testing"

func singleNumber(nums []int) int {
	s := 0
	for _, i := range nums {
		s ^= i
	}
	return s
}

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			name: "example 2",
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			name: "example 3 single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "single element zero",
			nums: []int{0},
			want: 0,
		},
		{
			name: "unique at beginning",
			nums: []int{9, 1, 1, 2, 2, 3, 3},
			want: 9,
		},
		{
			name: "unique in middle",
			nums: []int{5, 6, 6, 7, 7, 8, 8},
			want: 5,
		},
		{
			name: "unique at end",
			nums: []int{1, 1, 2, 2, 3, 3, 4},
			want: 4,
		},
		{
			name: "includes zero duplicated",
			nums: []int{0, 1, 1, 2, 2},
			want: 0,
		},
		{
			name: "unique is zero",
			nums: []int{4, 4, 0, 9, 9},
			want: 0,
		},
		{
			name: "all negative except one",
			nums: []int{-1, -1, -2, -2, -3},
			want: -3,
		},
		{
			name: "unique negative among positives",
			nums: []int{7, 7, 8, 8, -11},
			want: -11,
		},
		{
			name: "mixed negative and positive",
			nums: []int{-10, 4, 4, -10, 99},
			want: 99,
		},
		{
			name: "smallest constraint style negative unique",
			nums: []int{-30000, 5, 5},
			want: -30000,
		},
		{
			name: "largest constraint style positive unique",
			nums: []int{30000, -7, -7},
			want: 30000,
		},
		{
			name: "unordered pairs",
			nums: []int{10, 3, 10, 4, 3},
			want: 4,
		},
		{
			name: "longer odd sized array",
			nums: []int{12, 1, 2, 1, 2, 3, 3, 4, 4},
			want: 12,
		},
		{
			name: "repeated pair values far apart",
			nums: []int{6, 1, 2, 1, 2, 6, 99},
			want: 99,
		},
		{
			name: "negative unique with zero pair",
			nums: []int{0, -5, 0},
			want: -5,
		},
		{
			name: "two pairs and unique negative",
			nums: []int{-2, 8, 8, 11, 11},
			want: -2,
		},
		{
			name: "many pairs one unique",
			nums: []int{14, 21, 14, 35, 21, 42, 35},
			want: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber(tt.nums)
			if got != tt.want {
				t.Fatalf("singleNumber(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
