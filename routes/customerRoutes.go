package routes

import (
	"net/http"
	controller "trial-go/controllers"
)

func CustomerRoutes() {
	http.HandleFunc("/list-customers", controller.GetCustomers)
}
