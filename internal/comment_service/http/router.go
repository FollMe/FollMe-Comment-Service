package http

import (
	"follme/comment-service/internal/comment_service"

	"github.com/gorilla/mux"
)

func Route(router *mux.Router, protectedRouter *mux.Router, cmtHandler comment_service.CmtHandler) {

	router.HandleFunc("/api/comments/{postId}", cmtHandler.GetCommentsOfPost).Methods("GET")
	protectedRouter.HandleFunc("/api/comments", cmtHandler.CreateCommentsOfPost).Methods("POST")
	router.HandleFunc("/api/comments/count", cmtHandler.GetNumberCommentsOfPosts).Methods("POST")
	router.HandleFunc("/ws", cmtHandler.StartWSConnection).Methods("GET")

}
