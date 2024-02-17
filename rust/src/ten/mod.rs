use std::collections::HashSet;

// Problem: Check if the given slice contains duplicates.
// Approach:
//
//	-Iterate over the slice of ints
//	-Use a map to memorize previously seen integers
//	-Check if current integer is in memory
pub fn contains_duplicates(nums: Vec<i32>) -> bool {
    let mut uniques: HashSet<i32> = HashSet::new();
    for i in 0..nums.len() {
        if uniques.contains(&nums[i]) {
            return true;
        }
        uniques.insert(nums[i]);
    }
    false
}

// Problem: Check if two strings are anagrams of each other.
// Approach:
//  -Split strings into characters
//  -Sort characters
//  -Combine sorted characters and check for equality
pub fn valid_anagrams(one: &str, two: &str) -> bool {
    let mut one_chars: Vec<char> = one.chars().collect();
    let mut two_chars: Vec<char> = two.chars().collect();

    one_chars.sort();
    two_chars.sort();

    let one_sorted: String = one_chars.into_iter().collect();
    let two_sorted: String = two_chars.into_iter().collect();
    one_sorted == two_sorted
}