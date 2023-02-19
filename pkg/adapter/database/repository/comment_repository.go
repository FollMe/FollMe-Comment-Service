package repository

import (
	"database/sql"
	"follme/comment-service/pkg/domain"
	"follme/comment-service/pkg/port"
)

type commentRepository struct {
	DB *sql.DB
}

func NewCommentRepo(db *sql.DB) port.CommentRepository {
	return &commentRepository{
		DB: db,
	}
}

func (cmtRepo *commentRepository) GetAllCmtOfPost(postId string) ([]domain.Comment, error) {
	cmtList := make([]domain.Comment, 0)
	return cmtList, nil
}
