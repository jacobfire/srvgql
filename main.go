package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"graphql/book"
	"graphql/infrastructure"
	"log"
	"net/http"
	"net/url"
)

func main() {
	routes := chi.NewRouter()
	r := book.RegisterRoutes(routes)
	log.Println("Server ready at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func init() {
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	env := infrastructure.Environment{}
	env.SetEnvironment()
	env.LoadConfig()
	_, err := env.InitPostgres()
	if err != nil {
		fmt.Println("Init DB error", err)
	}
}
