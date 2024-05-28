package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	Id    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	id, _ := uuid.NewV7()
	return &Product{
		Id:    id.String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	database, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer database.Close()
	product := NewProduct("Coca Cola", 4.90)
	err = InsertProduct(database, product)
	if err != nil {
		panic(err)
	}
	product, err = SelectProduct(database, product.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Product is", *product)
	product.Name = "Coca Cola 2 Litros"
	product.Price = 9.50
	err = UpdateProduct(database, product)
	if err != nil {
		panic(err)
	}
	fmt.Println("Product is", *product)
	products, err := SelectProducts(database)
	if err != nil {
		panic(err)
	}
	fmt.Println("Products is", products)
	err = DeleteProduct(database, product.Id)
	fmt.Println("Product deleted")
	if err != nil {
		panic(err)
	}
	products, err = SelectProducts(database)
	if err != nil {
		panic(err)
	}
	fmt.Println("Products is", products)
}

func InsertProduct(database *sql.DB, product *Product) error {
	statement, err := database.Prepare("INSERT INTO products(id, name, price) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(product.Id, product.Name, product.Price)
	return err
}

func UpdateProduct(database *sql.DB, product *Product) error {
	statement, err := database.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(product.Name, product.Price, product.Id)
	return err
}

func SelectProduct(database *sql.DB, id string) (*Product, error) {
	statement, err := database.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	var product Product
	err = statement.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price)
	return &product, err
}

func SelectProducts(database *sql.DB) ([]Product, error) {
	rows, err := database.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var product Product
		err = rows.Scan(&product.Id, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, err
}

func DeleteProduct(database *sql.DB, id string) error {
	statement, err := database.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(id)
	return err
}
