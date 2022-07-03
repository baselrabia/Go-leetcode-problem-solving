package main

import (
	"fmt"
	"math"
)

func lengthOfLongestSubstring(s string) int {

	m1 := make(map[string]int)
	longest := 0
	l := 0

	chars := []rune(s)
	for r := 0; r < len(chars); r++ {
		char := string(chars[r])

		for _, ok := m1[char]; ok; {
			if _, ok2 := m1[char]; ok2 == false {
				break
			}
			dChar := string(chars[l])

			delete(m1, dChar)
			l++
		}

		m1[char] = r
		longest = int(math.Max(float64(longest), float64(r-l+1)))
	}

	return longest
}

func main() {

	s1 := "abcabcbb"
	//s1 := "bbbbb"
	//s1 := "pwwkew"

	i := lengthOfLongestSubstring(s1)

	fmt.Println(i)
}
