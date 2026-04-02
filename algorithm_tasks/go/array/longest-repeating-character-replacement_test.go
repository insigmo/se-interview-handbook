// https://leetcode.com/problems/longest-repeating-character-replacement

package array

import (
	"testing"
)

func characterReplacement(s string, k int) int {
	counts := make([]int, 26)
	maxCount := 0
	left, right := 0, 0
	length := len(s)
	for right < length {
		idx := s[right] - 'A'
		counts[idx]++
		if counts[idx] > maxCount {
			maxCount = counts[idx]
		}
		currentSize := right - left + 1
		if currentSize-maxCount > k {
			counts[s[left]-'A']--
			left++
		}
		right++
	}
	return right - left
}

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "basic example from problem",
			s:    "AABABBA",
			k:    1,
			want: 4,
		},
		{
			name: "replace all to same with k=2",
			s:    "XYYX",
			k:    2,
			want: 4,
		},
		{
			name: "another basic example",
			s:    "AAABABB",
			k:    1,
			want: 5,
		},
		{
			name: "k=0 no replacements allowed",
			s:    "AABBA",
			k:    0,
			want: 2,
		},
		{
			name: "k=0 all different characters",
			s:    "ABCDE",
			k:    0,
			want: 1,
		},
		{
			name: "all same characters k=0",
			s:    "AAAA",
			k:    0,
			want: 4,
		},
		{
			name: "all same characters large k",
			s:    "AAAA",
			k:    3,
			want: 4,
		},
		{
			name: "single character string",
			s:    "A",
			k:    0,
			want: 1,
		},
		{
			name: "single character string with k>0",
			s:    "A",
			k:    5,
			want: 1,
		},
		{
			name: "k large enough to cover entire string",
			s:    "ABCD",
			k:    10,
			want: 4,
		},
		{
			name: "two alternating characters enough k",
			s:    "ABAB",
			k:    2,
			want: 4,
		},
		{
			name: "two blocks of two k=1",
			s:    "AABB",
			k:    1,
			want: 3,
		},
		{
			name: "best window is in the middle",
			s:    "BAAABB",
			k:    1,
			want: 4,
		},
		{
			name: "all different characters k covers all but one",
			s:    "ABCDE",
			k:    4,
			want: 5,
		},
		{
			name: "two characters k=0 pick the longer run",
			s:    "AAABBB",
			k:    0,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := characterReplacement(tt.s, tt.k)
			if got != tt.want {
				t.Fatalf("characterReplacement(%q, %d) = %d, want %d", tt.s, tt.k, got, tt.want)
			}
		})
	}
}
