package model

// CreateContent :: 建立會員請求資料
type CreateContent struct {
	Account  string
	Password string
}

// CreateResult :: 建立會員結果
type CreateResult struct {
	IsOK bool
}

// DeleteContent :: 刪除會員請求資料
type DeleteContent struct {
	Account string
}

// DeleteResult :: 刪除會員結果
type DeleteResult struct {
	IsOK bool
}

// ModifyPasswordContent :: 修改會員密碼請求資料
type ModifyPasswordContent struct {
	Account  string
	Password string
}

// ModifyPasswordResult :: 修改會員密碼結果
type ModifyPasswordResult struct {
	IsOK bool
}

// VerificationContent :: 驗證帳號密碼請求資料
type VerificationContent struct {
	Account  string
	Password string
}

// UserDao :: 會員資料
type UserDao struct {
	Account  string
	Password string
}
