package interfacesegregation

type VisitanteServiceInterface interface {
	BuscarProdutos(nome string) string
	AdicionarAoCarrinho(productId int32) bool
	//
	// outros métodos adicionais
	//
}