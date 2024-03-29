package ten

import (
	"fmt"
	"regexp"
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

// Problem: Return a list of grouped anagrams.
// Approach:
//
//	-Create a map: string -> [string]
//	-Sort the strings
//	-Check if sorted string exists in map
func values(anagrams map[string][]string) [][]string {
	vals := make([][]string, 0)
	for key := range anagrams {
		vals = append(vals, anagrams[key])
	}
	return vals
}
func splitAndSort(word string) string {
	chars := make([]string, 0)
	for _, char := range word {
		chars = append(chars, string(char))
	}
	sort.Strings(chars)
	return strings.Join(chars, "")
}
func GroupAnagrams(words []string) [][]string {
	anagrams := make(map[string][]string)
	for _, w := range words {
		sw := splitAndSort(w)
		if _, exists := anagrams[sw]; exists {
			anagrams[sw] = append(anagrams[sw], w)
		} else {
			anagrams[sw] = []string{w}
		}
	}
	return values(anagrams)
}

// Problem: Return the 'k' most frequent elements.
// Approach:
//
//	-Use a map to create frequencies count
//	-Get the map K:V pairs
//	-Sort pairs by value
func TopKFrequencies(nums []int, k int) []int {
	frequencies := make(map[int]int)
	for _, num := range nums {
		frequencies[num] += 1
	}
	entries := make([][]int, 0, len(frequencies))
	for key, val := range frequencies {
		entries = append(entries, []int{key, val})
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[j][1] < entries[i][1]
	})
	fmt.Println(entries)
	return []int{}
}

// Problem: Return true if the given string is a palindrome.
// Approach:
//
//	-Remove all non-alphanumeric characters
//	-Use two pointers to work inwards
func ValidPalindrome(s string) bool {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	result := reg.ReplaceAllString(s, "")
	chars := []rune(strings.ToLower(result))
	lhs := 0
	rhs := len(chars) - 1
	for lhs != rhs {
		if string(chars[lhs]) != string(chars[rhs]) {
			return false
		}
		lhs += 1
		rhs -= 1
	}
	return true
}
