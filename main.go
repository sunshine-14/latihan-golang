package main

import (
	"fmt"
	"net/http"
	"os"
	"trial-go/routes"
)

func main() {
	routes.CustomerRoutes()
	routes.UserRoutes()

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	fmt.Println("Server jalan di", addr)
	http.ListenAndServe(addr, nil)
}
