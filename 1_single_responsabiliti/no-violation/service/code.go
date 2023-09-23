package service

import (
	"github.com/sebastiaofortes/solid-go/1_single_responsabiliti/no-violation/dao"
)

type ProductService struct {
	d dao.ProductDao
}

func NewProductDao(d dao.ProductDao) (ProductService, error) {

	return ProductService{d: d}, nil
}

func (p *ProductService) CalcularTotalCompra(quantidade int64, produtoID int64) (float64, error) {

	// procuro dados no banco
	preço, _ := p.d.ObterPreço(produtoID)

	total := preço * float64(quantidade)

	// retorna o resultado
	return total, nil
}
