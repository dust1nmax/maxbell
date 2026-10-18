package models

import "time"

type Post struct {
	Title       string    `json:"title" db:"title" binging:"required"`
	Context     string    `json:"context" db:"context" binging:"required"`
	Id          int64     `json:"id" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id" binging:"required"`
	CommunityID int64     `json:"community_id" db:"community_id" binging:"required"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}
