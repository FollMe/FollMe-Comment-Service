package domain

import (
	"time"
)

type Comment struct {
	Id        int        `json:"id"`
	PostId    string     `json:"postId"`
	ParentId  int        `json:"parentId"`
	Author    string     `json:"author"`
	Content   string     `json:"content"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
