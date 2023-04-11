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
	router.HandleFunc(BaseUrl+"/comments/{postId}", app.CmtHandler.GetCommentsOfPost).Methods("GET")
	router.HandleFunc(BaseUrl+"/comments", app.CmtHandler.CreateCommentsOfPost).Methods("POST")
	http.ListenAndServe(":3001", router)
}
