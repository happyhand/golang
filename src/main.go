package main

import (
	"net/http"

	_ "./controllers/user"
	route "./route"
)

func main() {
	http.ListenAndServe(":3000", route.RouteManager)
}
