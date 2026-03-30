// https://leetcode.com/problems/valid-anagram/description/

package array

import (
	"testing"
)

func isAnagram(s string, t string) bool {
	const (
		englishLettersCount = 26
		aLetter             = 'a'
	)
	sRunes := [englishLettersCount]rune{}
	tRunes := [englishLettersCount]rune{}

	for _, ch := range s {
		sRunes[ch-aLetter]++
	}

	for _, ch := range t {
		tRunes[ch-aLetter]++
	}
	return sRunes == tRunes
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "long text example false",
			s:    "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
			t:    "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbba",
			want: false,
		},
		{
			name: "example true",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "example false",
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			name: "single same char",
			s:    "a",
			t:    "a",
			want: true,
		},
		{
			name: "single different char",
			s:    "a",
			t:    "b",
			want: false,
		},
		{
			name: "different lengths",
			s:    "ab",
			t:    "a",
			want: false,
		},
		{
			name: "same letters different order",
			s:    "listen",
			t:    "silent",
			want: true,
		},
		{
			name: "same string",
			s:    "abcabc",
			t:    "abcabc",
			want: true,
		},
		{
			name: "same length different counts",
			s:    "aacc",
			t:    "ccac",
			want: false,
		},
		{
			name: "repeated letters true",
			s:    "zzzxy",
			t:    "yxzzz",
			want: true,
		},
		{
			name: "repeated letters false",
			s:    "zzzxy",
			t:    "yyzzz",
			want: false,
		},
		{
			name: "all same letters true",
			s:    "bbbbbb",
			t:    "bbbbbb",
			want: true,
		},
		{
			name: "all same letters false",
			s:    "bbbbbb",
			t:    "bbbbba",
			want: false,
		},
		{
			name: "alphabet permutation",
			s:    "abcdefghijklmnopqrstuvwxyz",
			t:    "zyxwvutsrqponmlkjihgfedcba",
			want: true,
		},
		{
			name: "last char mismatch",
			s:    "abcdefghijklmnopqrstuvwxyy",
			t:    "abcdefghijklmnopqrstuvwxyz",
			want: false,
		},
		{
			name: "many duplicates true",
			s:    "aaabbbcccdddeeefff",
			t:    "fedcbafedcbafedcba",
			want: true,
		},
		{
			name: "many duplicates false",
			s:    "aaabbbcccdddeeefff",
			t:    "fedcbafedcbafedcbb",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAnagram(tt.s, tt.t)
			if got != tt.want {
				t.Fatalf("isAnagram(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
