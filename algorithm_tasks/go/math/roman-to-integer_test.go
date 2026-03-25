package math

import "testing"

var romans = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) int {
	target := 0
	nums := []byte(s)

	counter := 0
	next := 0

	for {
		if counter > len(nums)-1 {
			break
		}
		if counter == len(nums)-1 {
			next = 0
		} else {
			next = romans[nums[counter+1]]
		}
		current := romans[nums[counter]]

		if current < next {
			target += next - current
			counter += 2
		} else {
			counter++
			target += current
		}
	}
	return target
}

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example 1",
			s:    "III",
			want: 3,
		},
		{
			name: "example 2",
			s:    "LVIII",
			want: 58,
		},
		{
			name: "example 3",
			s:    "MCMXCIV",
			want: 1994,
		},
		{
			name: "minimum value",
			s:    "I",
			want: 1,
		},
		{
			name: "maximum value",
			s:    "MMMCMXCIX",
			want: 3999,
		},
		{
			name: "simple additive five",
			s:    "V",
			want: 5,
		},
		{
			name: "simple additive ten",
			s:    "X",
			want: 10,
		},
		{
			name: "simple additive fifty",
			s:    "L",
			want: 50,
		},
		{
			name: "simple additive hundred",
			s:    "C",
			want: 100,
		},
		{
			name: "simple additive five hundred",
			s:    "D",
			want: 500,
		},
		{
			name: "simple additive thousand",
			s:    "M",
			want: 1000,
		},
		{
			name: "subtractive four",
			s:    "IV",
			want: 4,
		},
		{
			name: "subtractive nine",
			s:    "IX",
			want: 9,
		},
		{
			name: "subtractive forty",
			s:    "XL",
			want: 40,
		},
		{
			name: "subtractive ninety",
			s:    "XC",
			want: 90,
		},
		{
			name: "subtractive four hundred",
			s:    "CD",
			want: 400,
		},
		{
			name: "subtractive nine hundred",
			s:    "CM",
			want: 900,
		},
		{
			name: "repeated ones",
			s:    "III",
			want: 3,
		},
		{
			name: "repeated tens",
			s:    "XXX",
			want: 30,
		},
		{
			name: "repeated hundreds",
			s:    "CCC",
			want: 300,
		},
		{
			name: "repeated thousands",
			s:    "MMM",
			want: 3000,
		},
		{
			name: "mixed additive and subtractive 14",
			s:    "XIV",
			want: 14,
		},
		{
			name: "mixed additive and subtractive 19",
			s:    "XIX",
			want: 19,
		},
		{
			name: "mixed additive and subtractive 44",
			s:    "XLIV",
			want: 44,
		},
		{
			name: "mixed additive and subtractive 49",
			s:    "XLIX",
			want: 49,
		},
		{
			name: "mixed additive and subtractive 99",
			s:    "XCIX",
			want: 99,
		},
		{
			name: "mixed additive and subtractive 444",
			s:    "CDXLIV",
			want: 444,
		},
		{
			name: "mixed additive and subtractive 944",
			s:    "CMXLIV",
			want: 944,
		},
		{
			name: "mixed additive and subtractive 1666",
			s:    "MDCLXVI",
			want: 1666,
		},
		{
			name: "complex 1987",
			s:    "MCMLXXXVII",
			want: 1987,
		},
		{
			name: "complex 2421",
			s:    "MMCDXXI",
			want: 2421,
		},
		{
			name: "complex 3888",
			s:    "MMMDCCCLXXXVIII",
			want: 3888,
		},
		{
			name: "multiple subtractive groups",
			s:    "CMXCIX",
			want: 999,
		},
		{
			name: "another common form 2024",
			s:    "MMXXIV",
			want: 2024,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := romanToInt(tt.s)
			if got != tt.want {
				t.Fatalf("romanToInt(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
