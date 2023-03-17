package main

import (
	"database/sql"
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
	_ = database.NewProductRepositoryGormImpl(ormDB)

}
