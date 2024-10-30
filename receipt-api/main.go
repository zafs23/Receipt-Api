package main

import (
	"fmt"
	"log"
	"net/http"
	"receipt-api/routes"
)

func main() {
	r := routes.RegisterRoutes()
	fmt.Println("Server listening on local port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
