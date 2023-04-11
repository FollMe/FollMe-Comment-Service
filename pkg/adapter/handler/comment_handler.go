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
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			serializer.NewFailHttpRes(""),
		)
		return
	}

	cmtsRes := []serializer.Comment{}
	for _, cmt := range cmts {
		cmtRes := serializer.Comment{
			ID:        cmt.ID(),
			PostSlug:  cmt.PostSlug(),
			Author:    cmt.Author(),
			Content:   cmt.Content(),
			CreatedAt: cmt.CreatedAt(),
			UpdatedAt: cmt.UpdatedAt(),
		}
		replyCmtRes := []serializer.Comment{}
		for _, replyCmt := range cmt.Replies() {
			replyCmtRes = append(replyCmtRes, serializer.Comment{
				ID:        replyCmt.ID(),
				Author:    replyCmt.Author(),
				Content:   replyCmt.Content(),
				CreatedAt: replyCmt.CreatedAt(),
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
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			serializer.NewFailHttpRes(""),
		)
		return
	}

	w.Write(httpRes)
}

func (h CmtHandler) CreateCommentsOfPost(w http.ResponseWriter, r *http.Request) {
	var req serializer.CreateCommentOfPostReq
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			serializer.NewFailHttpRes(""),
		)
		return
	}

	err = serializer.Validate(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			serializer.NewFailHttpRes("Invalid request body"),
		)
		return
	}

	user := r.Context().Value("UserInfo").(*model.User)

	_, err = h.cmtSvc.InsertCommentOfPost(r.Context(), model.CreateCommentOpts{
		PostSlug: req.PostSlug,
		ParentId: req.ParentId,
		Content:  req.Content,
		Author:   user.ID,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			serializer.NewFailHttpRes(""),
		)
		return
	}

	json.NewEncoder(w).Encode(
		serializer.NewSuccessHttpRes("Post comment successful", nil),
	)
}
