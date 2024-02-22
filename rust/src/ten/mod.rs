use std::{collections::{HashMap, HashSet}, i32, iter::zip, vec};

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

// Problem: Return indices of two numbers that sum to target
// Approach:
//
//	-Create a mapping of numbers to locations
//	-Iterate over numbers, and calculate delta
//	-Check if delta exists within map
pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<usize> {
    let mut locations: HashMap<i32, Vec<usize>> = HashMap::new();
    for i in 0..nums.len() {
        locations
        .entry(nums[i])
        .or_insert(Vec::new())
        .push(i)
    }

    for i in 0..nums.len() {
        let delta = target - nums[i];
        if locations.contains_key(&delta) {
            let locs = locations.get(&delta).unwrap();
            if locs.len() >= 2 {
                return vec![i, locs[1]];
            }
            if locs[0] != i {
                return vec![i, locs[0]];
            }
        }
    }
    vec![]
}

// Problem: Return a list of grouped anagrams.
// Approach:
//   -Create a map: string -> [string]
//   -Sort the strings
//   -Check if sorted string exists in map
pub fn group_anagrams(words: Vec<String>) -> Vec<Vec<String>> {
    let mut anagrams: HashMap<String, Vec<String>> = HashMap::new();

    for i in 0..words.len() {
        let mut w: Vec<char> = words[i].chars().collect();
        w.sort();
        let w: String = w.into_iter().collect();
        anagrams
        .entry(w)
        .or_insert(Vec::new())
        .push(words[i].clone());
    }

    anagrams.values().cloned().collect()
}

// Problem: Return the 'k' most frequent elements.
// Approach:
//   -Use a map to create frequencies count
//   -Get the map K:V pairs
//   -Sort pairs by value
pub fn top_k_frequencies(nums: Vec<i32>, k: i32) -> Vec<i32> {
    let mut frequencies: HashMap<i32, i32> = HashMap::new();
    for i in 0..nums.len() {
        frequencies
            .entry(nums[i])
            .and_modify(|v| *v += 1)
            .or_insert(1);
    }

    let mut pairs: Vec<(i32, i32)> = frequencies.into_iter().collect();
    pairs.sort_by(|a, b| b.1.cmp(&a.1));

    let mut result: Vec<i32> = Vec::new();
    for i in 0..k {
        result.push(pairs[i as usize].0)

    }
    result
}

// Problem: Return true if the given string is a palindrome.
// Approach:
//   -Remove all non-alphanumeric characters
//   -Use two pointers to work inwards
use regex::Regex;
pub fn valid_palindrome(s: String) -> bool {
    let regex = Regex::new(r"[^a-zA-Z]").unwrap();
    let result = regex.replace_all(&s, "").to_lowercase();
    let pairs = zip(result.chars(), result.chars().rev());
    for pair in pairs {
        if pair.0 != pair.1 {
            return false
        }
    }
    true
}