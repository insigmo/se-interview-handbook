// https://leetcode.com/problems/group-anagrams

package array

import (
	"sort"
	"testing"
)

func groupAnagrams(strs []string) [][]string {
	const englishLettersCount = 26
	const aLetter byte = 'a'
	type letters [englishLettersCount]byte

	buf := make(map[letters][]string, len(strs)/4)

	for _, str := range strs {
		counter := letters{}
		for _, letter := range []byte(str) {
			counter[letter-aLetter]++
		}
		if _, ok := buf[counter]; !ok {
			buf[counter] = []string{}
		}
		buf[counter] = append(buf[counter], str)
	}
	result := make([][]string, 0, len(buf))
	for _, words := range buf {
		result = append(result, words)
	}
	return result
}

// with Sorting - O(m*n*log n) too big
func groupAnagramsWithSorting(strs []string) [][]string {
	buf := make(map[string][]string, len(strs)/4)
	for _, str := range strs {
		w := []byte(str)
		sort.Slice(w, func(i, j int) bool { return i < j })
		buf[string(w)] = append(buf[string(w)], str)
	}
	res := make([][]string, 0, len(buf))
	for _, words := range buf {
		res = append(res, words)
	}
	return res
}

func TestGroupAnagrams(t *testing.T) {
	type sig [26]int

	makeSig := func(s string) sig {
		var res sig
		for i := 0; i < len(s); i++ {
			res[s[i]-'a']++
		}
		return res
	}

	build := func(t *testing.T, groups [][]string) map[sig]map[string]int {
		t.Helper()

		res := make(map[sig]map[string]int)
		for _, group := range groups {
			if len(group) == 0 {
				t.Fatalf("group must not be empty: %v", groups)
			}

			key := makeSig(group[0])
			if _, exists := res[key]; exists {
				t.Fatalf("duplicate anagram group for signature: %v", group)
			}

			res[key] = make(map[string]int)
			for _, word := range group {
				if makeSig(word) != key {
					t.Fatalf("group contains non-anagrams: %v", group)
				}
				res[key][word]++
			}
		}

		return res
	}

	assertGroups := func(t *testing.T, got [][]string, want [][]string) {
		t.Helper()

		gotMap := build(t, got)
		wantMap := build(t, want)

		if len(gotMap) != len(wantMap) {
			t.Fatalf("got %v, want %v", got, want)
		}

		for key, wantWords := range wantMap {
			gotWords, ok := gotMap[key]
			if !ok {
				t.Fatalf("missing group for signature in got=%v want=%v", got, want)
			}

			if len(gotWords) != len(wantWords) {
				t.Fatalf("group mismatch for signature in got=%v want=%v", got, want)
			}

			for word, wantCount := range wantWords {
				if gotWords[word] != wantCount {
					t.Fatalf("word count mismatch for %q in got=%v want=%v", word, got, want)
				}
			}

			for word := range gotWords {
				if _, ok := wantWords[word]; !ok {
					t.Fatalf("unexpected word %q in got=%v want=%v", word, got, want)
				}
			}
		}
	}

	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "example one",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			name: "example two",
			strs: []string{""},
			want: [][]string{{""}},
		},
		{
			name: "example three",
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
		{
			name: "multiple empty strings",
			strs: []string{"", "", ""},
			want: [][]string{{"", "", ""}},
		},
		{
			name: "no anagrams at all",
			strs: []string{"ab", "cd", "ef", "gh"},
			want: [][]string{{"ab"}, {"cd"}, {"ef"}, {"gh"}},
		},
		{
			name: "all strings in one group",
			strs: []string{"abc", "bca", "cab", "cba", "bac", "acb"},
			want: [][]string{{"abc", "bca", "cab", "cba", "bac", "acb"}},
		},
		{
			name: "duplicate words preserved",
			strs: []string{"abc", "bca", "abc", "cab", "foo", "ofo", "foo"},
			want: [][]string{{"abc", "bca", "abc", "cab"}, {"foo", "ofo", "foo"}},
		},
		{
			name: "single character groups",
			strs: []string{"a", "b", "a", "c", "b"},
			want: [][]string{{"a", "a"}, {"b", "b"}, {"c"}},
		},
		{
			name: "same letters different counts are not anagrams",
			strs: []string{"a", "aa", "aaa", "aaaa"},
			want: [][]string{{"a"}, {"aa"}, {"aaa"}, {"aaaa"}},
		},
		{
			name: "mixed sizes and repeated groups",
			strs: []string{"ab", "ba", "abc", "cba", "bca", "z", "zz", "zz"},
			want: [][]string{{"ab", "ba"}, {"abc", "cba", "bca"}, {"z"}, {"zz", "zz"}},
		},
		{
			name: "words with repeated letters",
			strs: []string{"aabb", "bbaa", "abab", "abba", "baab", "baba", "abcd"},
			want: [][]string{{"aabb", "bbaa", "abab", "abba", "baab", "baba"}, {"abcd"}},
		},
		{
			name: "longer words",
			strs: []string{"listen", "silent", "enlist", "google", "gogole", "abc", "def"},
			want: [][]string{{"listen", "silent", "enlist"}, {"google", "gogole"}, {"abc"}, {"def"}},
		},
		{
			name: "anagrams with empty and non-empty",
			strs: []string{"", "b", "", "bb", "b"},
			want: [][]string{{"", ""}, {"b", "b"}, {"bb"}},
		},
		{
			name: "many isolated and grouped words",
			strs: []string{"rat", "tar", "art", "star", "tars", "cheese", "seeche", "hello"},
			want: [][]string{{"rat", "tar", "art"}, {"star", "tars"}, {"cheese", "seeche"}, {"hello"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.strs)
			assertGroups(t, got, tt.want)
		})
	}
}
