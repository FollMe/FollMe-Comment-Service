package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Route() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Tui oke nha")
	}).Methods("GET")
	http.ListenAndServe(":3000", router)
}
