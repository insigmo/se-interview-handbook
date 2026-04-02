// https://leetcode.com/problems/best-time-to-buy-and-sell-stock

package greedy

import "testing"

func maxProfit(prices []int) int {
	res := 0
	lowest := prices[0]
	for _, price := range prices[1:] {
		profit := price - lowest
		if lowest > price {
			lowest = price
		}
		res = max(profit, res)
	}
	return res
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "example with clear buy low sell high window",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			name:   "strictly decreasing prices yields zero",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			name:   "single day cannot make transaction",
			prices: []int{5},
			want:   0,
		},
		{
			name:   "all equal prices yields zero",
			prices: []int{3, 3, 3, 3},
			want:   0,
		},
		{
			name:   "best profit is from first day to last day",
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			name:   "best sell happens before a later lower buy candidate is invalid",
			prices: []int{2, 4, 1},
			want:   2,
		},
		{
			name:   "minimum appears late but still allows small profit",
			prices: []int{9, 8, 1, 2},
			want:   1,
		},
		{
			name:   "multiple valleys and peaks choose best single transaction",
			prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			want:   4,
		},
		{
			name:   "zero price to max price",
			prices: []int{10000, 0, 10000},
			want:   10000,
		},
		{
			name:   "profit comes from middle segment not endpoints",
			prices: []int{8, 2, 6, 1, 7},
			want:   6,
		},
		{
			name:   "repeated minimum before best sell",
			prices: []int{5, 1, 1, 1, 6},
			want:   5,
		},
		{
			name:   "early spike is still optimal over later smaller gains",
			prices: []int{1, 10, 2, 3, 4},
			want:   9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.prices)
			if got != tt.want {
				t.Fatalf("maxProfit(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}
