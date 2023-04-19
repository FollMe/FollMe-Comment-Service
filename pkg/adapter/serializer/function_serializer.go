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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			NewFailHttpRes(err.Error()),
		)
		panic(nil)
	}
}
