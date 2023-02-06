package singleresponsabiliti

import (
	"database/sql"
	"net/http"
)

func Incorrect(resposta http.ResponseWriter, requisicao *http.Request) {
	// crio a conexao com o banco
	conexao, _ := sql.Open("mysql", "datasource")
	
	// procuro dados no banco
	consulta, _ := conexao.Query("SELECT * FROM products")

	// guardo resultado da procura em uma variável
	var resultado []byte 
	consulta.Scan(&resultado)

	// envio os dados para o cliente (navegador)
	resposta.Write(resultado)
}