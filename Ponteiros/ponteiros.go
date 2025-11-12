package main

import "fmt"

func main() {

	x := 10
	z := &x

	x++
	fmt.Println(x, *z)

}
