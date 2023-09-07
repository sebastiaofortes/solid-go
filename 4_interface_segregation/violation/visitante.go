package interfacesegregation

type Visitante struct {
	//
	// some fields
	//
}

func (v *Visitante) BuscarProdutos(nome string) string {
	//
	// código
	//
	return "product json"
}

func (v *Visitante) AdicionarAoCarrinho(productId int32) bool {
	//
	// código
	//
	return false
}

func (v* Visitante) FormaDePagamento(forma string) {

}