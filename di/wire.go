//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/becardine/learning-go/di/product"
	"github.com/google/wire"
)

var setRepositoryDependency = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)))

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(setRepositoryDependency, product.NewProductUseCase)
	return &product.ProductUseCase{}
}
