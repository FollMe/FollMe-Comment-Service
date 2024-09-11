package server

import (
	"follme/comment-service/internal/comment_service"
	"follme/comment-service/internal/story_with_you"
	"follme/comment-service/pkg/adapter/database"
	"follme/comment-service/pkg/config"
)

type ApplicationContext struct {
	CmtHandler          comment_service.CmtHandler
	StoryWithYouHandler story_with_you.StoryWithYouHandler
}

func NewApp() *ApplicationContext {
	db := database.ConnectDB()
	cmtHandler := comment_service.NewCmtHandler(db, config.AppConfig.WSToken)
	storyWithYouHandler := story_with_you.NewStoryWithYouHandler(db)

	return &ApplicationContext{
		CmtHandler:          cmtHandler,
		StoryWithYouHandler: storyWithYouHandler,
	}
}
