package server

import (
	"follme/comment-service/pkg/adapter/database"
	"follme/comment-service/pkg/adapter/database/repository"
	"follme/comment-service/pkg/adapter/handler"
	"follme/comment-service/pkg/service"
)

type ApplicationContext struct {
	CmtHandler *handler.CmtHandler
}

func NewApp() *ApplicationContext {
	db := database.ConnectDB()
	cmtRepo := repository.NewCommentRepo(db)
	wsSvc := service.NewWebSocketService()
	cmtSvc := service.NewCommentSvc(cmtRepo, wsSvc)
	cmtHandler := handler.NewCmtHandler(cmtSvc, wsSvc)

	return &ApplicationContext{
		CmtHandler: cmtHandler,
	}
}
