package create_product

import (
	"context"
	"pos-graduacao-go-lang/graphql/internal/entity"
	"pos-graduacao-go-lang/graphql/internal/repository"
)

type Input struct {
	Context     context.Context
	Name        string
	Description *string
	Price       float64
	Categories  []string
}

type Output struct {
	Id string `json:"id"`
}

type UseCase struct {
	productRepository repository.ProductRepository
}

func New(productRepository repository.ProductRepository) *UseCase {
	return &UseCase{productRepository: productRepository}
}

func (useCase *UseCase) Execute(input Input) (*Output, error) {
	product := entity.NewProduct(input.Name, input.Description, input.Price, input.Categories)
	err := useCase.productRepository.Create(input.Context, product)
	if err != nil {
		return nil, err
	}
	return &Output{Id: product.ID}, nil
}
