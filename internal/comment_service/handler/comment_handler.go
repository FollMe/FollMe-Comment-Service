package handler

import (
	"encoding/json"
	"follme/comment-service/internal/comment_service/domain"
	"follme/comment-service/pkg/adapter/serializer"
	"follme/comment-service/pkg/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type CmtHandler struct {
	cmtSvc domain.CommentSvc
	wsSvc  domain.WebSocketSvc
}

func NewCmtHandler(c domain.CommentSvc, ws domain.WebSocketSvc) *CmtHandler {
	return &CmtHandler{
		cmtSvc: c,
		wsSvc:  ws,
	}
}

var upgrader = websocket.Upgrader{}

func (h CmtHandler) GetCommentsOfPost(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["postId"]
	cmts, err := h.cmtSvc.GetCommentsOfPost(r.Context(), postId)
	if err != nil {
		panic(err)
	}

	cmtsRes := []domain.CommentRes{}
	for _, cmt := range cmts {
		cmtRes := domain.CommentRes{
			ID:        cmt.ID(),
			PostSlug:  cmt.PostSlug(),
			Author:    cmt.Author(),
			Content:   cmt.Content(),
			CreatedAt: cmt.CreatedAt(),
			UpdatedAt: cmt.UpdatedAt(),
		}
		replyCmtRes := []domain.CommentRes{}
		for _, replyCmt := range cmt.Replies() {
			replyCmtRes = append(replyCmtRes, domain.CommentRes{
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

	json.NewEncoder(w).Encode(
		serializer.NewSuccessHttpRes("", domain.GetCommentsOfPostRes{
			Comments: cmtsRes,
		}),
	)
}

func (h CmtHandler) CreateCommentsOfPost(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateCommentOfPostReq
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	serializer.ValidateOrPanic(w, req)

	user := r.Context().Value("UserInfo").(*user.User)

	cmt, err := h.cmtSvc.InsertCommentOfPost(r.Context(), domain.CreateCommentOpts{
		PostSlug: req.PostSlug,
		ParentId: req.ParentId,
		Content:  req.Content,
		Author:   user.ID,
	})

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(
		serializer.NewSuccessHttpRes("Post comment successful", domain.CreateCommentOfPostRes{
			ID: cmt.ID(),
		}),
	)
}

func (h CmtHandler) GetNumberCommentsOfPosts(w http.ResponseWriter, r *http.Request) {
	var req domain.GetNumberCommentsOfPostsReq
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	serializer.ValidateOrPanic(w, req)

	result, err := h.cmtSvc.GetNumberCommentsOfPosts(r.Context(), req.PostSlugs)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(
		serializer.NewSuccessHttpRes("", map[string]interface{}{
			"numsOfCmt": result,
		}),
	)
}

func (h *CmtHandler) StartWSConnection(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	defer ws.Close()
	log.Println("Connected")
	h.wsSvc.HandleConnection(ws)
	log.Println("Disconnected")
}
