package middleware

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"follme/comment-service/pkg/adapter/serializer"
	"follme/comment-service/pkg/user"

	"net/http"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/comment-svc/ws" {
			next.ServeHTTP(w, req)
		}
		userInfos := req.Header["X-User-Info"]
		if len(userInfos) <= 0 {
			serializer.ResponseJSON(w)(serializer.NewFailHttpRes("Vui lòng đăng nhập"), http.StatusUnauthorized)
			return
		}

		decodeInfo, _ := base64.StdEncoding.DecodeString(userInfos[0])
		user := user.User{}
		err := json.Unmarshal(decodeInfo, &user)
		if err != nil {
			panic(err)
		}

		reqContext := context.WithValue(req.Context(), "UserInfo", &user)
		req = req.WithContext(reqContext)
		next.ServeHTTP(w, req)
	})
}
