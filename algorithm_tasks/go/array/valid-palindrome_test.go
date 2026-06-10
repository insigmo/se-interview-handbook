// https://leetcode.com/problems/valid-palindrome/

package array

import (
	"testing"
)

const letterBigA = 'A'
const letterBigZ = 'Z'
const letterA = 'a'
const letterZ = 'z'
const letter0 = '0'
const letter9 = '9'

func isLetterNum(r rune) bool {
	if r >= letter0 && r <= letter9 {
		return true
	}
	if r >= letterBigA && r <= letterBigZ {
		return true
	}
	if r >= letterA && r <= letterZ {
		return true
	}

	return false
}

// A = 65, D = 69
// D - A = 4
// D - A + a = d
func toLower(r rune) rune {
	if r >= letterBigA && r <= letterBigZ {
		return r - letterBigA + letterA
	}
	return r
}

func isPalindrome(s string) bool {
	length := len(s)
	if length <= 0 {
		return true
	}

	left := 0
	right := length - 1
	for left <= right {
		leftRune := toLower(rune(s[left]))
		rightRune := toLower(rune(s[right]))

		if !isLetterNum(leftRune) {
			left++
			continue
		}

		if !isLetterNum(rightRune) {
			right--
			continue
		}

		if leftRune != rightRune {
			return false
		}
		left++
		right--
	}

	return true
}

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		check string
		want  bool
	}{
		{
			name:  "check with unknown symbols",
			check: "A man, a plan, a canal: Panama",
			want:  true,
		},
		{
			name:  "check with odd symbols count",
			check: "raceacar",
			want:  false,
		},
		{
			name:  "check with number",
			check: "0f",
			want:  true,
		},
		{
			name:  "check random word",
			check: "superman",
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := isPalindrome(test.check)
			if got != test.want {
				t.Fatalf("isPalindrome(%s) = %v, want %v", test.check, got, test.want)
			}
		})
	}
}
