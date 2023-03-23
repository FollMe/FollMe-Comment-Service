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
	cmtSvc := service.NewCommentSvc(cmtRepo)
	cmtHandler := handler.NewCmtHandler(cmtSvc)

	return &ApplicationContext{
		CmtHandler: cmtHandler,
	}
}
