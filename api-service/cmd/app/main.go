package main

import (
	"log"
	"net/http"

	"github.com/kokweikhong/msia-stock-app/api-service/internal/router"
)

func main() {

	r := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", r))

    // klsescreener.GetQuoteResults()

}

