package middleware

import (
	"follme/comment-service/pkg/adapter/serializer"
	"log"

	"net/http"
)

func FilterPanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				serializer.ResponseJSON(w)(serializer.NewFailHttpRes(""), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, req)
	})
}
