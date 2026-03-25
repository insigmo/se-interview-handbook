package math

import "testing"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	inverted := 0
	target := x
	for target != 0 {
		inverted = inverted*10 + target%10
		target /= 10
	}
	return inverted == x
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{
			name: "example 1 positive palindrome",
			x:    121,
			want: true,
		},
		{
			name: "example 2 negative number",
			x:    -121,
			want: false,
		},
		{
			name: "example 3 trailing zero",
			x:    10,
			want: false,
		},
		{
			name: "zero",
			x:    0,
			want: true,
		},
		{
			name: "single digit positive",
			x:    7,
			want: true,
		},
		{
			name: "single digit negative",
			x:    -7,
			want: false,
		},
		{
			name: "two same digits",
			x:    11,
			want: true,
		},
		{
			name: "two different digits",
			x:    12,
			want: false,
		},
		{
			name: "even length palindrome",
			x:    1221,
			want: true,
		},
		{
			name: "odd length palindrome",
			x:    12321,
			want: true,
		},
		{
			name: "odd length non palindrome",
			x:    12341,
			want: false,
		},
		{
			name: "even length non palindrome",
			x:    1231,
			want: false,
		},
		{
			name: "ends with zero but not zero itself",
			x:    100,
			want: false,
		},
		{
			name: "internal zeros palindrome",
			x:    1001,
			want: true,
		},
		{
			name: "internal zeros non palindrome",
			x:    1002,
			want: false,
		},
		{
			name: "large palindrome near int32 limit",
			x:    2147447412,
			want: true,
		},
		{
			name: "max int32 not palindrome",
			x:    2147483647,
			want: false,
		},
		{
			name: "min int32 negative",
			x:    -2147483648,
			want: false,
		},
		{
			name: "palindrome with many repeated digits",
			x:    1111111111,
			want: true,
		},
		{
			name: "non palindrome with repeated prefix",
			x:    1112111111,
			want: false,
		},
		{
			name: "middle zero palindrome",
			x:    10501,
			want: true,
		},
		{
			name: "middle zero non palindrome",
			x:    10531,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.x)
			if got != tt.want {
				t.Fatalf("isPalindrome(%d) = %v, want %v", tt.x, got, tt.want)
			}
		})
	}
}
