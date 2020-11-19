package controller

import (
	"net/http"

	userModel "../../models/user"
	route "../../route"
	apiService "../../service/api"
	userService "../../service/user"
	"github.com/gorilla/mux"
)

func init() {
	route.RegisterAPI("/v1/user/create", create, http.MethodPost)
	route.RegisterAPI("/v1/user/delete", delete, http.MethodPost)
	route.RegisterAPI("/v1/user/pwd/change", modifyPassword, http.MethodPost)
	route.RegisterAPI("/v1/user/login/{Account}/{Password}", verification, http.MethodGet)
}

// create :: API 接口 - 新增會員
func create(writer http.ResponseWriter, request *http.Request) {
	var content userModel.CreateContent
	apiService.RequestHandler(request, &content)
	httpStatus, response := userService.Create(content)
	apiService.ResponseHandler(writer, httpStatus, response)
}

// delete :: API 接口 - 刪除會員
func delete(writer http.ResponseWriter, request *http.Request) {
	var content userModel.DeleteContent
	apiService.RequestHandler(request, &content)
	httpStatus, response := userService.Delete(content)
	apiService.ResponseHandler(writer, httpStatus, response)
}

// modifyPassword :: API 接口 - 修改會員密碼
func modifyPassword(writer http.ResponseWriter, request *http.Request) {
	var content userModel.ModifyPasswordContent
	apiService.RequestHandler(request, &content)
	httpStatus, response := userService.ModifyPassword(content)
	apiService.ResponseHandler(writer, httpStatus, response)
}

// modifyPassword :: API 接口 - 修改會員密碼
func verification(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	httpStatus, response := userService.Verification(userModel.VerificationContent{Account: vars["Account"], Password: vars["Password"]})
	apiService.ResponseHandler(writer, httpStatus, response)
}
