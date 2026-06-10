// https://leetcode.com/problems/plus-one/

package array

import (
	"reflect"
	"testing"
)

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 != 10 {
			digits[i]++
			return digits
		}
		digits[i] = 0

		if i == 0 {
			return append([]int{1}, digits...)
		}

	}
	return digits
}

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		// Базовые случаи
		{
			name:  "обычный случай — нет переноса",
			input: []int{1, 2, 3},
			want:  []int{1, 2, 4},
		},
		{
			name:  "последняя цифра 0",
			input: []int{1, 2, 0},
			want:  []int{1, 2, 1},
		},
		// Краевые случаи с переносом
		{
			name:  "одна девятка — расширение массива",
			input: []int{9},
			want:  []int{1, 0},
		},
		{
			name:  "все девятки — расширение массива",
			input: []int{9, 9, 9},
			want:  []int{1, 0, 0, 0},
		},
		{
			name:  "перенос только с последней цифры",
			input: []int{1, 9},
			want:  []int{2, 0},
		},
		{
			name:  "перенос через несколько цифр",
			input: []int{1, 9, 9},
			want:  []int{2, 0, 0},
		},
		// Граничные значения по условию задачи
		{
			name:  "массив из одного нуля",
			input: []int{0},
			want:  []int{1},
		},
		{
			name:  "массив из одной единицы",
			input: []int{1},
			want:  []int{2},
		},
		{
			name:  "длинный массив без переноса",
			input: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:  []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		},
		{
			name:  "длинный массив — все девятки",
			input: []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			want:  []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Копируем input, т.к. функция модифицирует слайс in-place
			input := make([]int, len(tt.input))
			copy(input, tt.input)

			got := plusOne(input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
