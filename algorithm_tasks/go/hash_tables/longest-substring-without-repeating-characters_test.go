// https://leetcode.com/problems/single-number/description/

package hash_tables

import "testing"

func lengthOfLongestSubstring(s string) int {
	var res, left int
	knownChars := make(map[rune]struct{})

	for right, ch := range s {
		for {
			if _, ok := knownChars[ch]; !ok {
				break
			}
			delete(knownChars, rune(s[left]))
			left++
		}
		knownChars[ch] = struct{}{}
		res = max(res, right-left+1)
	}
	return res
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"empty", "", 0},
		{"one char", "a", 1},
		{"all unique", "abcdef", 6},
		{"all same", "bbbbb", 1},
		{"example 1", "abcabcbb", 3},
		{"example 2", "pwwkew", 3},
		{"example 3", "dvdf", 3},
		{"repeating at start", "abba", 2},
		{"long unique tail", "tmmzuxt", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Fatalf("lengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
