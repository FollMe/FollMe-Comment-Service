package server

import (
	"follme/comment-service/pkg/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func Route() {
	app := NewApp()

	router := mux.NewRouter().StrictSlash(true)
	const BaseUrl = "/comment-svc/api"
	router.Use(middleware.AuthenticationMiddleware)
	router.Use(middleware.FilterPanicMiddleware)
	router.HandleFunc(BaseUrl+"/comments/{postId}", app.CmtHandler.GetCommentsOfPost).Methods("GET")
	router.HandleFunc(BaseUrl+"/comments", app.CmtHandler.CreateCommentsOfPost).Methods("POST")
	router.HandleFunc("/comment-svc/ws", app.CmtHandler.StartWSConnection).Methods("GET")

	// Commit date
	router.HandleFunc(BaseUrl+"/commit-date/{id}", app.CommitDateHandler.GetCommitDate).Methods("GET")

	http.ListenAndServe(":3000", router)
}
