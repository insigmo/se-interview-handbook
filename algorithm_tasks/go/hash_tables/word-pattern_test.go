// https://leetcode.com/problems/word-pattern/

package hash_tables

import (
	"strings"
	"testing"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	length := len(words)
	if len(pattern) != length {
		return false
	}

	charToWord := make(map[uint8]string, length)
	wordToChar := make(map[string]uint8, length)

	for i := range length {
		p, w := pattern[i], words[i]
		if _, ok := charToWord[p]; ok && charToWord[p] != w {
			return false
		}
		if _, ok := wordToChar[w]; ok && wordToChar[w] != p {
			return false
		}
		charToWord[p] = w
		wordToChar[w] = p
	}
	return true
}

func TestWordPattern(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		s        string
		expected bool
	}{
		// --- примеры из условия ---
		{"example 1: abba / dog cat cat dog", "abba", "dog cat cat dog", true},
		{"example 2: abba / dog cat cat fish", "abba", "dog cat cat fish", false},
		{"example 3: aaaa / dog cat cat dog", "aaaa", "dog cat cat dog", false},
		{"example 4: abba / dog dog dog dog", "abba", "dog dog dog dog", false},

		// --- один символ / одно слово ---
		{"одно слово совпадает", "a", "dog", true},
		{"один символ, одно слово", "z", "hello", true},

		// --- разная длина ---
		{"паттерн длиннее", "aab", "dog cat", false},
		{"слов больше чем символов", "ab", "dog cat cat", false},

		// --- одинаковые слова для разных символов (не биекция) ---
		{"два разных символа → одно слово", "ab", "dog dog", false},

		// --- все символы одинаковые ---
		{"aaa / dog dog dog", "aaa", "dog dog dog", true},
		{"aaa / dog dog cat", "aaa", "dog dog cat", false},

		// --- полный паттерн, все разные ---
		{"abcd / dog cat fish bird", "abcd", "dog cat fish bird", true},
		{"abcd / dog cat fish dog", "abcd", "dog cat fish dog", false},

		// --- симметричный паттерн ---
		{"abba / dog cat cat dog", "abba", "dog cat cat dog", true},
		{"abab / dog cat dog cat", "abab", "dog cat dog cat", true},
		{"abab / dog cat cat dog", "abab", "dog cat cat dog", false},

		// --- одно слово повторяется, символ разный ---
		{"aa / cat cat", "aa", "cat cat", true},
		{"ab / cat cat", "ab", "cat cat", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordPattern(tt.pattern, tt.s)
			if got != tt.expected {
				t.Errorf("wordPattern(%q, %q) = %v, want %v",
					tt.pattern, tt.s, got, tt.expected)
			}
		})
	}
}
