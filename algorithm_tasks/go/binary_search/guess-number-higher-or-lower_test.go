// https://leetcode.com/problems/guess-number-higher-or-lower/

package binary_search

import "testing"

var pickedNumber = make([]int, 1)

func guessNumber(n int) int {
	left := 1

	for left <= n {
		mid := (left + n) / 2
		switch guess(mid) {
		case 0:
			return mid
		case 1:
			left = mid + 1
		case -1:
			n = mid - 1
		}
	}
	return -1
}

func guess(num int) int {
	pick := pickedNumber[0]
	switch {
	case num > pick:
		return -1
	case num < pick:
		return 1
	default:
		return 0
	}
}

func TestGuessNumber(t *testing.T) {
	tests := []struct {
		name string
		n    int
		pick int
		want int
	}{
		// примеры из условия задачи
		{"example1_n10_pick6", 10, 6, 6},
		{"example2_n1_pick1", 1, 1, 1},
		{"example3_n2_pick1", 2, 1, 1},

		// граничные: минимальный диапазон
		{"min_n1", 1, 1, 1},

		// граничные: pick на краях диапазона
		{"pick_equals_n", 10, 10, 10},
		{"pick_equals_1", 10, 1, 1},

		// середина и соседние значения
		{"pick_mid", 100, 50, 50},
		{"pick_below_mid", 100, 49, 49},
		{"pick_above_mid", 100, 51, 51},

		// n = 2, оба варианта
		{"n2_pick1", 2, 1, 1},
		{"n2_pick2", 2, 2, 2},

		// большой диапазон
		{"large_n_pick_1", 1_000_000, 1, 1},
		{"large_n_pick_end", 1_000_000, 1_000_000, 1_000_000},
		{"large_n_pick_mid", 1_000_000, 499_999, 499_999},

		// максимальное n = 2^31 - 1
		{"max_n_pick_1", 1<<31 - 1, 1, 1},
		{"max_n_pick_max", 1<<31 - 1, 1<<31 - 1, 1<<31 - 1},
		{"max_n_pick_mid", 1<<31 - 1, 1 << 30, 1 << 30},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pickedNumber[0] = tc.pick
			got := guessNumber(tc.n)
			if got != tc.want {
				t.Errorf("guessNumber(%d) pick=%d = %d; want %d",
					tc.n, tc.pick, got, tc.want)
			}
		})
	}
}
