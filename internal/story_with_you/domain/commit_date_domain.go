package domain

import (
	"context"
	"time"
)

type CommitDate struct {
	ID      int        `json:"id"`
	Partner string     `json:"partner"`
	Date    *time.Time `json:"commitDate"`
}

type CommitDateRepo interface {
	Get(ctx context.Context, id string) (*CommitDate, error)
	UpdateCommitDate(ctx context.Context, id string, commitDate time.Time) error
}

type CommitDateSvc interface {
	GetCommitDate(ctx context.Context, id string) (*CommitDate, error)
	UpdateCommitDate(ctx context.Context, id string, commitDate time.Time) error
}

type UpdateCommitDateReq struct {
	Date time.Time `json:"date" validate:"required"`
}
