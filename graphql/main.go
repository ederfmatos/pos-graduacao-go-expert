package main

import (
	"database/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"pos-graduacao-go-lang/graphql/graph"
	"pos-graduacao-go-lang/graphql/internal/repository"
	"pos-graduacao-go-lang/graphql/internal/usecase/create_category"
	"pos-graduacao-go-lang/graphql/internal/usecase/create_product"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	database, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer database.Close()

	categoryRepository := repository.NewSqlCategoryRepository(database)
	productRepository := repository.NewSqlProductRepository(database)

	resolvers := &graph.Resolver{
		CreateCategoryUseCase: create_category.New(categoryRepository),
		CreateProductUseCase:  create_product.New(productRepository),
		ProductRepository:     productRepository,
		CategoryRepository:    categoryRepository,
	}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolvers}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
