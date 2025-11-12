package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)

	go DoisTresQuatro(5, c)

	a, b := <-c, <-c

	fmt.Println(a, b)
	fmt.Println(<-c)

}

func DoisTresQuatro(b int, c chan int) {

	c <- b * 2

	time.Sleep(time.Second * 2)
	c <- b * 3

	time.Sleep(time.Second * 4)
	c <- b * 4
}
