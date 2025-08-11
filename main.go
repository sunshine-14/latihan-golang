package main

import (
	"fmt"
	"net/http"
	"trial-go/routes"
)

func main() {
	routes.CustomerRoutes()
	routes.UserRoutes()

	addr := ":8080"
	fmt.Println("Server jalan di", addr)
	http.ListenAndServe(addr, nil)
}
