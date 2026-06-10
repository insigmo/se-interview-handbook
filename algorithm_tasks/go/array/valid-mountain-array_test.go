// https://leetcode.com/problems/valid-mountain-array/

package array

import "testing"

func validMountainArray(arr []int) bool {
	length := len(arr)
	i := 1

	for i < length && arr[i] > arr[i-1] {
		i++
	}

	if i == 1 || i == length {
		return false
	}

	for i < length && arr[i] < arr[i-1] {
		i++
	}

	return i == length
}

func Test_validMountainArray(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{"too short", []int{1, 2}, false},
		{"example true", []int{0, 3, 2, 1}, true},
		{"only up", []int{1, 2, 3, 4, 5}, false},
		{"only down", []int{5, 4, 3, 2, 1}, false},
		{"flat peak", []int{0, 3, 3, 2, 1}, false},
		{"plateau start", []int{1, 2, 2, 3, 1}, false},
		{"plateau end", []int{1, 2, 3, 3, 2, 1}, false},
		{"valley", []int{2, 1, 4, 7, 3, 2, 5}, false},
		{"peak at edge right", []int{0, 1, 2, 3, 4}, false},
		{"valid longer", []int{2, 4, 6, 8, 7, 3, 1}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray(tt.arr); got != tt.want {
				t.Fatalf("validMountainArray(%v) = %v, want %v", tt.arr, got, tt.want)
			}
		})
	}
}
