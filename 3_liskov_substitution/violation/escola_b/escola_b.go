package escola_b

import (
	"fmt"
)

type Aluno struct {
	Nome string
	Nota string
	Sexo string
}

type Turma struct {
	Alunos []Aluno
}

func (t Turma) ObterNota(nomeAluno string) (string, error) {
	for _, aluno := range t.Alunos {
		if aluno.Nome == nomeAluno {

			return aluno.Nota, nil
		}
	}
	return "", fmt.Errorf("Aluno '%s' não encontrado na turma", nomeAluno)
}
