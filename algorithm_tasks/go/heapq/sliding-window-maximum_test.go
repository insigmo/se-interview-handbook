// https://leetcode.com/problems/sliding-window-maximum/

package heapq

import (
	"fmt"
	"reflect"
	"testing"
)

const N = 100000

var deque, ans [N]int

func maxSlidingWindow(nums []int, k int) []int {
	// front is the oldest, back is the newest
	front, back := 0, 0
	for i, val := range nums {
		// limit deque to window size
		for ; front < back && i-deque[front]+1 > k; front++ {
			fmt.Println(front)
		}
		fmt.Println()
		// discard values smaller than current
		for ; front < back && nums[deque[back-1]] <= val; back-- {
			fmt.Println(back)
		}
		// queue current
		deque[back] = i
		back++
		// record max ans
		ans[i] = nums[deque[front]]
	}
	return ans[k-1 : len(nums)]
}

func TestMaxSlidingWindow(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "basic example",
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "with 1 k",
			nums: []int{1, -1},
			k:    1,
			want: []int{1, -1},
		},
		{
			name: "window size equals array length",
			nums: []int{4, 2, 7, 1},
			k:    4,
			want: []int{7},
		},
		{
			name: "window size one",
			nums: []int{5, 3, 8, 1},
			k:    1,
			want: []int{5, 3, 8, 1},
		},
		{
			name: "single element array",
			nums: []int{42},
			k:    1,
			want: []int{42},
		},
		{
			name: "strictly increasing array",
			nums: []int{1, 2, 3, 4, 5},
			k:    3,
			want: []int{3, 4, 5},
		},
		{
			name: "strictly decreasing array",
			nums: []int{5, 4, 3, 2, 1},
			k:    3,
			want: []int{5, 4, 3},
		},
		{
			name: "all elements equal",
			nums: []int{7, 7, 7, 7, 7},
			k:    3,
			want: []int{7, 7, 7},
		},
		{
			name: "all negative numbers",
			nums: []int{-5, -3, -7, -1, -4},
			k:    2,
			want: []int{-3, -3, -1, -1},
		},
		{
			name: "mixed positive and negative",
			nums: []int{-2, 1, -3, 4, -1},
			k:    3,
			want: []int{1, 4, 4},
		},
		{
			name: "max at start then drops",
			nums: []int{9, 1, 1, 1, 1},
			k:    3,
			want: []int{9, 1, 1},
		},
		{
			name: "max at end of each window",
			nums: []int{1, 1, 1, 9, 1},
			k:    3,
			want: []int{1, 9, 9},
		},
		{
			name: "window size two",
			nums: []int{3, 1, 4, 1, 5, 9},
			k:    2,
			want: []int{3, 4, 4, 5, 9},
		},
		{
			name: "large negative values",
			nums: []int{-100, -200, -50, -300},
			k:    2,
			want: []int{-100, -50, -50},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSlidingWindow(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("maxSlidingWindow(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
