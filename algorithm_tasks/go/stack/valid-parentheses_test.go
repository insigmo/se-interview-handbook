// https://leetcode.com/problems/valid-parentheses

package stack

import "testing"

func isValid(s string) bool {
	openCloseParentphrases := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}
	stack := make([]byte, 0)

	for _, char := range []byte(s) {
		stackLength := len(stack)
		if v, ok := openCloseParentphrases[char]; ok {
			if !(stackLength != 0 && stack[stackLength-1] == v) {
				return false
			}
			stack = stack[:stackLength-1]
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "example simple pair",
			s:    "()",
			want: true,
		},
		{
			name: "example multiple pairs",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "example mismatched types",
			s:    "(]",
			want: false,
		},
		{
			name: "example nested valid",
			s:    "([])",
			want: true,
		},
		{
			name: "example wrong order",
			s:    "([)]",
			want: false,
		},
		{
			name: "single opening bracket",
			s:    "(",
			want: false,
		},
		{
			name: "single closing bracket",
			s:    "]",
			want: false,
		},
		{
			name: "two opening brackets",
			s:    "((",
			want: false,
		},
		{
			name: "two closing brackets",
			s:    "))",
			want: false,
		},
		{
			name: "simple curly braces",
			s:    "{}",
			want: true,
		},
		{
			name: "simple square brackets",
			s:    "[]",
			want: true,
		},
		{
			name: "nested different types",
			s:    "{[]}",
			want: true,
		},
		{
			name: "deeply nested valid",
			s:    "({[]})",
			want: true,
		},
		{
			name: "deeply nested invalid closing order",
			s:    "({[})",
			want: false,
		},
		{
			name: "valid sequential and nested mix",
			s:    "([]){}[()]",
			want: true,
		},
		{
			name: "starts with closing bracket",
			s:    ")([])",
			want: false,
		},
		{
			name: "extra opening at end",
			s:    "()[]{}{",
			want: false,
		},
		{
			name: "extra closing at end",
			s:    "()[]{}]",
			want: false,
		},
		{
			name: "repeated same type valid",
			s:    "()()()()",
			want: true,
		},
		{
			name: "repeated same type invalid",
			s:    "())(()",
			want: false,
		},
		{
			name: "complex valid",
			s:    "(([]){})",
			want: true,
		},
		{
			name: "complex invalid leftover opening",
			s:    "(([]){})(",
			want: false,
		},
		{
			name: "complex invalid mismatch near end",
			s:    "(([]){]}",
			want: false,
		},
		{
			name: "adjacent mismatches",
			s:    "][",
			want: false,
		},
		{
			name: "long balanced",
			s:    "(((((((((())))))))))",
			want: true,
		},
		{
			name: "long unbalanced missing one close",
			s:    "(((((((((()))))))))",
			want: false,
		},
		{
			name: "alternating nested valid",
			s:    "[{()}([])]",
			want: true,
		},
		{
			name: "alternating nested invalid",
			s:    "[{()}(][)]",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.s)
			if got != tt.want {
				t.Fatalf("isValid(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
