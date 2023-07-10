package violation2

import (
	"testing"

	"github.com/sebastiaofortes/solid-go/3_liskov_substitution/no-violation/escola_a"
	"github.com/sebastiaofortes/solid-go/3_liskov_substitution/no-violation/escola_b"
)

func TestA(t *testing.T) {
	turma1 := escola_a.Turma{
		Alunos: []escola_a.Aluno{
			{Nome: "João", Nota: "8.5"},
			{Nome: "Maria", Nota: "7.2"},
			{Nome: "Pedro", Nota: "4.9"},
		},
	}

	_, err := VerificarAluno(turma1, "Maria")
	if err != nil {
		t.Error(err)
	}
}

func TestB(t *testing.T) {
	turma2 := escola_b.Turma{
		Alunos: []escola_b.Aluno{
			{Nome: "João", Nota: "A"},
			{Nome: "Maria", Nota: "C"},
			{Nome: "Pedro", Nota: "D"},
		},
	}

	_, err := VerificarAluno(turma2, "Maria")
	if err != nil {
		t.Error(err)
	}
}
