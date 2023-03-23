package handler

import (
	"encoding/json"
	"follme/comment-service/pkg/adapter/serializer"
	"follme/comment-service/pkg/model"
	"net/http"

	"github.com/gorilla/mux"
)

type CmtHandler struct {
	cmtSvc model.CommentSvc
}

func NewCmtHandler(c model.CommentSvc) *CmtHandler {
	return &CmtHandler{
		cmtSvc: c,
	}
}

func (h CmtHandler) GetCommentsOfPost(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["postId"]
	cmts, err := h.cmtSvc.GetCommentsOfPost(r.Context(), postId)
	if err != nil {
		json.NewEncoder(w).Encode(
			serializer.NewFailHttpRes(""),
		)
	}

	cmtsRes := []serializer.Comment{}
	for _, cmt := range cmts {
		cmtRes := serializer.Comment{
			ID:        cmt.ID(),
			PostID:    cmt.PostID(),
			Author:    cmt.Author(),
			Content:   cmt.Content(),
			CreatedAt: *cmt.CreatedAt(),
			UpdatedAt: cmt.UpdatedAt(),
		}
		replyCmtRes := []serializer.Comment{}
		for _, replyCmt := range cmt.Replies() {
			replyCmtRes = append(replyCmtRes, serializer.Comment{
				ID:        replyCmt.ID(),
				Author:    replyCmt.Author(),
				Content:   replyCmt.Content(),
				CreatedAt: *replyCmt.CreatedAt(),
				UpdatedAt: replyCmt.UpdatedAt(),
			})
		}
		cmtRes.Replies = replyCmtRes

		cmtsRes = append(cmtsRes, cmtRes)
	}

	dataRes := serializer.GetCommentsOfPostRes{
		Comments: cmtsRes,
	}
	httpRes, err := json.Marshal(serializer.NewSuccessHttpRes("", dataRes))
	if err != nil {
		json.NewEncoder(w).Encode(
			serializer.NewFailHttpRes(""),
		)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(httpRes)
}
