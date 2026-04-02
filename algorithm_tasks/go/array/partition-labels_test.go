// https://leetcode.com/problems/partition-labels

package array

import (
	"reflect"
	"testing"
)

func partitionLabels(s string) []int {
	lastIndexes := make(map[byte]int)
	sBytes := []byte(s)
	for i, b := range sBytes {
		lastIndexes[b] = i
	}
	res := make([]int, 0, len(sBytes)/4)
	size, end := 0, 0
	for i, b := range sBytes {
		if end < lastIndexes[b] {
			end = lastIndexes[b]
		}
		size++
		if end == i {
			res = append(res, size)
			size = 0
		}
	}
	return res
}

func TestPartitionLabels(t *testing.T) {
	tests := []struct {
		s    string
		want []int
	}{
		{"ababcbacadefegdehijhklij", []int{9, 7, 8}},
		{"a", []int{1}},
		{"z", []int{1}},
		{"ab", []int{1, 1}},
		{"aaaa", []int{4}},
		{"abcd", []int{1, 1, 1, 1}},
		{"abca", []int{4}},
		{"aab", []int{2, 1}},
		{"aaabbbccc", []int{3, 3, 3}},
		{"abcabc", []int{6}},
		{"caedbdedda", []int{1, 9}},
		{"eccbbbbdec", []int{10}},
	}
	for _, tt := range tests {
		got := partitionLabels(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("partitionLabels(%q) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
