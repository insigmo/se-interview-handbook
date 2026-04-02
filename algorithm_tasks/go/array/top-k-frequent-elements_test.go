// https://leetcode.com/problems/top-k-frequent-elements

package array

import (
	"sort"
	"testing"
)

type numStruct struct {
	num  int
	freq int
}

// change to topKFrequent
func topKFrequentElement(nums []int, k int) []int {
	freq := make(map[int]int, len(nums))
	for _, num := range nums {
		freq[num]++
	}

	var uniqueWords []numStruct
	for word, count := range freq {
		uniqueWords = append(uniqueWords, numStruct{word, count})
	}

	sort.Slice(uniqueWords, func(i, j int) bool {
		return uniqueWords[i].freq > uniqueWords[j].freq
	})

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = uniqueWords[i].num
	}

	return result

}

func TestTopKFrequentElement(t *testing.T) {
	assertResult := func(t *testing.T, got []int, want []int) {
		t.Helper()

		if len(got) != len(want) {
			t.Fatalf("got %v (len=%d), want %v (len=%d)", got, len(got), want, len(want))
		}

		wantSet := make(map[int]struct{})
		for _, v := range want {
			wantSet[v] = struct{}{}
		}

		gotSet := make(map[int]struct{})
		for _, v := range got {
			gotSet[v] = struct{}{}
		}

		if len(gotSet) != len(got) {
			t.Fatalf("result contains duplicates: %v", got)
		}

		for v := range wantSet {
			if _, ok := gotSet[v]; !ok {
				t.Fatalf("missing element %d in %v", v, got)
			}
		}

		for v := range gotSet {
			if _, ok := wantSet[v]; !ok {
				t.Fatalf("unexpected element %d in %v", v, got)
			}
		}
	}

	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "example one",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "example two single element",
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "example three interleaved",
			nums: []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "k equals unique count returns all",
			nums: []int{3, 1, 2},
			k:    3,
			want: []int{1, 2, 3},
		},
		{
			name: "k equals one returns top",
			nums: []int{5, 5, 5, 3, 3, 1},
			k:    1,
			want: []int{5},
		},
		{
			name: "all same value",
			nums: []int{7, 7, 7, 7},
			k:    1,
			want: []int{7},
		},
		{
			name: "negative numbers",
			nums: []int{-1, -1, -1, -2, -2, -3},
			k:    2,
			want: []int{-1, -2},
		},
		{
			name: "mix of positive and negative",
			nums: []int{-1, 1, -1, 1, -1, 2},
			k:    2,
			want: []int{-1, 1},
		},
		{
			name: "zero is most frequent",
			nums: []int{0, 0, 0, 1, 2, 3},
			k:    1,
			want: []int{0},
		},
		{
			name: "zero included in top k",
			nums: []int{0, 0, 1, 1, 2},
			k:    2,
			want: []int{0, 1},
		},
		{
			name: "large gap between top two frequencies",
			nums: []int{4, 4, 4, 4, 4, 4, 9, 9, 1},
			k:    2,
			want: []int{4, 9},
		},
		{
			name: "top k from five distinct",
			nums: []int{1, 1, 2, 2, 3, 3, 3, 4, 5},
			k:    3,
			want: []int{1, 2, 3},
		},
		{
			name: "extreme negative values",
			nums: []int{-10000, -10000, -9999, -9999, -9999, 0},
			k:    2,
			want: []int{-10000, -9999},
		},
		{
			name: "extreme positive values",
			nums: []int{10000, 10000, 9999, 9999, 9999, 0},
			k:    2,
			want: []int{10000, 9999},
		},
		{
			name: "single pair of duplicates",
			nums: []int{3, 3},
			k:    1,
			want: []int{3},
		},
		{
			name: "k equals one with many candidates",
			nums: []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4},
			k:    1,
			want: []int{4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequentElement(tt.nums, tt.k)
			assertResult(t, got, tt.want)
		})
	}
}
