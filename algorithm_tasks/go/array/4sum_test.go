// https://leetcode.com/problems/4sum

package array

import (
	"slices"
	"testing"
)

type Summarizer struct {
	numbers      []int
	result       [][]int
	prefix       []int
	countInArray int
}

func New(numbers []int) *Summarizer {
	slices.Sort(numbers)
	return &Summarizer{
		numbers: numbers,
		result:  make([][]int, 0, len(numbers)/2),
		prefix:  make([]int, 0, len(numbers)/4),
	}
}

func (s *Summarizer) twoSumCalc(start, target int) {
	left, right := start, len(s.numbers)-1

	for left < right {
		sum := s.numbers[left] + s.numbers[right]
		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			s.result = append(s.result, []int{
				s.prefix[0],
				s.prefix[1],
				s.numbers[left],
				s.numbers[right],
			})

			left++

			for s.numbers[left] == s.numbers[left-1] && left < right {
				left++
			}
		}
	}
}

func (s *Summarizer) kSum(countInArray, start, target int) [][]int {
	if countInArray == 2 {
		s.twoSumCalc(start, target)
		return s.result
	}
	for i := start; i < len(s.numbers)-countInArray+1; i++ {
		if i > start && s.numbers[i] == s.numbers[i-1] {
			continue
		}
		s.prefix = append(s.prefix, s.numbers[i])
		s.kSum(countInArray-1, i+1, target-s.numbers[i])
		s.prefix = s.prefix[:len(s.prefix)-1]
	}
	return s.result
}

func fourSum(nums []int, target int) [][]int {
	summarizer := New(nums)
	return summarizer.kSum(4, 0, target)
}

func TestFourSum(t *testing.T) {
	normalizeQuad := func(v []int) [4]int {
		a := [4]int{v[0], v[1], v[2], v[3]}
		for i := 0; i < 4; i++ {
			for j := i + 1; j < 4; j++ {
				if a[i] > a[j] {
					a[i], a[j] = a[j], a[i]
				}
			}
		}
		return a
	}

	buildSet := func(v [][]int) map[[4]int]struct{} {
		res := make(map[[4]int]struct{})
		for _, quad := range v {
			if len(quad) != 4 {
				continue
			}
			res[normalizeQuad(quad)] = struct{}{}
		}
		return res
	}

	assertEqualQuads := func(t *testing.T, got [][]int, want [][]int, target int) {
		t.Helper()

		for _, quad := range got {
			if len(quad) != 4 {
				t.Fatalf("each quadruplet must have length 4, got %v", quad)
			}
			sum := int64(quad[0]) + int64(quad[1]) + int64(quad[2]) + int64(quad[3])
			if sum != int64(target) {
				t.Fatalf("quadruplet must sum to target %d, got %v", target, quad)
			}
		}

		gotSet := buildSet(got)
		wantSet := buildSet(want)

		if len(got) != len(gotSet) {
			t.Fatalf("duplicate quadruplets found in result: %v", got)
		}

		if len(gotSet) != len(wantSet) {
			t.Fatalf("got %v, want %v", got, want)
		}

		for quad := range wantSet {
			if _, ok := gotSet[quad]; !ok {
				t.Fatalf("missing quadruplet %v in %v", quad, got)
			}
		}

		for quad := range gotSet {
			if _, ok := wantSet[quad]; !ok {
				t.Fatalf("unexpected quadruplet %v in %v", quad, got)
			}
		}
	}

	tests := []struct {
		name   string
		nums   []int
		target int
		want   [][]int
	}{
		{
			name:   "example one",
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			want:   [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			name:   "example two",
			nums:   []int{2, 2, 2, 2, 2},
			target: 8,
			want:   [][]int{{2, 2, 2, 2}},
		},
		{
			name:   "less than four elements",
			nums:   []int{1, 2, 3},
			target: 6,
			want:   [][]int{},
		},
		{
			name:   "exactly four elements valid",
			nums:   []int{1, 2, 3, 4},
			target: 10,
			want:   [][]int{{1, 2, 3, 4}},
		},
		{
			name:   "exactly four elements invalid",
			nums:   []int{1, 2, 3, 4},
			target: 11,
			want:   [][]int{},
		},
		{
			name:   "all zeroes deduplicated",
			nums:   []int{0, 0, 0, 0, 0, 0},
			target: 0,
			want:   [][]int{{0, 0, 0, 0}},
		},
		{
			name:   "all positive no solution",
			nums:   []int{1, 2, 3, 4, 5, 6},
			target: 0,
			want:   [][]int{},
		},
		{
			name:   "all negative no solution",
			nums:   []int{-6, -5, -4, -3, -2, -1},
			target: 0,
			want:   [][]int{},
		},
		{
			name:   "duplicate values multiple unique answers",
			nums:   []int{-2, -1, -1, 1, 1, 2, 2},
			target: 0,
			want:   [][]int{{-2, -1, 1, 2}, {-1, -1, 1, 1}},
		},

		{
			name:   "target negative",
			nums:   []int{-5, -4, -3, -2, -1, 0, 1, 2},
			target: -10,
			want:   [][]int{{-5, -4, -3, 2}, {-5, -4, -2, 1}, {-5, -4, -1, 0}, {-5, -3, -2, 0}, {-4, -3, -2, -1}},
		},
		{
			name:   "extreme values require int64-safe sum",
			nums:   []int{1000000000, 1000000000, 1000000000, 1000000000, -1000000000, -1000000000, -1000000000, -1000000000, 0},
			target: 0,
			want:   [][]int{{-1000000000, -1000000000, 1000000000, 1000000000}},
		},
		{
			name:   "many duplicates single answer",
			nums:   []int{1, 1, 1, 1, 1, -3, -3, -3, 2, 2, 2},
			target: 1,
			want:   [][]int{{-3, 1, 1, 2}},
		},
		{
			name:   "no answer with many numbers",
			nums:   []int{-8, -6, -4, -2, 2, 4, 6, 8},
			target: 1,
			want:   [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fourSum(tt.nums, tt.target)
			assertEqualQuads(t, got, tt.want, tt.target)
		})
	}
}
