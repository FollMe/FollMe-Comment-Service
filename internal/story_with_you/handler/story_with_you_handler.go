package handler

import (
	"encoding/json"
	"follme/comment-service/internal/story_with_you/domain"
	"follme/comment-service/pkg/adapter/serializer"
	"net/http"

	"github.com/gorilla/mux"
)

type CommitDateHandler struct {
	commitDateSvc domain.CommitDateSvc
}

func NewStoryWithYouHandler(commitDateSvc domain.CommitDateSvc) *CommitDateHandler {
	return &CommitDateHandler{
		commitDateSvc: commitDateSvc,
	}
}

func (h CommitDateHandler) GetCommitDate(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	commitDate, err := h.commitDateSvc.GetCommitDate(r.Context(), id)
	if err != nil {
		panic(err)
	}

	serializer.WriteSuccessResponse(w, "", commitDate)
}

func (h CommitDateHandler) UpdateCommitDate(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var req domain.UpdateCommitDateReq
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	err = h.commitDateSvc.UpdateCommitDate(r.Context(), id, req.Date)
	if err != nil {
		panic(err)
	}

	serializer.WriteSuccessResponse(w, "", nil)
}
