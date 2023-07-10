package dao

import (
	"database/sql"
	"fmt"
)

type ProductDao struct {
	conexao *sql.DB
}

func NewProductDao() (ProductDao, error) {
	conexao, err := sql.Open("mysql", "datasource")
	if err != nil {
		return ProductDao{}, fmt.Errorf("Erro ao criar objeto")
	}
	return ProductDao{conexao: conexao}, nil
}

func (p *ProductDao) ObterPreço(produtoID int64) (float64, error) {
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

func (p *ProductDao) calcularTotalCompra(quantidade int64, produtoID int64) (float64, error) {
	
	// procuro dados no banco
	preço, _ := p.ObterPreço(produtoID)

	total := preço * float64(quantidade)


	// retorna o resultado
	return total, nil
}