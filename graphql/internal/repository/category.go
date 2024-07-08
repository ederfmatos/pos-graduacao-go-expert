package repository

import (
	"context"
	"database/sql"
	"pos-graduacao-go-lang/graphql/internal/entity"
)

type ListCategoryOutput struct {
	Id          string
	Name        string
	Description *string
}

type CategoryRepository interface {
	Create(ctx context.Context, category *entity.Category) error
	List(ctx context.Context) ([]*ListCategoryOutput, error)
	FindById(ctx context.Context, id string) (*entity.Category, error)
	ListByProduct(ctx context.Context, id string) ([]*entity.Category, error)
}

const (
	sqlInsertCategory            = `INSERT INTO categories (id, name, description, created_at) VALUES (?, ?, ?, ?)`
	sqlListCategories            = "SELECT id, name, description FROM categories ORDER BY created_at DESC"
	sqlListCategoriesByProductId = `
		SELECT c.id, c.name, c.description 
		FROM categories as c 
		LEFT JOIN product_categories pc on c.id = pc.category_id
		WHERE pc.product_id = ?
		ORDER BY created_at DESC
	`
	sqlCategoryById = "SELECT id, name, description FROM categories WHERE id = ?"
)

type SqlCategoryRepository struct {
	database *sql.DB
}

func NewSqlCategoryRepository(database *sql.DB) CategoryRepository {
	return &SqlCategoryRepository{database: database}
}

func (repository *SqlCategoryRepository) Create(ctx context.Context, category *entity.Category) error {
	stmt, err := repository.database.PrepareContext(ctx, sqlInsertCategory)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, category.ID, category.Name, category.Description, category.CreatedAt)
	return err
}

func (repository *SqlCategoryRepository) List(ctx context.Context) ([]*ListCategoryOutput, error) {
	result, err := repository.database.QueryContext(ctx, sqlListCategories)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var categories []*ListCategoryOutput
	for result.Next() {
		var category ListCategoryOutput
		err = result.Scan(&category.Id, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (repository *SqlCategoryRepository) FindById(ctx context.Context, id string) (*entity.Category, error) {
	statement, err := repository.database.PrepareContext(ctx, sqlCategoryById)
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	var category entity.Category
	err = statement.QueryRow(id).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (repository *SqlCategoryRepository) ListByProduct(ctx context.Context, id string) ([]*entity.Category, error) {
	statement, err := repository.database.PrepareContext(ctx, sqlListCategoriesByProductId)
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	result, err := statement.QueryContext(ctx, id)
	if err != nil {
		return nil, err
	}
	var categories []*entity.Category
	for result.Next() {
		var category entity.Category
		err = result.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}
