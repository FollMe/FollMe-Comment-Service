package service

import (
	"context"
	"follme/comment-service/pkg/model"
)

type CommentSvc struct {
	repo model.CommentRepo
}

func NewCommentSvc(repo model.CommentRepo) *CommentSvc {
	return &CommentSvc{
		repo: repo,
	}
}

var _ model.CommentSvc = &CommentSvc{}

func (c CommentSvc) GetCommentsOfPost(ctx context.Context, postSlug string) ([]model.Comment, error) {
	return c.repo.List(ctx, model.ListOpts{
		PostSlug: postSlug,
	})
}

func (c CommentSvc) InsertCommentOfPost(ctx context.Context, opts model.CreateCommentOpts) (*model.Comment, error) {
	cmt := model.NewComment(opts)
	return c.repo.CreateOne(ctx, *cmt)
}
