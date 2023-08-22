package main

import (
	"fmt"
)

// Definição da interface
type Shape interface {
	//Interfaces possuem somente métodos
	// Interfaces são abstratas, ou seja: não é possível criar um objeto do tipo shape
	Area() float64
}

// Definição das structs que implementam a interface
type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// Implementação dos métodos da interface para Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Implementação dos métodos da interface para Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	// Criando objetos das structs
	circle := Circle{Radius: 3}
	rectangle := Rectangle{Width: 4, Height: 5}

	// Chamando a função que calcula a área utilizando a interface
	// Isso permite que objetos de diferentes tipos (Circle e Rectangle) sejam tratados de forma uniforme
	printArea(circle)
	printArea(rectangle)
}

// Função que recebe um objeto que implementa a interface Shape e imprime sua área
func printArea(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
}