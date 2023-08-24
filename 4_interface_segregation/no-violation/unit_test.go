package interfacesegregation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVisitanteBuscaProdutos(t *testing.T) {
	var v IVisitante
	v = &Visitante{}
	r := v.BuscarProdutos("23")
	assert.NotEmpty(t, r)
}

func TestVisitanteAdicionaAoCarrinho(t *testing.T) {
	var v IVisitante
	v = &Visitante{}
	v.BuscarProdutos("23")
	ok := v.AdicionarAoCarrinho(23)

	assert.Equal(t, true, ok)

}