package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

var RouteManager *mux.Router

func init() {
	RouteManager = mux.NewRouter()
	http.Handle("/", RouteManager)
}

// RegisterAPI :: 註冊路由
func RegisterAPI(path string, action func(http.ResponseWriter, *http.Request), methods string) {
	RouteManager.HandleFunc(path, action).Methods(methods)
}
