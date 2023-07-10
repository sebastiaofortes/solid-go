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

func (t Turma) ObterNota(nomeAluno string) (float64, error) {
	for _, aluno := range t.Alunos {
		if aluno.Nome == nomeAluno {
			nota, err := t.converterNota(aluno.Nota)
			if err != nil {
				return 0, err
			}
			return nota, nil
		}
	}

	return 0, fmt.Errorf("Aluno '%s' não encontrado na turma", nomeAluno)
}

func (t Turma) converterNota(nota string) (float64, error) {

	dataMap := map[string]float64{
		"A": 10.00,
		"B": 8.00,
		"C": 6.00,
		"D": 4.00,
		"E": 2.00,
		"F": 0.00,
	}

	notaF, ok := dataMap[nota]
	if !ok {
		return 0, fmt.Errorf("Nota correspondente não encontrada")
	}

	return notaF, nil

}
