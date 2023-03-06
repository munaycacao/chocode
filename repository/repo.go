package repository

import (
	"database/sql"
	"os"
	"fmt"
)

func GetDBConnString() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)
}

type Repo struct {
	db *sql.DB
	Queries *Queries
}

func NewRepo() *Repo {
	fmt.Println(GetDBConnString())
	db, err := sql.Open("mysql", GetDBConnString())
	if err != nil {
		panic(err.Error())
	}
	return &Repo{
		db: db,
		Queries: New(db),
	}
}