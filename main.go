package main

import "fmt"

// LongestConsecutive returns the length of the longest run of consecutive
// integers present in nums; order and duplicates don't matter.
func LongestConsecutive(nums []int) int {
	set := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		set[v] = struct{}{}
	}
	best := 0
	for v := range set {
		if _, ok := set[v-1]; ok {
			continue // v is inside a run, not its start; skip so each run is walked once
		}
		length := 1
		for {
			if _, ok := set[v+length]; !ok {
				break
			}
			length++
		}
		if length > best {
			best = length
		}
	}
	return best
}

func main() {
	fmt.Println(LongestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
