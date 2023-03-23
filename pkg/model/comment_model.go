package model

import (
	"context"
	"time"
)

type Comment struct {
	id        int
	postID    string
	replies   []Comment
	author    string
	content   string
	createdAt *time.Time
	updatedAt *time.Time
}

func (c Comment) ID() int {
	return c.id
}
func (c Comment) PostID() string {
	return c.postID
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
func (c Comment) CreatedAt() *time.Time {
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
}

type CommentSvc interface {
	GetCommentsOfPost(ctx context.Context, postId string) ([]Comment, error)
}

type ListOpts struct {
	ParentID int
	PostID   string
}

type CommentFactoryOpts struct {
	ID        int
	PostID    string
	Replies   []Comment `json:"author"`
	Author    string
	Content   string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func CommentFactory(co CommentFactoryOpts) *Comment {
	rp := Comment{
		id:        co.ID,
		postID:    co.PostID,
		replies:   co.Replies,
		author:    co.Author,
		content:   co.Content,
		createdAt: co.CreatedAt,
		updatedAt: co.UpdatedAt,
	}

	return &rp
}
