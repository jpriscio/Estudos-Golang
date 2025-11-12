package main

import "fmt"

type usuario struct {
	nome  string
	idade int8
}

type estudante struct {
	usuario
	curso string
}

func main() {
	fmt.Println("Struct")

	estudante1 := estudante{usuario{"Priscio", 21}, "ciencia"}

	fmt.Println(estudante1.nome)
}
