// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown

package greedy

import "testing"

// change to maxProfit
func maxProfitWithCooldown(prices []int) int {
	hold, cooldown := -prices[0], -prices[0]
	unHold := 0

	for _, price := range prices {
		hold = max(hold, unHold-price)
		unHold = max(unHold, cooldown)
		cooldown = hold + price
	}
	return max(cooldown, unHold)
}

func TestMaxProfitWithCooldown(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "example 1",
			prices: []int{1, 2, 3, 0, 2},
			want:   3,
		},
		{
			name:   "example 2 - single element",
			prices: []int{1},
			want:   0,
		},
		{
			name:   "two elements profit",
			prices: []int{1, 5},
			want:   4,
		},
		{
			name:   "two elements no profit",
			prices: []int{5, 1},
			want:   0,
		},
		{
			name:   "cooldown forces skip best buy",
			prices: []int{6, 1, 3, 2, 4, 7},
			want:   6,
		},
		{
			name:   "continuously rising",
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			name:   "continuously falling",
			prices: []int{5, 4, 3, 2, 1},
			want:   0,
		},
		{
			name:   "all same prices",
			prices: []int{3, 3, 3, 3},
			want:   0,
		},
		{
			name:   "sell then cooldown then buy again",
			prices: []int{1, 5, 1, 1, 5},
			want:   8,
		},
		{
			name:   "sell then cooldown blocks next buy",
			prices: []int{1, 5, 2, 1, 4},
			want:   7,
		},
		{
			name:   "three elements with cooldown",
			prices: []int{2, 1, 4},
			want:   3,
		},
		{
			name:   "alternating with cooldown penalty",
			prices: []int{1, 4, 1, 4, 1, 4},
			want:   6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfitWithCooldown(tt.prices)
			if got != tt.want {
				t.Fatalf("maxProfit(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}
