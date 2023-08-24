package violation2

type iTurma interface {
	ObterNota(nomeAluno string) (float64, error)
}

type Secretaria struct{
	turma []iTurma
}

func (s Secretaria)VerificarAluno(t int32, aluno string) (string, error) {
	nota, err := s.turma[t].ObterNota(aluno)
	if err != nil {
		return "", err
	} else {
		if nota >= 6.0 {
			return "O aluno foi aprovado!", nil
		} else {
			return "O aluno não foi aprovado.", nil
		}
	}
}
