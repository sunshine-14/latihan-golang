package routes

import (
	"net/http"
	controller "trial-go/controllers"
)

func UserRoutes() {
	http.HandleFunc("/list-users", controller.ListUser)
}
