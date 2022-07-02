package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Comment struct {
	bun.BaseModel `bun:"table:comments"`
	ID            int64     `bun:",pk"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	ChapterID     int64     `form:"chapter"`
	UserID        int64     `form:"author"`
	Content       string    `form:"comment"`
}
