package violation2

import (
	"strconv"
)

type iTurma interface {
	ObterNota(nomeAluno string) (string, error)
}

type Secretaria struct{
	turma []iTurma
}

func (s Secretaria)VerificarAluno(t int32, aluno string) (string, error) {
	nota, err := s.turma[t].ObterNota(aluno)
	if err != nil {
		return "", err
	} else {
		aprovado, err := converterNota(nota)
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

func converterNota(nota string) (float64, error) {
	notaF, err := strconv.ParseFloat(nota, 64)
	if err != nil {
		return 0, err
	}
	return notaF, nil
}


