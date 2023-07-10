package escola_b

import (
	"testing"
)

func TestConverSucess(t *testing.T) {
	turma2 := Turma{
		Alunos: []Aluno{
			{Nome: "João", Nota: "A"},
			{Nome: "Maria", Nota: "C"},
			{Nome: "Pedro", Nota: "D"},
		},
	}

	n, err := turma2.converterNota("A")
	if err != nil{
		t.Error(err)
	}

	if n != 10.00 {
		t.Errorf("Valor incorreto %f", n)
	}
}
