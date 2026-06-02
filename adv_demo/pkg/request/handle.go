package request

import (
	"adv_demo/pkg/response"
	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := DecodeRequest[T](r.Body)
	if err != nil {
		response.PlainText(*w, 400, err.Error())
		return nil, err
	}
	err = IsValid(body)
	if err != nil {
		response.PlainText(*w, 400, err.Error())
		return nil, err
	}
	return &body, nil
}
