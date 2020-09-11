package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var (
	db *sql.DB
)

type Article struct {
	ID        int    `json:"id"`
	Libelle string `json:"libelle"`
	StartPrice int `json:"start_price"`
	CurrentPrice int `json:"current_price"`
}

type UpdateArticle struct {
	 id int
	 currentPrice int
}

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var articles []Article

	result, err := db.Query("SELECT * from article")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var article Article
		err := result.Scan(&article.ID, &article.Libelle, &article.StartPrice, &article.CurrentPrice)
		if err != nil {
			panic(err.Error())
		}
		articles = append(articles, article)
	}
	json.NewEncoder(w).Encode(articles)
}

func newArticle(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT INTO article(id,libelle, start_price, current_price) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	id := keyVal["id"]
	libelle := keyVal["libelle"]
	startPrice := keyVal["startprice"]
	currentPrice := keyVal["currentprice"]
	_, err = stmt.Exec(id, libelle,  startPrice, currentPrice)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New article was created")
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM article WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Article with ID = %s was deleted", params["id"])
}

func getArticleHandler(id int) Article  {
	result, err := db.Query("SELECT * FROM article WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var article Article
	for result.Next() {
		err := result.Scan(&article.ID, &article.Libelle, &article.StartPrice, &article.CurrentPrice)
		if err != nil {
			panic(err.Error())
		}
	}
	return article
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	article := getArticleHandler(id)
	json.NewEncoder(w).Encode(article)
}

func updateArticle(channel chan UpdateArticle, chanOut chan string) func (w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err.Error())
		}
		keyVal := make(map[string]string)
		json.Unmarshal(body, &keyVal)
		newPrice, _ := strconv.Atoi(keyVal["currentPrice"])
		var data UpdateArticle
		data.id = id
		data.currentPrice = newPrice
		channel <- data
		res := <- chanOut
		fmt.Println(res)
		json.NewEncoder(w).Encode(res)
	}
}

func manageUpdate(chanIn chan UpdateArticle, chanOut chan string) {
	for {
		request := <- chanIn
		article := getArticleHandler(request.id)
		fmt.Println(article.CurrentPrice)
		if request.currentPrice <= article.CurrentPrice {
			chanOut <- "Error, price lower than the current one"
		} else {
			stmt, err := db.Prepare("UPDATE article SET current_price = $1 WHERE id = $2")
			if err != nil {
				panic(err.Error())
			}
			_, err = stmt.Exec(request.currentPrice, request.id)
			if err != nil {
				panic(err.Error())
			}
			chanOut <- "Article was updated"
		}

	}
}

func main() {
	/*launchRoutine()*/
	channelInt, channelOut := make(chan UpdateArticle), make(chan string)
	go manageUpdate(channelInt, channelOut)
	db = connectToDatabase()
	r := mux.NewRouter()
	api := r.PathPrefix("/api/article").Subrouter()
	api.HandleFunc("", getAll).Methods(http.MethodGet)
	api.HandleFunc("", newArticle).Methods(http.MethodPost)
	api.HandleFunc("/{id}", updateArticle(channelInt, channelOut)).Methods(http.MethodPost)
	api.HandleFunc("/{id}", deleteArticle).Methods(http.MethodDelete)
	api.HandleFunc("/{id}", getArticle).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))
}
