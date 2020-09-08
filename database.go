package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
	"time"

)

const (
	host     = "localhost"
	port     = 5432
	user     = "golang"
	password = "golang"
	dbname   = "golang"
)

func connectToDatabase() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to postgres!")

	sqlStatement := `

		CREATE TABLE IF NOT EXISTS ARTICLE (
			id serial, 
			date_start date not null,
			date_end date not null,
			libelle varchar(255) not null,
			description text,
			immediate_price float not null,
			start_price float not null,
			current_price float
		)`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	fmt.Println("tables created!")

	for i := 0; i < 10; i++ {
		query := `
			INSERT INTO article (id, date_start, date_end, libelle, description, immediate_price, start_price)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`
		_, err = db.Exec(query,
			i,
			time.Date(2020, time.Month(2), 21, 1, 10, 30, 0, time.UTC),
			time.Date(2020, time.Month(2), 23, 1, 10, 30, 0, time.UTC),
			"article-" + strconv.Itoa(i) ,
			"ceci est une description",
			3000,
			1200,
		)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("fixtures loaded")
}