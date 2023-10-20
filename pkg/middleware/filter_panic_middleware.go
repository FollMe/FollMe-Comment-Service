package middleware

import (
	"encoding/json"
	"follme/comment-service/pkg/adapter/serializer"
	"log"

	"net/http"
)

func FilterPanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(
					serializer.NewFailHttpRes(""),
				)
			}
		}()
		next.ServeHTTP(w, req)
	})
}
