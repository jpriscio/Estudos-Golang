package main

import "fmt"

func main() {

	mapas := make(map[int]string)

	mapas[123] = "joao"
	mapas[456] = "beatriz"

	mapa2 := map[string]int{
		"joao":  13,
		"pedro": 23,
		"laion": 56,
	}

	for nome, idade := range mapa2 {
		fmt.Printf("nome: %s idade %d\n", nome, idade)
	}

}
