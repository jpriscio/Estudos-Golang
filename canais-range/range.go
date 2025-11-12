package main

import "fmt"

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			fmt.Println(false, i)
		}
	}

	return true
}

func main() {
	isPrimo(100)
}
