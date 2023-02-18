package interfacesegregation

type UserLoggedService interface {
	BuscarProdutos(nome string)
	AdicionarAoCarrinho(productId int32)
	FormaDePagamento(forma string)
	GetNivelUsuario()
	//
	// outros métodos adicionais
	//
}

type VisitanteService interface {
	BuscarProdutos(nome string)
	AdicionarAoCarrinho(productId int32)
	//
	// outros métodos adicionais
	//
}

type VisitanteImplements struct {
	//
	// some fields
	//
}

func (v *VisitanteImplements) BuscarProdutos(nome string) {
	//
	// código
	//
}

func (v *VisitanteImplements) AdicionarAoCarrinho(productId int32) {
	//
	// código
	//
}

func (v* VisitanteImplements) FormaDePagamento(forma string) {

}

func (v* VisitanteImplements) GetNivelUsuario() {

}


type ServicoUsuarioLogado struct {
	//
	// atributos
	//
}

func (u *ServicoUsuarioLogado) BuscarProdutos(nome string) {
	//
	// código
	//
}

func (u *ServicoUsuarioLogado) AdicionarAoCarrinho(productId int32) {
	//
	// código
	//
}

func (u *ServicoUsuarioLogado) FormaDePagamento(forma string) {
	//
	// código
	//
}

func (u *ServicoUsuarioLogado) GetNivelUsuario() {
	//
	// código
	//
}
