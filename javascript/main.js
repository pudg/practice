import {
   containsDuplicates,
    validAnagram,
    twoSum,
    groupAnagrams,
    topKFrequencies,
    validPalindrome
 } from "./ten/ten.js";


 const main = () => {
    console.log(containsDuplicates([1, 1]));
    console.log(validAnagram("test", "ttes"));
    console.log(twoSum([8, 5, 1, 8, 5], 6));
    console.log(groupAnagrams(["eat","tea","tan","ate","nat","bat"]));
    console.log(topKFrequencies([1, 1, 2, 2, 2, 3], 1));
    console.log(validPalindrome("A man, a plan, a canal: Panama"));
    console.log(validPalindrome("race a car"));
 }
 main();