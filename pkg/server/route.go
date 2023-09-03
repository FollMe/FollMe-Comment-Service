package server

import (
	"follme/comment-service/pkg/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func Route() {
	app := NewApp()

	router := mux.NewRouter().StrictSlash(true)
	protectedRouter := router.NewRoute().Subrouter()
	const BaseUrl = "/comment-svc/api"

	router.Use(middleware.FilterPanicMiddleware)
	protectedRouter.Use(middleware.AuthenticationMiddleware)

	router.HandleFunc(BaseUrl+"/comments/{postId}", app.CmtHandler.GetCommentsOfPost).Methods("GET")
	protectedRouter.HandleFunc(BaseUrl+"/comments", app.CmtHandler.CreateCommentsOfPost).Methods("POST")
	router.HandleFunc(BaseUrl+"/comments/count", app.CmtHandler.GetNumberCommentsOfPosts).Methods("POST")
	router.HandleFunc("/comment-svc/ws", app.CmtHandler.StartWSConnection).Methods("GET")

	// Commit date
	router.HandleFunc(BaseUrl+"/commit-date/{id}", app.CommitDateHandler.GetCommitDate).Methods("GET")
	router.HandleFunc(BaseUrl+"/commit-date/{id}", app.CommitDateHandler.UpdateCommitDate).Methods("POST")

	http.ListenAndServe(":3000", router)
}
