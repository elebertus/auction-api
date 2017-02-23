package main

import (
	"fmt"
	"net/http"

	"github.com/elebertus/auction-api/routes"
)

func main() {

	fmt.Printf("Serving on http://localhost:8080\n")

	router := routes.NewRouter()

	http.ListenAndServe(":8080", router)
}
