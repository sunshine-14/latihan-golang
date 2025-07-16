package main

import (
	"fmt"
	"net/http"
	"trial-go/routes"
)

func main() {
	routes.CustomerRoutes()
	routes.UserRoutes()

	fmt.Println("Server jalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
