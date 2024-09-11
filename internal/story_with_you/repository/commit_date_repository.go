package repository

import (
	"context"
	"database/sql"
	"time"

	"follme/comment-service/internal/story_with_you/domain"
)

type CommitDateRepo struct {
	DB *sql.DB
}

func NewCommitDateRepo(db *sql.DB) *CommitDateRepo {
	return &CommitDateRepo{
		DB: db,
	}
}

var _ domain.CommitDateRepo = &CommitDateRepo{}

func (c CommitDateRepo) Get(ctx context.Context, id string) (*domain.CommitDate, error) {
	query := `
		select id, partner, date
		from commit_date
		where id = $1
	`

	var commitDate domain.CommitDate

	row := c.DB.QueryRow(query, id)

	err := row.Scan(&commitDate.ID, &commitDate.Partner, &commitDate.Date)

	if err != nil {
		if err == sql.ErrNoRows {
			return &commitDate, nil
		}
		return nil, err
	}

	return &commitDate, nil
}

func (c CommitDateRepo) UpdateCommitDate(ctx context.Context, id string, commitDate time.Time) error {
	query := "update commit_date set date = $1 where id = $2"
	_, err := c.DB.Exec(query, commitDate, id)
	if err != nil {
		return err
	}

	return nil
}
