package comment_service

import (
	"database/sql"
	"follme/comment-service/internal/comment_service/handler"
	"follme/comment-service/internal/comment_service/repository"
	"follme/comment-service/internal/comment_service/service"
	"net/http"
)

type CmtHandler interface {
	GetCommentsOfPost(w http.ResponseWriter, r *http.Request)
	CreateCommentsOfPost(w http.ResponseWriter, r *http.Request)
	GetNumberCommentsOfPosts(w http.ResponseWriter, r *http.Request)
	StartWSConnection(w http.ResponseWriter, r *http.Request)
}

func NewCmtHandler(db *sql.DB, wsToken string) CmtHandler {
	wsSvc := service.NewWebSocketService(wsToken)
	cmtRepo := repository.NewCommentRepo(db)
	cmtSvc := service.NewCommentSvc(cmtRepo, wsSvc)

	return handler.NewCmtHandler(cmtSvc, wsSvc)
}
