package main

import "testing"

func TestRescue(t *testing.T) {
	tests := []struct {
		n         int
		k         int
		positions []int
		want      int
	}{
		{5, 5, []int{2, 5, 10, 12, 15}, 2},
		{6, 10, []int{1, 11, 30, 34, 35, 37}, 4},
		{10, 3, []int{1, 20, 33, 41, 42, 60, 70, 80, 90, 100}, 2},
	}

	for _, tt := range tests {
		got := rescue(tt.n, tt.k, tt.positions)
		if got != tt.want {
			t.Errorf("rescue(%v, %v, %v) = %v; want %v", tt.n, tt.k, tt.positions, got, tt.want)
		}
	}
}
