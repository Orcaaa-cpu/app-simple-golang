package config

import (
	"os"
	"product/entities"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var db *pg.DB

func CreateCon() *pg.DB {
	if db != nil {
		return db
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	// dbUsername := "postgres"
	// dbPassword := "12345678"
	// dbHost := "localhost"
	// dbPort := "5432"
	// dbName := "postgres"

	opt := &pg.Options{
		User:     dbUsername,
		Password: dbPassword,
		Addr:     dbHost + ":" + dbPort,
		Database: dbName,
	}

	db = pg.Connect(opt)

	err := db.Model((*entities.Users)(nil)).CreateTable(&orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})
	if err != nil {
		panic(err)
	}

	err = db.Model((*entities.Pasiens)(nil)).CreateTable(&orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})
	if err != nil {
		panic(err)
	}

	return db
}
