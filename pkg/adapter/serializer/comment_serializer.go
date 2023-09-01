package serializer

import (
	"time"
)

type Comment struct {
	ID        int        `json:"id"`
	PostSlug  string     `json:"postSlug,omitempty"`
	Author    string     `json:"author"`
	Content   string     `json:"content"`
	ParentID  *int       `json:"parentId"`
	Replies   []Comment  `json:"replies,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type GetCommentsOfPostRes struct {
	Comments []Comment `json:"comments"`
}

type CreateCommentOfPostReq struct {
	PostSlug string `json:"postSlug" validate:"required"`
	ParentId *int   `json:"parentId,omitempty" validate:"omitempty,number"`
	Content  string `json:"content" validate:"required"`
}

type CreateCommentOfPostRes struct {
	ID int `json:"id"`
}

type GetNumberCommentsOfPostsReq struct {
	PostSlugs []string `json:"postSlugs" validate:"required"`
}
