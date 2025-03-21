package serializer

import (
	"encoding/json"
	"net/http"

	validator "github.com/go-playground/validator/v10"
)

func Validate(obj interface{}) error {
	validate := validator.New()
	return validate.Struct(obj)
}

func ValidateOrPanic(w http.ResponseWriter, obj interface{}) {
	err := Validate(obj)

	if err != nil {
		ResponseJSON(w)(
			NewFailHttpRes(err.Error()),
			http.StatusBadRequest,
		)
		panic(nil)
	}
}

func ResponseJSON(w http.ResponseWriter) func(body interface{}, statusCode_optional ...int) {
	return func(body interface{}, statusCode_optional ...int) {
		statusCode := http.StatusOK
		if len(statusCode_optional) > 0 {
			statusCode = statusCode_optional[0]
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(body)
	}
}

func WriteSuccessResponse(w http.ResponseWriter, message string, data interface{}, statusCode_optional ...int) {
	ResponseJSON(w)(NewSuccessHttpRes(message, data), statusCode_optional...)
}
