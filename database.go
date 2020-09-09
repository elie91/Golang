package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "golang"
	password = "golang"
	dbname   = "golang"
)

func connectToDatabase() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to postgres!")

	sqlStatement := `
		CREATE TABLE IF NOT EXISTS ARTICLE (
			id serial, 
			libelle varchar(255) not null,
			start_price float not null,
			current_price float
		)`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	fmt.Println("tables created!")

	_, _ = db.Exec("TRUNCATE article")

	for i := 0; i < 10; i++ {
		query := `
			INSERT INTO article (id, libelle, start_price, current_price)
			VALUES ($1, $2, $3, $4)
		`
		_, err = db.Exec(query,
			i,
			"article-" + strconv.Itoa(i) ,
			1200,
			1200,
		)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("fixtures loaded")

	return db
}