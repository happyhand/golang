package model

// ResponseDto :: Api 回應資料
type ResponseDto struct {
	httpStatus int
	Code       int
	Message    string
	Result     interface{}
}
