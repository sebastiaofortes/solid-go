package liskovsubstitution

import "log"

type NotasPai struct {
	lis []float32
}

func (n NotasPai) Filtrar(f float32) []float32 {
	var result []float32
	for i := range n.lis {
		if n.lis[i] > f {
			result = append(result, n.lis[i])
		}
	}
	return result
}

func (n NotasPai) PrintPrimeiro() {
	log.Println(n.lis[0])
}



type NotasFilho struct {
	NotasPai
}

//overide
func (n NotasFilho) Filtrar(f float32) []float32 {
	var result []float32
	for i := range n.lis {
		if n.lis[i] < f {
			result = append(result, n.lis[i])
		}
	}
	return result
}

func Situacão1() {
	//quero obter o primeiro elemento maior que determinado número
	a := new(NotasPai)

	a.lis = []float32{1,2, 3,4,5}

	a.lis = a.Filtrar(1)

	a.PrintPrimeiro()

}

func Situacão2() {
	//quero obter a lista de elementos menor que detarminado númoro
	b := new(NotasFilho)

	b.lis = []float32{1,2, 3,4,5}

	b.lis = b.Filtrar(3)

	b.PrintPrimeiro()

}

func Situacão3() {
	//aplicando filho onde está o pai
	a := new(NotasFilho)

	a.lis = []float32{1,2, 3,4,5}

	a.lis = a.Filtrar(1)

	a.PrintPrimeiro()
}