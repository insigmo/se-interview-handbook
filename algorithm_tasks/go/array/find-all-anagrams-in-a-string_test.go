// https://leetcode.com/problems/find-all-anagrams-in-a-string

package array

import (
	"reflect"
	"sort"
	"testing"
)

func findAnagrams(s string, p string) []int {
	const (
		englishLettersCount = 26
		aLetter             = 'a'
	)
	lengthS := len(s)
	lengthP := len(p)

	if lengthS < lengthP {
		return []int{}
	}

	var ans []int
	type letters [englishLettersCount]rune

	// if count has only zeros, then p has at least 1 anagram in s
	counter := letters{}
	for i := 0; i < lengthP; i++ {
		counter[p[i]-aLetter]++
		counter[s[i]-aLetter]--
	}

	// check that all elements is zero
	// if it's zero array, it's anagram
	isZero := func(c letters) bool {
		for _, v := range c {
			if v != 0 {
				return false
			}
		}
		return true
	}

	if isZero(counter) {
		ans = append(ans, 0)
	}

	for i := 1; i <= lengthS-lengthP; i++ {
		counter[s[i-1]-aLetter]++
		counter[s[i+lengthP-1]-aLetter]--

		if isZero(counter) {
			ans = append(ans, i)
		}
	}
	return ans
}

func TestFindAnagrams(t *testing.T) {
	assertEqual := func(t *testing.T, got, want []int) {
		t.Helper()

		gotCopy := append([]int(nil), got...)
		wantCopy := append([]int(nil), want...)

		sort.Ints(gotCopy)
		sort.Ints(wantCopy)

		if !reflect.DeepEqual(gotCopy, wantCopy) {
			t.Fatalf("got %v, want %v", got, want)
		}
	}

	tests := []struct {
		name string
		s    string
		p    string
		want []int
	}{
		{
			name: "example one",
			s:    "cbaebabacd",
			p:    "abc",
			want: []int{0, 6},
		},
		{
			name: "example two overlapping matches",
			s:    "abab",
			p:    "ab",
			want: []int{0, 1, 2},
		},
		{
			name: "single exact match",
			s:    "a",
			p:    "a",
			want: []int{0},
		},
		{
			name: "single no match",
			s:    "a",
			p:    "b",
			want: []int{},
		},
		{
			name: "p longer than s",
			s:    "ab",
			p:    "abc",
			want: []int{},
		},
		{
			name: "same length anagram",
			s:    "bca",
			p:    "abc",
			want: []int{0},
		},
		{
			name: "same length not anagram",
			s:    "abd",
			p:    "abc",
			want: []int{},
		},
		{
			name: "all same characters many overlaps",
			s:    "aaaaa",
			p:    "aa",
			want: []int{0, 1, 2, 3},
		},
		{
			name: "all same characters exact window",
			s:    "aaaa",
			p:    "aaaa",
			want: []int{0},
		},
		{
			name: "repeated letters in pattern",
			s:    "baa",
			p:    "aa",
			want: []int{1},
		},
		{
			name: "multiple repeated letter matches",
			s:    "aaabaaa",
			p:    "aaa",
			want: []int{0, 4},
		},
		{
			name: "matches at beginning middle end",
			s:    "abcxbacab",
			p:    "abc",
			want: []int{0, 4, 6},
		},
		{
			name: "no matches despite shared letters",
			s:    "abcdefg",
			p:    "hij",
			want: []int{},
		},
		{
			name: "window slides over extra repeated chars",
			s:    "abababab",
			p:    "aab",
			want: []int{0, 2, 4},
		},
		{
			name: "pattern with duplicate counts matters",
			s:    "abbcabc",
			p:    "abb",
			want: []int{0},
		},
		{
			name: "dense overlapping matches",
			s:    "aaaaaaaaaa",
			p:    "aaaa",
			want: []int{0, 1, 2, 3, 4, 5, 6},
		},
		{
			name: "alternating characters",
			s:    "babab",
			p:    "ab",
			want: []int{0, 1, 2, 3},
		},
		{
			name: "pattern appears only once after noise",
			s:    "zzzzabczz",
			p:    "bca",
			want: []int{4},
		},
		{
			name: "repeated windows with exact counts",
			s:    "abcabcabc",
			p:    "cba",
			want: []int{0, 1, 2, 3, 4, 5, 6},
		},
		{
			name: "count mismatch blocks near matches",
			s:    "aabaa",
			p:    "aaa",
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findAnagrams(tt.s, tt.p)
			assertEqual(t, got, tt.want)
		})
	}
}
