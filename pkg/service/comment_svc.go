package service

import (
	"context"
	"follme/comment-service/pkg/model"
)

type CommentSvc struct {
	repo  model.CommentRepo
	wsSvc model.WebSocketSvc
}

func NewCommentSvc(repo model.CommentRepo, wsSvc model.WebSocketSvc) *CommentSvc {
	return &CommentSvc{
		repo:  repo,
		wsSvc: wsSvc,
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
	cmt, err := c.repo.CreateOne(ctx, *cmt)
	if err != nil {
		return nil, err
	}
	c.wsSvc.BroadCastToPostRoom(cmt)
	return cmt, nil
}
