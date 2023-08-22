package main

import (
	"fmt"
)

// Definindo um tipo chamado "Pessoa"
type Pessoa struct {
	Nome  string
	Idade int
}

// Método associado ao tipo "Pessoa"
func (p Pessoa) Apresentar() {
	fmt.Printf("Olá, eu sou %s e tenho %d anos.\n", p.Nome, p.Idade)
}

func main() {
	// Criando uma instância do tipo "Pessoa"
	pessoa1 := Pessoa{
		Nome:  "João",
		Idade: 30,
	}

	// Chamando o método "Apresentar" da instância "pessoa1"
	pessoa1.Apresentar()
}