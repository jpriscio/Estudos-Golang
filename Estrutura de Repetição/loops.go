package main

import (
	"fmt"
)

func main() {

	fmt.Println(mergeAlternately("abc", "efdgh"))

}

func mergeAlternately(word1 string, word2 string) string {

	a := make([]byte, 0, len(word1)+len(word2))
	i, j := 0, 0
	for i < len(word1) || j < len(word2) {
		if i < len(word1) {
			a = append(a, word1[i])
			a = append(a, ' ')
			i++
		}

		if j < len(word2) {
			a = append(a, word2[j])
			a = append(a, ' ')
			j++
		}
	}

	return string(a)

}

/*Input: word1 = "abc", word2 = "pqr"
Output: "apbqcr"
Explanation: The merged string will be merged as so:
word1:  a   b   c
word2:    p   q   r */
