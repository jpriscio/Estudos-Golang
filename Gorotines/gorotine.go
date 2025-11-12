// Spinner shows an animated spinner while computing Fibonacci.
package main

import (
	"fmt"
	"time"
)

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {
	go escrever("ola mundo")
	escrever("Larissa")
}
