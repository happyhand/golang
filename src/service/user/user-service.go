package service

import (
	"net/http"

	apiModel "../../models/api"
	userModel "../../models/user"
	userRepository "../../repository/user"
)

// Create :: 新增會員
func Create(content userModel.CreateContent) (int, apiModel.ResponseDto) {
	if content.Account == "" || content.Password == "" {
		return http.StatusBadRequest, apiModel.ResponseDto{Code: 1, Message: "Invalid Input", Result: userModel.CreateResult{IsOK: false}}
	}

	user := userRepository.Get(content.Account)
	if user != nil {
		return http.StatusConflict, apiModel.ResponseDto{Code: 2, Message: "Reject Create", Result: userModel.CreateResult{IsOK: false}}
	}

	result := userRepository.Create(content.Account, content.Password)
	if !result {
		return http.StatusConflict, apiModel.ResponseDto{Code: 2, Message: "Reject Create", Result: userModel.CreateResult{IsOK: false}}
	}

	return http.StatusOK, apiModel.ResponseDto{Code: 0, Message: "", Result: userModel.CreateResult{IsOK: true}}
}

// Delete :: 刪除會員
func Delete(content userModel.DeleteContent) (int, apiModel.ResponseDto) {
	if content.Account == "" {
		return http.StatusBadRequest, apiModel.ResponseDto{Code: 1, Message: "Invalid Input", Result: userModel.DeleteResult{IsOK: false}}
	}

	user := userRepository.Get(content.Account)
	if user == nil {
		return http.StatusConflict, apiModel.ResponseDto{Code: 2, Message: "Reject Create", Result: userModel.DeleteResult{IsOK: false}}
	}

	result := userRepository.Delete(content.Account)
	if !result {
		return http.StatusConflict, apiModel.ResponseDto{Code: 2, Message: "Reject Delete", Result: userModel.DeleteResult{IsOK: false}}
	}

	return http.StatusOK, apiModel.ResponseDto{Code: 0, Message: "", Result: userModel.DeleteResult{IsOK: true}}
}

// ModifyPassword :: 修改會員密碼
func ModifyPassword(content userModel.ModifyPasswordContent) (int, apiModel.ResponseDto) {
	if content.Account == "" || content.Password == "" {
		return http.StatusBadRequest, apiModel.ResponseDto{Code: 1, Message: "Invalid Input", Result: userModel.ModifyPasswordResult{IsOK: false}}
	}

	user := userRepository.Get(content.Account)
	if user == nil {
		return http.StatusConflict, apiModel.ResponseDto{Code: 2, Message: "Reject Modify Password", Result: userModel.ModifyPasswordResult{IsOK: false}}
	}

	result := userRepository.ModifyPassword(content.Account, content.Password)
	if !result {
		return http.StatusConflict, apiModel.ResponseDto{Code: 2, Message: "Reject Modify Password", Result: userModel.ModifyPasswordResult{IsOK: false}}
	}

	return http.StatusOK, apiModel.ResponseDto{Code: 0, Message: "", Result: userModel.ModifyPasswordResult{IsOK: true}}
}

// Verification :: 驗證帳號密碼
func Verification(content userModel.VerificationContent) (int, apiModel.ResponseDto) {
	if content.Account == "" || content.Password == "" {
		return http.StatusBadRequest, apiModel.ResponseDto{Code: 2, Message: "Login Failed", Result: nil}
	}

	user := userRepository.Get(content.Account)
	if user == nil || user.Password != content.Password {
		return http.StatusBadRequest, apiModel.ResponseDto{Code: 2, Message: "Login Failed", Result: nil}
	}

	return http.StatusOK, apiModel.ResponseDto{Code: 0, Message: "", Result: nil}
}
