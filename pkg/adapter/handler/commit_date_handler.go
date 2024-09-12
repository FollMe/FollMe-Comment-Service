package handler

import (
	"encoding/json"
	"follme/comment-service/pkg/adapter/serializer"
	"follme/comment-service/pkg/model"
	"net/http"

	"github.com/gorilla/mux"
)

type CommitDateHandler struct {
	commitDateSvc model.CommitDateSvc
}

func NewCommitDateHandler(c model.CommitDateSvc) *CommitDateHandler {
	return &CommitDateHandler{
		commitDateSvc: c,
	}
}

func (h CommitDateHandler) GetCommitDate(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	commitDate, err := h.commitDateSvc.GetCommitDate(r.Context(), id)
	if err != nil {
		panic(err)
	}

	serializer.ResponseJSON(w)(serializer.NewSuccessHttpRes("", commitDate))
}

func (h CommitDateHandler) UpdateCommitDate(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var req serializer.UpdateCommitDateReq
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	err = h.commitDateSvc.UpdateCommitDate(r.Context(), id, req.Date)
	if err != nil {
		panic(err)
	}

	serializer.ResponseJSON(w)(serializer.NewSuccessHttpRes("", nil))
}
