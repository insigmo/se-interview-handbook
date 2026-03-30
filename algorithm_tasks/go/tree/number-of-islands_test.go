// https://leetcode.com/problems/number-of-islands/

package tree

import "testing"

type OceanMap struct {
	grid [][]byte
	n, m int
}

func (oceanMap *OceanMap) dfsNumIslands(i, j int) {
	if i < 0 || i >= oceanMap.n || j < 0 || j >= oceanMap.m || oceanMap.grid[i][j] != '1' {
		return
	}

	oceanMap.grid[i][j] = 0
	oceanMap.dfsNumIslands(i, j+1)
	oceanMap.dfsNumIslands(i+1, j)
	oceanMap.dfsNumIslands(i, j-1)
	oceanMap.dfsNumIslands(i-1, j)
}

func numIslands(grid [][]byte) int {
	islandsAmount := 0
	n, m := len(grid), len(grid[0])
	o := &OceanMap{grid: grid, n: n, m: m}
	for i := range n {
		for j := range m {
			if o.grid[i][j] == '1' {
				islandsAmount++
				o.dfsNumIslands(i, j)
			}
		}
	}
	return islandsAmount
}

func TestNumIslands(t *testing.T) {
	copyGrid := func(grid [][]byte) [][]byte {
		res := make([][]byte, len(grid))
		for i := range grid {
			res[i] = make([]byte, len(grid[i]))
			for j := range grid[i] {
				res[i][j] = grid[i][j]
			}
		}
		return res
	}

	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "example one",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			name: "example two",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
		{
			name: "single land cell",
			grid: [][]byte{
				{'1'},
			},
			want: 1,
		},
		{
			name: "single water cell",
			grid: [][]byte{
				{'0'},
			},
			want: 0,
		},
		{
			name: "all water",
			grid: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			want: 0,
		},
		{
			name: "all land",
			grid: [][]byte{
				{'1', '1', '1'},
				{'1', '1', '1'},
				{'1', '1', '1'},
			},
			want: 1,
		},
		{
			name: "diagonal cells are not connected",
			grid: [][]byte{
				{'1', '0', '1'},
				{'0', '1', '0'},
				{'1', '0', '1'},
			},
			want: 5,
		},
		{
			name: "single row alternating",
			grid: [][]byte{
				{'1', '0', '1', '0', '1', '0', '1'},
			},
			want: 4,
		},
		{
			name: "single column alternating",
			grid: [][]byte{
				{'1'},
				{'0'},
				{'1'},
				{'0'},
				{'1'},
			},
			want: 3,
		},
		{
			name: "narrow horizontal bridge makes one island",
			grid: [][]byte{
				{'1', '1', '0', '0'},
				{'0', '1', '1', '1'},
				{'0', '0', '0', '1'},
			},
			want: 1,
		},
		{
			name: "center island surrounded by water",
			grid: [][]byte{
				{'0', '0', '0', '0', '0'},
				{'0', '1', '1', '1', '0'},
				{'0', '1', '0', '1', '0'},
				{'0', '1', '1', '1', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			name: "four corner islands",
			grid: [][]byte{
				{'1', '0', '0', '1'},
				{'0', '0', '0', '0'},
				{'0', '0', '0', '0'},
				{'1', '0', '0', '1'},
			},
			want: 4,
		},
		{
			name: "checkerboard 4 by 4",
			grid: [][]byte{
				{'1', '0', '1', '0'},
				{'0', '1', '0', '1'},
				{'1', '0', '1', '0'},
				{'0', '1', '0', '1'},
			},
			want: 8,
		},
		{
			name: "vertical strips",
			grid: [][]byte{
				{'1', '0', '1', '0', '1'},
				{'1', '0', '1', '0', '1'},
				{'1', '0', '1', '0', '1'},
			},
			want: 3,
		},
		{
			name: "touching only by corners still separate",
			grid: [][]byte{
				{'1', '0', '0'},
				{'0', '1', '1'},
				{'0', '1', '0'},
			},
			want: 2,
		},
		{
			name: "snake shaped one island",
			grid: [][]byte{
				{'1', '1', '1', '0'},
				{'0', '0', '1', '0'},
				{'1', '1', '1', '0'},
				{'1', '0', '0', '0'},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numIslands(copyGrid(tt.grid))
			if got != tt.want {
				t.Fatalf("numIslands() = %d, want %d", got, tt.want)
			}
		})
	}
}
