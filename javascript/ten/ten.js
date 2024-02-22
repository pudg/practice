// Problem: Check if the given slice contains duplicates.
// Approach:
//	-Iterate over the slice of ints
//	-Use a map to memorize previously seen integers
//	-Check if current integer is in memory
export const containsDuplicates = (nums) => {
    let memory = new Set();
    for (const num of nums) {
        if (memory.has(num)) {
            return true;
        }
        memory.add(num);
    }
    return false;
}

// Problem: Check if two strings are anagrams of each other.
// Approach:
//  -Split strings into characters
//  -Sort characters
//  -Combine sorted characters and check for equality
export const validAnagram = (one, two) => {
    let one_loc = one.split("");
    let two_loc = two.split("");

    one_loc.sort();
    two_loc.sort();

    return one_loc.join("") == two_loc.join("");
}

// Problem: Return indices of two numbers that sum to target
// Approach:
//   -Create a mapping of numbers to locations
//   -Iterate over numbers, and calculate delta
//   -Check if delta exists within map
export const twoSum = (nums, target) => {
    const locations = new Map();
    for (const idx in nums) {
        if (locations.has(nums[idx])) {
            const locs = locations.get(nums[idx]);
            locs.push(idx);
            locations.set(nums[idx], locs);
        } else {
            locations.set(nums[idx], [idx]);
        }
    }

    for (const idx in nums) {
        const delta = target - nums[idx];
        if (locations.has(delta)) {
            const locs = locations.get(delta);
            if (locs.length >= 2) {
                return [idx, locs[1]];
            }
            if (locs[0] != idx) {
                return [idx, locs[0]];
            }
        }
    }
    return [];
}

// Problem: Return a list of grouped anagrams.
// Approach:
//   -Create a map: string -> [string]
//   -Sort the strings
//   -Check if sorted string exists in map
export const groupAnagrams = (words) => {
    const anagrams = {};
    for (const w of words) {
        let sw = w.split("");
        sw.sort()
        sw = sw.join("");
        if (sw in anagrams) {
            anagrams[sw].push(w);
        } else {
            anagrams[sw] = [w];
        }
    }
    return Object.values(anagrams);
}

// Problem: Return the 'k' most frequent elements.
// Approach:
//   -Use a map to create frequencies count
//   -Get the map K:V pairs
//   -Sort pairs by value
export const topKFrequencies = (nums, k) => {
    const frequencies = new Map();
    for (const num of nums) {
        frequencies.has(num) 
        ? frequencies.set(num, frequencies.get(num) + 1) 
        : frequencies.set(num, 1)
    }
    //convert map_entries to array, and sort in descending order
    let entries = [...frequencies.entries()].sort((a, b) => b[1] - a[1]);
    return entries.slice(0, k).map(entry => entry[0]);
}

// Problem: Return true if the given string is a palindrome.
// Approach:
//   -Remove all non-alphanumeric characters
//   -Use two pointers to work inwards
export const validPalindrome = (s) => {
    let result = s.replace(/[^a-zA-Z]/g, "");
    result = result.toLowerCase();
    let lhs = 0;
    let rhs = result.length - 1;
    while (lhs != rhs) {
        if (result[lhs] !== result[rhs]) {
            return false;
        }
        lhs += 1;
        rhs -= 1;
    }
    return true;
}