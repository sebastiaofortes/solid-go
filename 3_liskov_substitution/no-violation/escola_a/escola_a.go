package escola_a

import (
	"fmt"
	"strconv"
)

type Aluno struct {
	Nome string
	Nota string
}

type Turma struct {
	Alunos []Aluno
}

func (t Turma) ObterNota(nomeAluno string) (float64, error) {
	for _, aluno := range t.Alunos {
		if aluno.Nome == nomeAluno {
			nota, err := t.converterNota(aluno.Nota)
			if err != nil{
				return 0, err
			}
			return nota, nil
		}
	}

	return 0, fmt.Errorf("Aluno '%s' não encontrado na turma", nomeAluno)
}

func (t Turma) converterNota(nota string) (float64, error) {
	notaF, err := strconv.ParseFloat(nota, 64)
	if err != nil {
		return 0, err
	}
	return notaF, nil
}

