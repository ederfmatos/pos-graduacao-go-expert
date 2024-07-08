package repository

import (
	"context"
	"database/sql"
	"pos-graduacao-go-lang/graphql/internal/entity"
)

type ListProductOutput struct {
	Id          string
	Name        string
	Description *string
	Price       float64
}

type ProductRepository interface {
	Create(ctx context.Context, product *entity.Product) error
	List(ctx context.Context) ([]*ListProductOutput, error)
	FindById(ctx context.Context, id string) (*entity.Product, error)
	ListByCategory(ctx context.Context, id string) ([]*entity.Product, error)
}

const (
	sqlListProducts             = `SELECT id, name, description, price FROM products ORDER BY created_at DESC`
	sqlListProductsByCategoryId = `
		SELECT p.id, p.name, p.description, p.price 
		FROM products as p 
		LEFT JOIN product_categories pc on p.id = pc.product_id
		WHERE pc.category_id = ?
		ORDER BY created_at DESC
	`
	sqlInsertProduct         = `INSERT INTO products (id, name, description, price, created_at) VALUES (?, ?, ?, ?, ?)`
	sqlInsertProductCategory = `INSERT INTO product_categories (product_id, category_id) VALUES (?, ?)`
	sqlProductById           = `SELECT id, name, description, price FROM products WHERE id = ?`
)

type SqlProductRepository struct {
	database *sql.DB
}

func NewSqlProductRepository(database *sql.DB) ProductRepository {
	return &SqlProductRepository{database: database}
}

func (repository *SqlProductRepository) Create(ctx context.Context, product *entity.Product) error {
	transaction, err := repository.database.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	stmt, err := transaction.PrepareContext(ctx, sqlInsertProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, product.ID, product.Name, product.Description, product.Price, product.CreatedAt)
	if err != nil {
		transaction.Rollback()
		return err
	}
	for _, category := range product.Categories {
		err = func() error {
			stmt, err := transaction.PrepareContext(ctx, sqlInsertProductCategory)
			defer stmt.Close()
			if err != nil {
				return err
			}
			_, err = stmt.ExecContext(ctx, product.ID, category)
			if err != nil {
				return err
			}
			return nil
		}()
		if err != nil {
			transaction.Rollback()
			return err
		}
	}
	err = transaction.Commit()
	if err != nil {
		transaction.Rollback()
		return err
	}
	return nil
}

func (repository *SqlProductRepository) List(ctx context.Context) ([]*ListProductOutput, error) {
	result, err := repository.database.QueryContext(ctx, sqlListProducts)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var products []*ListProductOutput
	for result.Next() {
		var product ListProductOutput
		err = result.Scan(&product.Id, &product.Name, &product.Description, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (repository *SqlProductRepository) ListByCategory(ctx context.Context, id string) ([]*entity.Product, error) {
	statement, err := repository.database.PrepareContext(ctx, sqlListProductsByCategoryId)
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	result, err := statement.QueryContext(ctx, id)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var products []*entity.Product
	for result.Next() {
		var product entity.Product
		err = result.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (repository *SqlProductRepository) FindById(ctx context.Context, id string) (*entity.Product, error) {
	statement, err := repository.database.PrepareContext(ctx, sqlProductById)
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	var product entity.Product
	err = statement.QueryRow(id).Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
