package main

import (
	"fmt"
)

// Definição da classe pai (superclasse)
type Animal struct {
	Name   string
	Sound  string
}

// Método da classe pai
func (a *Animal) MakeSound() {
	fmt.Printf("%s makes a %s sound\n", a.Name, a.Sound)
}

// Definição da classe filha (subclasse)
type Dog struct {
	Animal      // Incorpora a classe Animal
	Breed string
}

func main() {
	// Criando um objeto da classe filha
	d := Dog{
		Animal: Animal{Name: "Buddy", Sound: "Woof"},
		Breed:  "Golden Retriever",
	}

	// Chamando um método da classe pai através da classe filha
	d.MakeSound()

	// Acessando diretamente os campos da classe pai através da classe filha
	fmt.Printf("%s is a %s\n", d.Name, d.Breed)

	//Polimorfismo por herança não funciona em GO
	//var a = Animal{}
	//a = d 
}