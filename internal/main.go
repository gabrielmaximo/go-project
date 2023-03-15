package main

import (
	"database/sql"
	"fmt"
	"github.com/gabrielmaximo/go-project/internal/domain/entity"
	"github.com/gabrielmaximo/go-project/internal/infra/database"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	sqlDB, err := sql.Open("postgres", "user=postgres password=postgres dbname=go-project sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(sqlDB)

	ormDB, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN: "host=localhost user=postgres password=postgres dbname=go-project port=5432 sslmode=disable",
			},
		),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(ormDB.Config.Dialector.Name())

	productRepository := database.NewProductRepositoryImpl(sqlDB)
	product := entity.NewProduct("Notebook", 499.90)
	product2 := entity.NewProduct("Monitor", 299.90)
	product3 := entity.NewProduct("Mouse", 99.90)
	err = productRepository.Create(product)
	if err != nil {
		panic(err)
	}
	err = productRepository.Create(product2)
	if err != nil {
		panic(err)
	}
	err = productRepository.Create(product3)
	if err != nil {
		panic(err)
	}
	products, err := productRepository.FindAll()

	fmt.Println(*products)
}
