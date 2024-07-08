package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"pos-graduacao-go-lang/graphql/graph/model"
	"pos-graduacao-go-lang/graphql/internal/usecase/create_category"
	"pos-graduacao-go-lang/graphql/internal/usecase/create_product"
)

// Products is the resolver for the products field.
func (r *categoryResolver) Products(ctx context.Context, obj *model.Category) ([]*model.Product, error) {
	products, err := r.ProductRepository.ListByCategory(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	var result []*model.Product
	for _, product := range products {
		result = append(result, &model.Product{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		})
	}
	return result, nil
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, category model.NewCategory) (string, error) {
	input := create_category.Input{
		Context:     ctx,
		Name:        category.Name,
		Description: category.Description,
	}
	output, err := r.CreateCategoryUseCase.Execute(input)
	if err != nil {
		return "", err
	}
	return output.Id, nil
}

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, product model.NewProduct) (string, error) {
	input := create_product.Input{
		Context:     ctx,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Categories:  product.Categories,
	}
	output, err := r.CreateProductUseCase.Execute(input)
	if err != nil {
		return "", err
	}
	return output.Id, nil
}

// Categories is the resolver for the categories field.
func (r *productResolver) Categories(ctx context.Context, obj *model.Product) ([]*model.Category, error) {
	categories, err := r.CategoryRepository.ListByProduct(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	var result []*model.Category
	for _, category := range categories {
		result = append(result, &model.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	return result, nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.CategoryList, error) {
	categories, err := r.CategoryRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	var result []*model.CategoryList
	for _, category := range categories {
		result = append(result, &model.CategoryList{
			ID:          category.Id,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	return result, nil
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.ProductList, error) {
	products, err := r.ProductRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	var result []*model.ProductList
	for _, product := range products {
		result = append(result, &model.ProductList{
			ID:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		})
	}
	return result, nil
}

// CategoryByID is the resolver for the categoryById field.
func (r *queryResolver) CategoryByID(ctx context.Context, id string) (*model.Category, error) {
	category, err := r.CategoryRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &model.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

// ProductByID is the resolver for the productById field.
func (r *queryResolver) ProductByID(ctx context.Context, id string) (*model.Product, error) {
	product, err := r.ProductRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &model.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}, nil
}

// Category returns CategoryResolver implementation.
func (r *Resolver) Category() CategoryResolver { return &categoryResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Product returns ProductResolver implementation.
func (r *Resolver) Product() ProductResolver { return &productResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type categoryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }