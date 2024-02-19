package main

import (
	"fmt"
	"practice/ten"
)

func main() {
	fmt.Println(ten.ValidAnagrams("cider", "cried"))
	fmt.Println(ten.TwoSum([]int{2, 7, 13, 13, 7}, 20))
	fmt.Println(ten.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
