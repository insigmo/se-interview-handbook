// https://leetcode.com/problems/sliding-window-median

package heapq

import (
	"container/heap"
	"math"
	"testing"
)

type MinH []int

func (h MinH) Len() int           { return len(h) }
func (h MinH) Less(i, j int) bool { return h[i] < h[j] }
func (h MinH) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinH) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinH) Pop() any          { old := *h; x := old[len(old)-1]; *h = old[:len(old)-1]; return x }

type MaxH []int

func (h MaxH) Len() int           { return len(h) }
func (h MaxH) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxH) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxH) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxH) Pop() any          { old := *h; x := old[len(old)-1]; *h = old[:len(old)-1]; return x }

func medianSlidingWindow(nums []int, k int) []float64 {
	small := MaxH{}
	large := MinH{}
	isEven := k%2 == 0
	n := len(nums)

	for i := 0; i < k; i++ {
		heap.Push(&small, nums[i])
	}

	for i := 0; i < k/2; i++ {
		heap.Push(&large, heap.Pop(&small))
	}

	outgoing := make(map[int]int)
	out := make([]float64, 0, n-k+1)
	var balance int

	for i := k; i <= n; i++ {
		if isEven {
			out = append(out, (float64(small[0])+float64(large[0]))/2)
		} else {
			out = append(out, float64(small[0]))
		}

		if i >= n {
			break
		}

		balance = 0

		outgoing[nums[i-k]]++

		if nums[i-k] <= small[0] {
			balance--
		} else {
			balance++
		}

		if nums[i] <= small[0] {
			heap.Push(&small, nums[i])
			balance++
		} else {
			heap.Push(&large, nums[i])
			balance--
		}

		if balance < 0 {
			heap.Push(&small, heap.Pop(&large))
		}

		if balance > 0 {
			heap.Push(&large, heap.Pop(&small))
		}

		for small.Len() > 0 && outgoing[small[0]] > 0 {
			outgoing[small[0]]--
			heap.Pop(&small)
		}

		for large.Len() > 0 && outgoing[large[0]] > 0 {
			outgoing[large[0]]--
			heap.Pop(&large)
		}
	}

	return out
}

func TestMedianSlidingWindow(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []float64
	}{
		{
			name: "odd k basic example",
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			want: []float64{1.0, -1.0, -1.0, 3.0, 5.0, 6.0},
		},
		{
			name: "even k averages two middle elements",
			nums: []int{1, 3, 5, 2, 8, 4},
			k:    2,
			want: []float64{2.0, 4.0, 3.5, 5.0, 6.0},
		},
		{
			name: "k equals length of array odd",
			nums: []int{3, 1, 2},
			k:    3,
			want: []float64{2.0},
		},
		{
			name: "k equals length of array even",
			nums: []int{3, 1, 2, 4},
			k:    4,
			want: []float64{2.5},
		},
		{
			name: "k equals 1 every element is its own median",
			nums: []int{3, 1, 5, 2},
			k:    1,
			want: []float64{3.0, 1.0, 5.0, 2.0},
		},
		{
			name: "all elements identical",
			nums: []int{2, 2, 2, 2, 2},
			k:    3,
			want: []float64{2.0, 2.0, 2.0},
		},
		{
			name: "all negative numbers",
			nums: []int{-5, -3, -1, -2, -4},
			k:    3,
			want: []float64{-3.0, -2.0, -2.0},
		},
		{
			name: "duplicates in window",
			nums: []int{1, 1, 1, 1, 1},
			k:    3,
			want: []float64{1.0, 1.0, 1.0},
		},
		{
			name: "single element array",
			nums: []int{5},
			k:    1,
			want: []float64{5.0},
		},
		{
			name: "large values even k no overflow",
			nums: []int{2147483647, 2147483647},
			k:    2,
			want: []float64{2147483647.0},
		},
		{
			name: "descending order odd k",
			nums: []int{9, 7, 5, 3, 1},
			k:    3,
			want: []float64{7.0, 5.0, 3.0},
		},
		{
			name: "ascending order even k",
			nums: []int{1, 2, 3, 4, 5, 6},
			k:    4,
			want: []float64{2.5, 3.5, 4.5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := medianSlidingWindow(tt.nums, tt.k)
			if len(got) != len(tt.want) {
				t.Fatalf("medianSlidingWindow(%v, %d) len=%d, want len=%d", tt.nums, tt.k, len(got), len(tt.want))
			}
			for i := range got {
				if math.Abs(got[i]-tt.want[i]) > 1e-5 {
					t.Fatalf("medianSlidingWindow(%v, %d)[%d] = %v, want %v", tt.nums, tt.k, i, got[i], tt.want[i])
				}
			}
		})
	}
}
