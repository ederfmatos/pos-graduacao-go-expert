package graph

import (
	"pos-graduacao-go-lang/graphql/internal/repository"
	"pos-graduacao-go-lang/graphql/internal/usecase/create_category"
	"pos-graduacao-go-lang/graphql/internal/usecase/create_product"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateCategoryUseCase *create_category.UseCase
	CreateProductUseCase  *create_product.UseCase
	ProductRepository     repository.ProductRepository
	CategoryRepository    repository.CategoryRepository
}
