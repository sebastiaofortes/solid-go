package interfacesegregation

type UsuarioLogado struct {
	//
	// atributos
	//
}

func (u *UsuarioLogado) BuscarProdutos(nome string) string {
	//
	// código
	//
	return "product json"
}

func (u *UsuarioLogado) AdicionarAoCarrinho(productId int32) bool {
	//
	// código
	//
	return true
}

func (u *UsuarioLogado) FormaDePagamento(forma string) {
	//
	// código
	//
}

func (u *UsuarioLogado) GetNivelUsuario() {
	//
	// código
	//
}
