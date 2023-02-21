package singleresponsabiliti

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

func (p *ProductDao) ProductListDao() ([]string, error) {
	// procuro dados no banco
	consulta, err := p.conexao.Query("SELECT 'name' FROM products")
	if err != nil {
		return []string{}, fmt.Errorf("Erro ao realizar consulta")
	}

	// guardo resultado da procura em uma variável
	var resultado []string
	consulta.Scan(&resultado)

	// retorna uma lista de produtos
	return resultado, nil
}
