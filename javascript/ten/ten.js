// Problem: Check if the given slice contains duplicates.
// Approach:
//
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