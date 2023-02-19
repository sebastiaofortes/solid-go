package interfacesegregation

type UserService interface {
	BuscarProdutos(nome string)
	AdicionarAoCarrinho(productId int32)
	FormaDePagamento(forma string)
	GetNivelUsuario()
	//
	// outros métodos adicionais
	//
}

type Visitante struct {
	//
	// some fields
	//
}

func (v *Visitante) BuscarProdutos(nome string) {
	//
	// código
	//
}

func (v *Visitante) AdicionarAoCarrinho(productId int32) {
	//
	// código
	//
}

func (v* Visitante) FormaDePagamento(forma string) {

}

func (v* Visitante) GetNivelUsuario() {

}


type UsuarioLogado struct {
	//
	// atributos
	//
}

func (u *UsuarioLogado) BuscarProdutos(nome string) {
	//
	// código
	//
}

func (u *UsuarioLogado) AdicionarAoCarrinho(productId int32) {
	//
	// código
	//
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
