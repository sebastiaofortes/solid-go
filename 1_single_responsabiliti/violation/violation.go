package singleresponsabiliti

import (
	"database/sql"
)

func ProductListDao() []string {
	// crio a conexao com o banco
	conexao, _ := sql.Open("mysql", "datasource")

	// procuro dados no banco
	consulta, _ := conexao.Query("SELECT 'name' FROM products")

	// guardo resultado da procura em uma variável
	var resultado []string
	consulta.Scan(&resultado)

	// retorna uma lista de produtos
	return resultado
}
