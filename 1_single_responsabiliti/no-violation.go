package singleresponsabiliti

import (
	"database/sql"
	"fmt"
	"net/http"
)

var conexao *sql.DB

func Conectar() {
	conexao, _ = sql.Open("mysql", "datasource")
}

func Consultar(query string) string{
	consulta := conexao.QueryRow(query)
	var resultado bool
	consulta.Scan(&resultado)
	return fmt.Sprintln(resultado)
}

func Correct(resposta http.ResponseWriter, requisicao *http.Request) {
	// procuro dados no banco
	resultado := Consultar("SELECT * FROM products WHERE id = 1")
	
	// envio os dados para o cliente (navegador)
	resposta.Write([]byte(resultado))
}