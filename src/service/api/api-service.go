package service

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	apiModel "../../models/api"
)

// RequestHandler :: Api 請求處理器
func RequestHandler(request *http.Request, data interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(request.Body, 1024))
	if err == nil {
		json.Unmarshal(body, &data)
	}

	defer request.Body.Close()
	return err
}

// ResponseHandler :: Api 回應處理器
func ResponseHandler(write http.ResponseWriter, status int, response apiModel.ResponseDto) {
	data, err := json.Marshal(response)
	write.Header().Set("Content-Type", "application/json")
	if err == nil {
		write.WriteHeader(status)
		write.Write(data)
	} else {
		write.WriteHeader(http.StatusInternalServerError)
		write.Write([]byte("Error"))
	}
}
