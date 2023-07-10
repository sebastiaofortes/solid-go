package violation2

import (
	"strconv"
)

func ConverterNota(nota string) (float64, error) {
	notaF, err := strconv.ParseFloat(nota, 64)
	if err != nil {
		return 0, err
	}
	return notaF, nil
}

type Turma interface {
	ObterNota(nomeAluno string) (string, error)
}

func VerificarAluno(t Turma, aluno string) (string, error) {
	nota, err := t.ObterNota(aluno)
	if err != nil {
		return "", err
	} else {
		aprovado, err := ConverterNota(nota)
		if err != nil {
			return "", err
		}
		if aprovado >= 6.0 {
			return "O aluno foi aprovado!", nil
		} else {
			return "O aluno não foi aprovado.", nil
		}
	}
}
