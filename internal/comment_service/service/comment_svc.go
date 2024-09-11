package service

import (
	"context"
	"follme/comment-service/internal/comment_service/domain"
)

type CommentSvc struct {
	repo  domain.CommentRepo
	wsSvc domain.WebSocketSvc
}

func NewCommentSvc(repo domain.CommentRepo, wsSvc domain.WebSocketSvc) *CommentSvc {
	return &CommentSvc{
		repo:  repo,
		wsSvc: wsSvc,
	}
}

var _ domain.CommentSvc = &CommentSvc{}

func (c CommentSvc) GetCommentsOfPost(ctx context.Context, postSlug string) ([]domain.Comment, error) {
	return c.repo.List(ctx, domain.ListOpts{
		PostSlug: postSlug,
	})
}

func (c CommentSvc) GetNumberCommentsOfPosts(ctx context.Context, postSlugs []string) (map[string]int, error) {
	return c.repo.GetNumberRecord(ctx, postSlugs)
}

func (c CommentSvc) InsertCommentOfPost(ctx context.Context, opts domain.CreateCommentOpts) (*domain.Comment, error) {
	cmt := domain.NewComment(opts)
	cmt, err := c.repo.CreateOne(ctx, *cmt)
	if err != nil {
		return nil, err
	}
	c.wsSvc.BroadCastToPostRoom(cmt)
	return cmt, nil
}
