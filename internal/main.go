package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go-project/internal/domain/entity"
	"go-project/internal/infra/database"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=go-project sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	productRepository := database.NewProductRepositoryImpl(db)
	product := entity.NewProduct("foo", 23.3)
	err = productRepository.Create(product)
	product2 := entity.NewProduct("foo2", 232.32)
	err = productRepository.Create(product2)
	product3 := entity.NewProduct("foo3", 2.28)
	err = productRepository.Create(product3)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
	product.Price = 100.00
	product.Name = "Notebook"
	err = productRepository.Update(product)
	if err != nil {
		panic(err)
	}
	err = productRepository.Delete(product3.ID)
	if err != nil {
		panic(err)
	}
	findedProduct, err := productRepository.FindById(product3.ID)
	if err != nil {
		panic(err)
	}
	productList, err := productRepository.FindAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(findedProduct)
	fmt.Println("-----------------------------------------------")
	fmt.Println(productList)
}
