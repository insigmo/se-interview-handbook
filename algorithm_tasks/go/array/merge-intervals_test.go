// https://leetcode.com/problems/merge-intervals/

package array

import (
	"reflect"
	"sort"
	"testing"
)

func merge(intervals [][]int) [][]int {
	var ans [][]int

	// Шаг 1: сортируем по левой границе интервала
	// [1,3],[8,10],[2,6] → [1,3],[2,6],[8,10]
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for _, interval := range intervals {
		n := len(ans)

		// Случай 1: ans пуст ИЛИ текущий интервал не пересекается с последним
		// ans = [[1,3]], interval = [8,10]
		// 8 > 3 → не пересекаются → просто добавляем
		if n == 0 || interval[0] > ans[n-1][1] {
			ans = append(ans, interval)

			// Случай 2: интервалы пересекаются И текущий правее
			// ans = [[1,3]], interval = [2,6]
			// 2 <= 3 → пересечение; 6 > 3 → расширяем правую границу: [1,6]
		} else if interval[1] > ans[n-1][1] {
			ans[n-1][1] = interval[1]
		}
		// Случай 3 (неявный): текущий полностью внутри последнего
		// ans = [[1,6]], interval = [2,4]
		// 2 <= 6 AND 4 <= 6 → ничего не делаем, пропускаем
	}

	return ans
}

func TestMerge(t *testing.T) {
	normalize := func(intervals [][]int) [][]int {
		cp := make([][]int, len(intervals))
		for i := range intervals {
			if len(intervals[i]) != 2 {
				t.Fatalf("each interval must have length 2, got %v", intervals[i])
			}
			a, b := intervals[i][0], intervals[i][1]
			if a > b {
				t.Fatalf("interval start must be <= end, got %v", intervals[i])
			}
			cp[i] = []int{a, b}
		}

		sort.Slice(cp, func(i, j int) bool {
			if cp[i][0] == cp[j][0] {
				return cp[i][1] < cp[j][1]
			}
			return cp[i][0] < cp[j][0]
		})
		return cp
	}

	assertMerged := func(t *testing.T, got, want [][]int) {
		t.Helper()

		gotNorm := normalize(got)
		wantNorm := normalize(want)

		for i := 1; i < len(gotNorm); i++ {
			prev, curr := gotNorm[i-1], gotNorm[i]
			if prev[1] >= curr[0] {
				t.Fatalf("intervals must be non-overlapping and merged at boundaries, got %v", gotNorm)
			}
		}

		if !reflect.DeepEqual(gotNorm, wantNorm) {
			t.Fatalf("got %v, want %v", gotNorm, wantNorm)
		}
	}

	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			name:      "example one",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "example two touching endpoints",
			intervals: [][]int{{1, 4}, {4, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			name:      "example three unsorted touching endpoints",
			intervals: [][]int{{4, 7}, {1, 4}},
			want:      [][]int{{1, 7}},
		},
		{
			name:      "single interval",
			intervals: [][]int{{1, 3}},
			want:      [][]int{{1, 3}},
		},
		{
			name:      "already disjoint no merge",
			intervals: [][]int{{1, 2}, {3, 4}, {6, 7}},
			want:      [][]int{{1, 2}, {3, 4}, {6, 7}},
		},
		{
			name:      "fully nested intervals",
			intervals: [][]int{{1, 10}, {2, 3}, {4, 8}},
			want:      [][]int{{1, 10}},
		},
		{
			name:      "duplicates collapse into one",
			intervals: [][]int{{1, 4}, {1, 4}, {1, 4}},
			want:      [][]int{{1, 4}},
		},
		{
			name:      "chain of touching intervals merged",
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}},
			want:      [][]int{{1, 4}},
		},
		{
			name:      "mixed overlapping and disjoint",
			intervals: [][]int{{5, 7}, {1, 3}, {2, 6}, {8, 10}, {9, 18}},
			want:      [][]int{{1, 7}, {8, 18}},
		},
		{
			name:      "zero length intervals",
			intervals: [][]int{{1, 1}, {1, 2}, {3, 3}},
			want:      [][]int{{1, 2}, {3, 3}},
		},
		{
			name:      "max bound values with overlaps",
			intervals: [][]int{{0, 0}, {0, 10000}, {5000, 10000}},
			want:      [][]int{{0, 10000}},
		},
		{
			name:      "intervals sharing same start different ends",
			intervals: [][]int{{1, 4}, {1, 5}, {1, 3}},
			want:      [][]int{{1, 5}},
		},
		{
			name:      "intervals sharing same end different starts",
			intervals: [][]int{{1, 5}, {2, 5}, {3, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			name:      "intervals sharing same end different starts",
			intervals: [][]int{{1, 4}, {0, 5}},
			want:      [][]int{{0, 5}},
		},
		{
			name:      "complex mixture of overlaps and gaps",
			intervals: [][]int{{1, 4}, {6, 9}, {2, 5}, {10, 13}, {12, 15}, {20, 21}, {18, 19}},
			want:      [][]int{{1, 5}, {6, 9}, {10, 15}, {18, 19}, {20, 21}},
		},
		{
			name:      "single point then overlapping range",
			intervals: [][]int{{2, 2}, {2, 3}, {4, 4}},
			want:      [][]int{{2, 3}, {4, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.intervals)
			assertMerged(t, got, tt.want)
		})
	}
}
