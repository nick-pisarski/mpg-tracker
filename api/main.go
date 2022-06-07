package main

import (
	"log"
	"mpg-tracker/api/router"
	"net/http"
)

func main() {
	fillUpRouter := router.NewRouter(router.Config{
		ConnectionString: SQLiteConnectionString,
	})

	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", fillUpRouter))
}
