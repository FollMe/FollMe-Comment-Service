package repository

import (
	"context"
	"database/sql"
	"time"

	"follme/comment-service/pkg/model"
)

type comment struct {
	ID        int
	PostSlug  string
	ParentID  *int
	Author    string
	Content   string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type replyComment struct {
	ID        *int
	ParentID  *int
	Author    *string
	Content   *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type CommentRepo struct {
	DB *sql.DB
}

func NewCommentRepo(db *sql.DB) *CommentRepo {
	return &CommentRepo{
		DB: db,
	}
}

var _ model.CommentRepo = &CommentRepo{}

func (c comment) toModel() *model.Comment {
	return model.CommentFactory(model.CommentFactoryOpts{
		ID:        c.ID,
		PostSlug:  c.PostSlug,
		Author:    c.Author,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	})
}

func (c replyComment) toModel() *model.Comment {
	return model.CommentFactory(model.CommentFactoryOpts{
		ID:        *c.ID,
		Author:    *c.Author,
		Content:   *c.Content,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	})
}

func (c CommentRepo) List(ctx context.Context, opts model.ListOpts) ([]model.Comment, error) {
	query := `
		select c.id, c.post_slug, c.author, c.content, c.created_at, c.updated_at,
		rc.id, rc.author, rc.content, rc.created_at, rc.updated_at
		from comment c left join comment rc on c.id = rc.parent_id
		where c.post_slug = $1 and c.parent_id is null
		order by c.id ASC, rc.id ASC 
	`
	cmts := []model.Comment{}
	rows, err := c.DB.Query(query, opts.PostSlug)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			cmt      comment
			replyCmt replyComment
		)
		err := rows.Scan(
			&cmt.ID, &cmt.PostSlug, &cmt.Author, &cmt.Content, &cmt.CreatedAt, &cmt.UpdatedAt,
			&replyCmt.ID, &replyCmt.Author, &replyCmt.Content, &replyCmt.CreatedAt, &replyCmt.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		cmtModel := cmt.toModel()
		if len(cmts) == 0 || cmt.ID != cmts[len(cmts)-1].ID() {
			if replyCmt.ID != nil {
				cmtModel.AppendReply(*replyCmt.toModel())
			}
			cmts = append(cmts, *cmtModel)
			continue
		}

		cmts[len(cmts)-1].AppendReply(*replyCmt.toModel())
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cmts, nil
}

func (c CommentRepo) CreateOne(ctx context.Context, opts model.Comment) (*model.Comment, error) {
	var (
		commentId int
		createdAt *time.Time
	)
	query := "insert into comment (post_slug, parent_id, author, content) values ($1, $2, $3, $4) returning id, created_at"
	err := c.DB.QueryRowContext(ctx, query, opts.PostSlug(), opts.ParentID(), opts.Author(), opts.Content()).Scan(&commentId, &createdAt)
	if err != nil {
		return nil, err
	}

	return model.CommentFactory(model.CommentFactoryOpts{
		ID:        commentId,
		PostSlug:  opts.PostSlug(),
		ParentId:  opts.ParentID(),
		Author:    opts.Author(),
		Content:   opts.Content(),
		CreatedAt: createdAt,
	}), nil
}
