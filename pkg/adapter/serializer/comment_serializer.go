package serializer

import "time"

type Comment struct {
	ID        int        `json:"id"`
	PostID    string     `json:"postID,omitempty"`
	Author    string     `json:"author"`
	Content   string     `json:"content"`
	Replies   []Comment  `json:"replies,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type GetCommentsOfPostRes struct {
	Comments []Comment `json:"comments"`
}
