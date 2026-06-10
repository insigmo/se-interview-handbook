// https://leetcode.com/problems/string-compression/
package linked_list

import (
	"strconv"
	"testing"
)

func compress(chars []byte) int {
	current, ind, counter, length := 0, 0, 0, len(chars)

	for ind < length {
		counter = 0
		ch := chars[ind]
		for ind < length && chars[ind] == ch {
			ind++
			counter++
		}
		chars[current] = ch
		current++

		if counter > 1 {
			for _, b := range []byte(strconv.Itoa(counter)) {
				chars[current] = b
				current++
			}
		}
	}
	return current
}

func TestCompress(t *testing.T) {
	testCases := []struct {
		name      string
		input     []byte
		wantLen   int
		wantChars []byte
	}{
		{
			name:      "пример 1: aabbccc -> a2b2c3",
			input:     []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			wantLen:   6,
			wantChars: []byte{'a', '2', 'b', '2', 'c', '3'},
		},
		{
			name:      "пример 2: один символ -> без числа",
			input:     []byte{'a'},
			wantLen:   1,
			wantChars: []byte{'a'},
		},
		{
			name:      "пример 3: 12 одинаковых -> двузначное число",
			input:     []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
			wantLen:   3,
			wantChars: []byte{'a', '1', '2'},
		},
		{
			name:      "все уникальные символы -> без изменений",
			input:     []byte{'a', 'b', 'c'},
			wantLen:   3,
			wantChars: []byte{'a', 'b', 'c'},
		},
		{
			name: "один символ повторяется 100 раз",
			input: func() []byte {
				b := make([]byte, 100)
				for i := range b {
					b[i] = 'z'
				}
				return b
			}(),
			wantLen:   4,
			wantChars: []byte{'z', '1', '0', '0'},
		},
		{
			name:      "чередующиеся группы: aabba",
			input:     []byte{'a', 'a', 'b', 'b', 'a'},
			wantLen:   5,
			wantChars: []byte{'a', '2', 'b', '2', 'a'},
		},
		{
			name:      "два символа подряд: aa -> a2",
			input:     []byte{'a', 'a'},
			wantLen:   2,
			wantChars: []byte{'a', '2'},
		},
		{
			name:      "группа ровно 10: a×10 -> a10",
			input:     []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
			wantLen:   3,
			wantChars: []byte{'a', '1', '0'},
		},
		{
			name:      "группа 9: a×9 -> a9 (однозначное число)",
			input:     []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
			wantLen:   2,
			wantChars: []byte{'a', '9'},
		},
		{
			name:      "смешанные группы разной длины",
			input:     []byte{'a', 'b', 'b', 'b', 'c', 'c'},
			wantLen:   5,
			wantChars: []byte{'a', 'b', '3', 'c', '2'},
		},
		{
			name: "группа 5000: граничный случай по constraints",
			input: func() []byte {
				b := make([]byte, 5000)
				for i := range b {
					b[i] = 'x'
				}
				return b
			}(),
			wantLen:   5,
			wantChars: []byte{'x', '5', '0', '0', '0'},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			gotLen := compress(tc.input)
			if gotLen != tc.wantLen {
				t.Errorf("\n\tgot:\t%d\n\twant:\t%d", gotLen, tc.wantLen)
			}
		})
	}
}
