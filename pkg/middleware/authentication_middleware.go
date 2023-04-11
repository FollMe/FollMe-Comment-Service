package middleware

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"follme/comment-service/pkg/adapter/serializer"
	"follme/comment-service/pkg/model"

	"net/http"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		userInfos := req.Header["X-User-Info"]
		w.Header().Set("Content-Type", "application/json")
		if len(userInfos) <= 0 {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(
				serializer.NewFailHttpRes("Vui lòng đăng nhập"),
			)
			return
		}

		decodeInfo, _ := base64.StdEncoding.DecodeString(userInfos[0])
		user := model.User{}
		err := json.Unmarshal(decodeInfo, &user)
		if err != nil {
			panic(err)
		}

		reqContext := context.WithValue(req.Context(), "UserInfo", &user)
		req = req.WithContext(reqContext)
		next.ServeHTTP(w, req)
	})
}
