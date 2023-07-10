package violation2

type Turma interface {
	ObterNota(nomeAluno string) (float64, error)
}

func VerificarAluno(t Turma, aluno string) (string, error) {
	nota, err := t.ObterNota(aluno)
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
