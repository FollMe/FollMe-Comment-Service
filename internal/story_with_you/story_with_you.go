package story_with_you

import (
	"database/sql"
	"follme/comment-service/internal/story_with_you/handler"
	"follme/comment-service/internal/story_with_you/repository"
	"follme/comment-service/internal/story_with_you/service"
	"net/http"
)

type StoryWithYouHandler interface {
	GetCommitDate(w http.ResponseWriter, r *http.Request)
	UpdateCommitDate(w http.ResponseWriter, r *http.Request)
}

func NewStoryWithYouHandler(db *sql.DB) StoryWithYouHandler {
	commitDateRepo := repository.NewCommitDateRepo(db)
	commitDateSvc := service.NewCommitDateSvc(commitDateRepo)

	return handler.NewStoryWithYouHandler(commitDateSvc)
}
