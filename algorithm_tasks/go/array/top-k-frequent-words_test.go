// https://leetcode.com/problems/top-k-frequent-words

package array

import (
	"sort"
	"testing"
)

type wordStruct struct {
	word string
	freq int
}

func topKFrequent(words []string, k int) []string {
	freq := make(map[string]int, len(words))
	for _, word := range words {
		freq[word]++
	}

	var uniqueWords []wordStruct
	for word, count := range freq {
		uniqueWords = append(uniqueWords, wordStruct{word, count})
	}

	sort.Slice(uniqueWords, func(i, j int) bool {
		if uniqueWords[i].freq == uniqueWords[j].freq {
			return uniqueWords[i].word < uniqueWords[j].word
		}
		return uniqueWords[i].freq > uniqueWords[j].freq
	})

	result := make([]string, k)
	for i := 0; i < k; i++ {
		result[i] = uniqueWords[i].word
	}

	return result
}

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		k     int
		want  []string
	}{
		{
			name:  "example one",
			words: []string{"i", "love", "leetcode", "i", "love", "coding"},
			k:     2,
			want:  []string{"i", "love"},
		},
		{
			name:  "example two",
			words: []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},
			k:     4,
			want:  []string{"the", "is", "sunny", "day"},
		},
		{
			name:  "k equals one returns most frequent",
			words: []string{"a", "b", "b", "c", "c", "c"},
			k:     1,
			want:  []string{"c"},
		},
		{
			name:  "k equals one tie broken lexicographically",
			words: []string{"a", "b"},
			k:     1,
			want:  []string{"a"},
		},
		{
			name:  "all same word",
			words: []string{"go", "go", "go", "go"},
			k:     1,
			want:  []string{"go"},
		},
		{
			name:  "single word k equals one",
			words: []string{"hello"},
			k:     1,
			want:  []string{"hello"},
		},
		{
			name:  "k equals total unique words",
			words: []string{"x", "y", "z", "x", "y", "x"},
			k:     3,
			want:  []string{"x", "y", "z"},
		},
		{
			name:  "prefix words tiebreak",
			words: []string{"ab", "abc", "ab", "abc", "abcd"},
			k:     2,
			want:  []string{"ab", "abc"},
		},
		{
			name:  "large k returns all words sorted",
			words: []string{"one", "two", "three", "two", "three", "three"},
			k:     3,
			want:  []string{"three", "two", "one"},
		},
		{
			name:  "large k returns all words sorted",
			words: []string{"i", "love", "leetcode", "i", "love", "coding"},
			k:     3,
			want:  []string{"i", "love", "coding"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.words, tt.k)

			if len(got) != len(tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}

			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Fatalf("position %d: got %q, want %q (full: got %v, want %v)", i, got[i], tt.want[i], got, tt.want)
				}
			}
		})
	}
}
