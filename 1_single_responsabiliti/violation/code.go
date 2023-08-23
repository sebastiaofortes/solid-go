package singleresponsabiliti

import (
	"database/sql"
	"fmt"
)

type Product struct {
	conexao *sql.DB
}

func NewProduct() (Product, error) {
	conexao, err := sql.Open("mysql", "datasource")
	if err != nil {
		return Product{}, fmt.Errorf("Erro ao criar objeto")
	}
	return Product{conexao: conexao}, nil
}

func (p *Product) ObterPreço(produtoID int64) (float64, error) {
	// procuro dados no banco
	consulta, err := p.conexao.Query("SELECT 'preço' FROM products WHERE id = ?", produtoID)
	if err != nil {
		return 0, fmt.Errorf("Erro ao realizar consulta")
	}

	// guardo resultado da procura em uma variável
	var resultado float64
	consulta.Scan(&resultado)

	// retorna o resultado
	return resultado, nil
}

func (p *Product) calcularTotalCompra(quantidade int64, produtoID int64) (float64, error) {
	
	// procuro dados no banco
	preço, _ := p.ObterPreço(produtoID)

	total := preço * float64(quantidade)


	// retorna o resultado
	return total, nil
}