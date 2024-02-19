package ten

import (
	"fmt"
	"sort"
	"strings"
)

// Problem: Check if the given slice contains duplicates.
// Approach:
//
//	-Iterate over the slice of ints
//	-Use a map to memorize previously seen integers
//	-Check if current integer is in memory
func ContainsDuplicates(nums []int) bool {
	memory := make(map[int]bool, 0)

	for _, num := range nums {
		if _, exists := memory[num]; exists {
			return true
		}
		memory[num] = true
	}
	return false
}

// Problem: Check if two strings are anagrams of each other.
// Approach:
//
//	-Sort both strings
//	-Check for equality
func ValidAnagrams(one, two string) bool {
	oneChars := strings.Split(one, "")
	twoChars := strings.Split(two, "")

	sort.Strings(oneChars)
	sort.Strings(twoChars)

	one = strings.Join(oneChars, "")
	two = strings.Join(twoChars, "")

	return one == two
}

// Problem: Return indices of two numbers that sum to target
// Approach:
//
//	-Create a mapping of numbers to locations
//	-Iterate over numbers, and calculate delta
//	-Check if delta exists within map
func TwoSum(nums []int, target int) []int {
	locations := make(map[int][]int)
	for idx, num := range nums {
		if _, exists := locations[num]; exists {
			locations[num] = append(locations[num], idx)
		} else {
			locations[num] = []int{idx}
		}
	}
	fmt.Println("{}", locations)
	for idx, num := range nums {
		delta := target - num
		if locs, exists := locations[delta]; exists {
			if len(locs) >= 2 {
				return []int{idx, locs[1]}
			}
			if locs[0] != idx {
				return []int{idx, locs[0]}
			}
		}
	}
	return []int{}
}
