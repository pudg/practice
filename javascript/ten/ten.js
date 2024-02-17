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