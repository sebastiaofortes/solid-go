package interfacesegregation

type UsuarioService interface {
	BuscarProdutos(nome string) string
	AdicionarAoCarrinho(productId int32) bool
	FormaDePagamento(forma string)
	GetNivelUsuario()
	//
	// outros métodos adicionais
	//
}