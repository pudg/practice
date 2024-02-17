# Problem: Return true if number array contains duplicates
# Approach: 
#   - Iterate over the nums array
#   - Populate a set of numbers
#   - Check if current number is in set
def contains_duplicate(nums):
    uniques = set()
    for num in nums:
        if num in uniques:
            return True
        uniques.add(num)
    return False

# Problem: Check if two strings are anagrams of each other.
# Approach:
#  -Split strings into characters
#  -Sort characters
#  -Combine sorted characters and check for equality
def valid_anagram(one, two):
    one, two = list(one), list(two)
    one.sort()
    two.sort()

    one_sorted = "".join([i for i in one])
    two_sorted = "".join([i for i in two])
    print(one_sorted, two_sorted)
    return one_sorted == two_sorted