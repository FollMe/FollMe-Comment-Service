package service

import (
	"context"
	"follme/comment-service/pkg/model"
	"time"
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

func (c CommitDateSvc) UpdateCommitDate(ctx context.Context, id string, commitDate time.Time) error {
	record, err := c.GetCommitDate(ctx, id)
	if err != nil {
		return err
	}
	if record.Date != nil {
		return nil
	}

	return c.repo.UpdateCommitDate(ctx, id, commitDate)
}
