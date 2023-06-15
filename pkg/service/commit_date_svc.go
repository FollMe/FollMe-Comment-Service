package service

import (
	"context"
	"follme/comment-service/pkg/model"
)

type CommitDateSvc struct {
	repo model.CommitDateRepo
}

func NewCommitDateSvc(repo model.CommitDateRepo) *CommitDateSvc {
	return &CommitDateSvc{
		repo: repo,
	}
}

var _ model.CommitDateSvc = &CommitDateSvc{}

func (c CommitDateSvc) GetCommitDate(ctx context.Context, id string) (*model.CommitDate, error) {
	return c.repo.Get(ctx, id)
}
