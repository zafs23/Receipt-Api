package main

import (
	"fmt"
	"net/http"
	"receipt-api/routes"
)

func main() {
	r := routes.SetupRouter()
	fmt.Println("Server listening on local port 8080")
	http.ListenAndServe(":8080", r)
}
