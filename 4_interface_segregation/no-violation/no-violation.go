package interfacesegregation

type UserService interface {
	BuscarProdutos(nome string) string
	AdicionarAoCarrinho(productId int32) bool
	FormaDePagamento(forma string)
	GetNivelUsuario()
	//
	// outros métodos adicionais
	//
}

type VisitanteServiceInterface interface {
	BuscarProdutos(nome string) string
	AdicionarAoCarrinho(productId int32) bool
	//
	// outros métodos adicionais
	//
}

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
	return true
}

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
