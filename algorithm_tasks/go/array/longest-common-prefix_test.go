package array

import "testing"

func longestCommonPrefix(strs []string) string {
	var res []uint8
	first := strs[0]
	for i := range len(first) {
		for _, word := range strs {
			if i == len(word) || first[i] != word[i] {
				return string(res)
			}
		}
		res = append(res, first[i])
	}
	return string(res)
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{
			name: "example 1",
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "example 2 no common prefix",
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			name: "single string",
			strs: []string{"alone"},
			want: "alone",
		},
		{
			name: "single empty string",
			strs: []string{""},
			want: "",
		},
		{
			name: "contains empty string",
			strs: []string{"abc", "", "ab"},
			want: "",
		},
		{
			name: "all empty strings",
			strs: []string{"", "", ""},
			want: "",
		},
		{
			name: "all identical strings",
			strs: []string{"same", "same", "same"},
			want: "same",
		},
		{
			name: "prefix is entire shortest string",
			strs: []string{"ab", "abc", "abcd"},
			want: "ab",
		},
		{
			name: "common prefix length one",
			strs: []string{"apple", "ape", "april"},
			want: "ap",
		},
		{
			name: "common prefix single character",
			strs: []string{"car", "cat", "cap"},
			want: "ca",
		},
		{
			name: "no common prefix from first character",
			strs: []string{"a", "b", "c"},
			want: "",
		},
		{
			name: "two strings one prefix of other",
			strs: []string{"test", "testing"},
			want: "test",
		},
		{
			name: "two strings no prefix",
			strs: []string{"ab", "cd"},
			want: "",
		},
		{
			name: "many strings same starting char only",
			strs: []string{"zebra", "zip", "zone", "z"},
			want: "z",
		},
		{
			name: "mismatch after long prefix",
			strs: []string{"interspecies", "interstellar", "interstate"},
			want: "inters",
		},
		{
			name: "mismatch at second char",
			strs: []string{"aa", "ab", "ac"},
			want: "a",
		},
		{
			name: "short string breaks prefix",
			strs: []string{"aaa", "aa", "aaaab"},
			want: "aa",
		},
		{
			name: "duplicate strings and one shorter",
			strs: []string{"prefix", "prefix", "pre"},
			want: "pre",
		},
		{
			name: "all one character equal",
			strs: []string{"x", "x", "x", "x"},
			want: "x",
		},
		{
			name: "all one character different",
			strs: []string{"x", "y", "z"},
			want: "",
		},
		{
			name: "common prefix with repeated letters",
			strs: []string{"aaaaab", "aaaaac", "aaaaa"},
			want: "aaaaa",
		},
		{
			name: "late mismatch with two strings",
			strs: []string{"abcdefghijkl", "abcdefghijkz"},
			want: "abcdefghijk",
		},
		{
			name: "empty first string",
			strs: []string{"", "abc", "abcd"},
			want: "",
		},
		{
			name: "prefix disappears because of last string",
			strs: []string{"flower", "flow", "flour", "x"},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonPrefix(tt.strs)
			if got != tt.want {
				t.Fatalf("longestCommonPrefix(%v) = %q, want %q", tt.strs, got, tt.want)
			}
		})
	}
}
