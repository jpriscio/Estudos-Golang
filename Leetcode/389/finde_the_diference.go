package main

import (
	"fmt"
)

func findTheDifference(s string, t string) byte {
	somaS, somaT := 0, 0
	var result int

	for _, ch := range s {
		somaS += int(ch)
	}

	for _, ch := range t {
		somaT += int(ch)
	}
	result = somaT - somaS
	return byte(result)

}

func main() {

	fmt.Println(string(findTheDifference("abc", "abcr")))

}
