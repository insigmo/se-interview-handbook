// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee

package greedy

import "testing"

// change to maxProfit
func maxProfitWithFee(prices []int, fee int) int {
	profit := 0
	hold := -prices[0]
	for _, price := range prices[1:] {
		hold = max(hold, profit-price)
		profit = max(profit, price+hold-fee)
	}

	return profit
}

func TestMaxProfitWithFee(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		fee    int
		want   int
	}{
		{
			name:   "example 1",
			prices: []int{1, 3, 2, 8, 4, 9},
			fee:    2,
			want:   8,
		},
		{
			name:   "example 2",
			prices: []int{1, 3, 7, 5, 10, 3},
			fee:    3,
			want:   6,
		},
		{
			name:   "fee eats all profit",
			prices: []int{1, 2},
			fee:    2,
			want:   0,
		},
		{
			name:   "fee exactly equals profit",
			prices: []int{1, 3},
			fee:    2,
			want:   0,
		},
		{
			name:   "single element",
			prices: []int{5},
			fee:    1,
			want:   0,
		},
		{
			name:   "continuously rising, fee makes some not worthwhile",
			prices: []int{1, 2, 3, 4, 5},
			fee:    2,
			want:   2,
		},
		{
			name:   "continuously falling prices",
			prices: []int{9, 7, 5, 3, 1},
			fee:    1,
			want:   0,
		},
		{
			name:   "all same prices",
			prices: []int{4, 4, 4, 4},
			fee:    1,
			want:   0,
		},
		{
			name:   "zero fee - same as stock ii",
			prices: []int{1, 5, 1, 5},
			fee:    0,
			want:   8,
		},
		{
			name:   "large fee only one transaction profitable",
			prices: []int{1, 5, 1, 10},
			fee:    4,
			want:   5,
		},
		{
			name:   "multiple small gains beaten by one big gain after fee",
			prices: []int{1, 3, 1, 3, 1, 10},
			fee:    2,
			want:   7,
		},
		{
			name:   "two elements no profit after fee",
			prices: []int{3, 4},
			fee:    2,
			want:   0,
		},
		{
			name:   "high volatility with fee",
			prices: []int{1, 4, 1, 4, 1, 4},
			fee:    1,
			want:   6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfitWithFee(tt.prices, tt.fee)
			if got != tt.want {
				t.Fatalf("maxProfit(%v, %d) = %d, want %d", tt.prices, tt.fee, got, tt.want)
			}
		})
	}
}
