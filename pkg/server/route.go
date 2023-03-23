package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Route() {
	app := NewApp()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/comments/{postId}", app.CmtHandler.GetCommentsOfPost).Methods("GET")
	http.ListenAndServe(":3002", router)
}
