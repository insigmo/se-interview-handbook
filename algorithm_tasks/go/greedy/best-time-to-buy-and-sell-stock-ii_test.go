// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii

package greedy

import "testing"

// change to maxProfit
func maxProfitII(prices []int) int {
	profit := 0
	hold := -prices[0]
	for _, price := range prices[1:] {
		hold = max(hold, profit-price)
		profit = max(profit, price+hold)
	}

	return profit
}

func TestMaxProfitII(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "example 1 - multiple profitable transactions",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   7,
		},
		{
			name:   "example 2 - continuously rising prices",
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			name:   "example 3 - continuously falling prices",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			name:   "single element",
			prices: []int{5},
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
			name:   "all same prices",
			prices: []int{3, 3, 3, 3},
			want:   0,
		},
		{
			name:   "alternating prices",
			prices: []int{1, 5, 1, 5, 1, 5},
			want:   12,
		},
		{
			name:   "valley then peak",
			prices: []int{3, 1, 4, 1, 5},
			want:   7,
		},
		{
			name:   "large spike then drop",
			prices: []int{1, 100, 1, 100},
			want:   198,
		},
		{
			name:   "two elements equal",
			prices: []int{4, 4},
			want:   0,
		},
		{
			name:   "profit only at the end",
			prices: []int{5, 5, 5, 6},
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfitII(tt.prices)
			if got != tt.want {
				t.Fatalf("maxProfit(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}
