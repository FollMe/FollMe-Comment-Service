package model

import (
	"context"
	"time"
)

type Comment struct {
	id        int
	postSlug  string
	parentId  *int
	replies   []Comment
	author    string
	content   string
	createdAt time.Time
	updatedAt *time.Time
}

func (c Comment) ID() int {
	return c.id
}
func (c Comment) PostSlug() string {
	return c.postSlug
}
func (c Comment) ParentID() *int {
	return c.parentId
}
func (c Comment) Replies() []Comment {
	return c.replies
}
func (c Comment) Author() string {
	return c.author
}
func (c Comment) Content() string {
	return c.content
}
func (c Comment) CreatedAt() time.Time {
	return c.createdAt
}
func (c Comment) UpdatedAt() *time.Time {
	return c.updatedAt
}

func (c *Comment) AppendReply(replyCmt Comment) {
	c.replies = append(c.replies, replyCmt)
}

type CommentRepo interface {
	List(ctx context.Context, opts ListOpts) ([]Comment, error)
	CreateOne(ctx context.Context, opts Comment) (*Comment, error)
}

type CommentSvc interface {
	GetCommentsOfPost(ctx context.Context, postId string) ([]Comment, error)
	InsertCommentOfPost(ctx context.Context, opts CreateCommentOpts) (*Comment, error)
}

type ListOpts struct {
	ParentID int
	PostSlug string
}

type CommentFactoryOpts struct {
	ID        int
	PostSlug  string
	ParentId  *int
	Replies   []Comment
	Author    string
	Content   string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func CommentFactory(co CommentFactoryOpts) *Comment {
	rp := Comment{
		id:        co.ID,
		postSlug:  co.PostSlug,
		parentId:  co.ParentId,
		replies:   co.Replies,
		author:    co.Author,
		content:   co.Content,
		createdAt: *co.CreatedAt,
		updatedAt: co.UpdatedAt,
	}

	return &rp
}

type CreateCommentOpts struct {
	PostSlug string `json:"postSlug,omitempty"`
	ParentId *int   `json:"parentId"`
	Content  string `json:"content"`
	Author   string `json:"author"`
}

func NewComment(opts CreateCommentOpts) *Comment {
	now := time.Now()
	return &Comment{
		postSlug:  opts.PostSlug,
		parentId:  opts.ParentId,
		content:   opts.Content,
		author:    opts.Author,
		createdAt: now,
	}
}
