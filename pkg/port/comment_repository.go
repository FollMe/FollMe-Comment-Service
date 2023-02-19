package port

import "follme/comment-service/pkg/domain"

type CommentRepository interface {
	GetAllCmtOfPost(postId string) ([]domain.Comment, error)
}
