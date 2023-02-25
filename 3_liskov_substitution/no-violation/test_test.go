package liskovsubstitution

import (
	"log"
	"testing"
)

func TestViolation(t *testing.T) {
	//quero obter o primeiro elemento maior que determinado número
	a := new(Notas)

	a.lis = []float32{1, 2, 3, 4, 5}

	log.Println(a.FiltrarAcima(3))

	log.Println(a.FiltrarAbaixo(2))
}

func TestViolation2(t *testing.T) {
	//quero obter a lista de elementos menor que detarminado númoro
	b := new(SubNotas)

	b.lis = []float32{1, 2, 3, 4, 5}

	log.Println(b.FiltrarAcima(3))

	log.Println(b.FiltrarAbaixo(2))
}
