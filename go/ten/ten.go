package ten

import (
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
