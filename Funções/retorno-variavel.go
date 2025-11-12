package main

import "fmt"

func mat(a, b int) (soma, sub int) {
	soma = a + b
	sub = a - b
	return
}

func main() {

	soma, sub := mat(10, 5)
	fmt.Println(soma, sub)

}
