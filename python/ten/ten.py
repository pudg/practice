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

# Problem: Return indices of two numbers that sum to target
# Approach:
#   -Create a mapping of numbers to locations
#   -Iterate over numbers, and calculate delta
#   -Check if delta exists within map
def two_sum(nums, target):
    locations = dict()
    for idx, num in enumerate(nums):
        if num not in locations:
            locations[num] = [idx]
        else:
            locations[num].append(idx)
    for idx, num in enumerate(nums):
        delta = target - num
        if delta in locations:
            if len(locations[delta]) >= 2:
                return [idx, locations[delta][1]]
            if locations[delta][0] != idx:
                return [idx, locations[delta][0]]
    return []


# Problem: Return a list of grouped anagrams.
# Approach:
#   -Create a map: string -> [string]
#   -Sort the strings
#   -Check if sorted string exists in map
def group_anagrams(words):
    anagrams = dict()
    for w in words:
        sw = sorted(list(w))
        sw = "".join([c for c in sw])
        if sw in anagrams:
            anagrams[sw].append(w)
        else:
            anagrams[sw] = [w]
    return anagrams.values()

# Problem: Return the 'k' most frequent elements.
# Approach:
#   -Use a map to create frequencies count
#   -Get the map K:V pairs
#   -Sort pairs by value
def top_k_frequencies(nums, k):
    from collections import defaultdict
    frequencies = defaultdict(int)
    for num in nums:
        frequencies[num] += 1
    pairs = sorted(frequencies.items(), reverse=True, key=lambda x : x[1])
    return [pairs[i][0] for i in range(k)]


# Problem: Return true if the given string is a palindrome.
# Approach:
#   -Remove all non-alphanumeric characters
#   -Use two pointers to work inwards
def valid_palindrome(s):
    import re
    regex = re.compile(r'[a-zA-Z]+')
    matches = regex.findall(s)
    result = "".join(matches).lower()
    lhs, rhs = 0, len(result) - 1
    while lhs != rhs:
        if result[lhs] != result[rhs]:
            return False
        lhs += 1
        rhs -= 1
    return True