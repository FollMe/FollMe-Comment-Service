package service

import (
	"context"
	"follme/comment-service/internal/story_with_you/domain"
	"time"
)

type CommitDateSvc struct {
	repo domain.CommitDateRepo
}

func NewCommitDateSvc(repo domain.CommitDateRepo) *CommitDateSvc {
	return &CommitDateSvc{
		repo: repo,
	}
}

var _ domain.CommitDateSvc = &CommitDateSvc{}

func (c CommitDateSvc) GetCommitDate(ctx context.Context, id string) (*domain.CommitDate, error) {
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
