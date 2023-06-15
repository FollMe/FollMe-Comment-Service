package repository

import (
	"context"
	"database/sql"

	"follme/comment-service/pkg/model"
)

type CommitDateRepo struct {
	DB *sql.DB
}

func NewCommitDateRepo(db *sql.DB) *CommitDateRepo {
	return &CommitDateRepo{
		DB: db,
	}
}

var _ model.CommitDateRepo = &CommitDateRepo{}

func (c CommitDateRepo) Get(ctx context.Context, id string) (*model.CommitDate, error) {
	query := `
		select id, partner, date
		from commit_date
		where id = $1
	`

	var commitDate model.CommitDate

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
