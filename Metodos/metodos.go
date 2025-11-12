package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (u usuario) salvar() {
	fmt.Println("Salvando dados do usuario", u.nome, "que tem idade", u.idade)
}

func (u usuario) verificarIdade() {
	if u.idade < 18 {
		fmt.Println("cadeia")
	} else {
		fmt.Println("É Guichooo")
	}
}

func (u *usuario) fazAniversario() {
	u.idade++
}

func main() {

	usuario1 := usuario{"Vanessa", 21}
	usuario1.fazAniversario()
	usuario1.salvar()

}
