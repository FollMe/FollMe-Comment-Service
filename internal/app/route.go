package server

import (
	cmt_service_http "follme/comment-service/internal/comment_service/http"
	story_with_you_http "follme/comment-service/internal/story_with_you/http"
	"follme/comment-service/pkg/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func Route() {
	app := NewApp()

	router := mux.NewRouter().StrictSlash(true).PathPrefix("/comment-svc").Subrouter()
	protectedRouter := router.NewRoute().Subrouter()

	router.Use(middleware.FilterPanicMiddleware)
	protectedRouter.Use(middleware.AuthenticationMiddleware)

	// Comment service
	cmt_service_http.Route(router, protectedRouter, app.CmtHandler)

	// Commit date
	story_with_you_http.Route(router, protectedRouter, app.StoryWithYouHandler)

	http.ListenAndServe(":3000", router)
}
