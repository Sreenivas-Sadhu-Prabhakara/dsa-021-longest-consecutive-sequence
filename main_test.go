package main

import "testing"

func TestLongestConsecutive(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{}, 0},
		{[]int{7}, 1},
		{[]int{1, 1, 1}, 1},
		{[]int{-2, -1, 0, 5}, 3},
	}
	for _, c := range cases {
		if got := LongestConsecutive(c.in); got != c.want {
			t.Errorf("LongestConsecutive(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
