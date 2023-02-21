package interfacesegregation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVisitanteNaoPodeUsarCarrinho(t *testing.T) {
	v := Visitante{}
	r := v.BuscarProdutos("23")
	assert.NotEmpty(t, r)
}

func TestVisitanteUsaCarrinho(t *testing.T) {
	v := Visitante{}
	v.BuscarProdutos("23")
	ok := v.AdicionarAoCarrinho(23)

	assert.Equal(t, true, ok)

}