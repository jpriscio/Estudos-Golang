package main

import "fmt"

func diasDaSemana(numero int) string {

	switch numero {

	case 1:
		return "Domingo"

	case 2:
		return "Segunda"

	default:
		return "Numero invalido"
	}

}

func diasDaSemana2(numero int) string {

	switch {

	case numero == 1:

		return "Domingo"

	case numero == 2:
		return "Segunda"

	default:
		return "Numero invalido"
	}

}

func main() {

	fmt.Println("Switch")

	fmt.Println(diasDaSemana2(20))
}
