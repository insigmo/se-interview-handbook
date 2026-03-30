// https://leetcode.com/problems/3sum/

package array

import (
	"slices"
	"testing"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	target := 0
	res := make([][]int, 0, len(nums)/4)

	for i, a := range nums {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := a + nums[left] + nums[right]

			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++

				for nums[left] == nums[left-1] && left < right {
					left++
				}
			}
		}
	}
	return res
}

func TestThreeSum(t *testing.T) {
	normalizeTriplet := func(v []int) [3]int {
		a := [3]int{v[0], v[1], v[2]}
		if a[0] > a[1] {
			a[0], a[1] = a[1], a[0]
		}
		if a[1] > a[2] {
			a[1], a[2] = a[2], a[1]
		}
		if a[0] > a[1] {
			a[0], a[1] = a[1], a[0]
		}
		return a
	}

	buildSet := func(v [][]int) map[[3]int]struct{} {
		res := make(map[[3]int]struct{})
		for _, triplet := range v {
			if len(triplet) != 3 {
				continue
			}
			res[normalizeTriplet(triplet)] = struct{}{}
		}
		return res
	}

	assertEqualTriplets := func(t *testing.T, got [][]int, want [][]int) {
		t.Helper()

		for _, triplet := range got {
			if len(triplet) != 3 {
				t.Fatalf("each triplet must have length 3, got %v", triplet)
			}
			if triplet[0]+triplet[1]+triplet[2] != 0 {
				t.Fatalf("triplet must sum to zero, got %v", triplet)
			}
		}

		gotSet := buildSet(got)
		wantSet := buildSet(want)

		if len(got) != len(gotSet) {
			t.Fatalf("duplicate triplets found in result: %v", got)
		}

		if len(gotSet) != len(wantSet) {
			t.Fatalf("got %v, want %v", got, want)
		}

		for triplet := range wantSet {
			if _, ok := gotSet[triplet]; !ok {
				t.Fatalf("missing triplet %v in %v", triplet, got)
			}
		}

		for triplet := range gotSet {
			if _, ok := wantSet[triplet]; !ok {
				t.Fatalf("unexpected triplet %v in %v", triplet, got)
			}
		}
	}

	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example one",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "example two",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "example three",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "more than three zeroes still one triplet",
			nums: []int{0, 0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "minimum valid length",
			nums: []int{-1, 0, 1},
			want: [][]int{{-1, 0, 1}},
		},
		{
			name: "minimum invalid length",
			nums: []int{1, 2, -2},
			want: [][]int{},
		},
		{
			name: "all positive",
			nums: []int{1, 2, 3, 4, 5},
			want: [][]int{},
		},
		{
			name: "all negative",
			nums: []int{-5, -4, -3, -2, -1},
			want: [][]int{},
		},
		{
			name: "duplicate candidates produce one unique triplet",
			nums: []int{-2, 0, 0, 2, 2},
			want: [][]int{{-2, 0, 2}},
		},
		{
			name: "mixed duplicates and multiple answers",
			nums: []int{-2, -1, -1, 0, 1, 1, 2, 2},
			want: [][]int{{-2, 0, 2}, {-2, 1, 1}, {-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "extreme values within constraints",
			nums: []int{-100000, 100000, 0, -99999, 99999, 1, -1},
			want: [][]int{
				{-100000, 0, 100000},
				{-99999, 0, 99999},
				{-100000, 1, 99999},
				{-99999, -1, 100000},
				{-1, 0, 1},
			},
		},
		{
			name: "many duplicates around zero",
			nums: []int{-1, -1, -1, 0, 0, 0, 1, 1, 1},
			want: [][]int{{-1, 0, 1}, {0, 0, 0}},
		},
		{
			name: "single solution using repeated negatives",
			nums: []int{-4, -2, -2, -2, 6},
			want: [][]int{{-4, -2, 6}},
		},
		{
			name: "multiple valid triplets with shared elements",
			nums: []int{-4, -2, -1, 0, 1, 2, 3, 5},
			want: [][]int{{-4, -1, 5}, {-4, 1, 3}, {-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			assertEqualTriplets(t, got, tt.want)
		})
	}
}
