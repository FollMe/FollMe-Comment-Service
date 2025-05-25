package http

import (
	"follme/comment-service/internal/story_with_you"

	"github.com/gorilla/mux"
)

func Route(router *mux.Router, protectedRouter *mux.Router, storyWithYouHandler story_with_you.StoryWithYouHandler) {

	router.HandleFunc("/api/commit-date/{id}", storyWithYouHandler.GetCommitDate).Methods("GET")
	router.HandleFunc("/api/commit-date/{id}", storyWithYouHandler.UpdateCommitDate).Methods("POST")
}
