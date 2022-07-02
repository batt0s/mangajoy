package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Chapter struct {
	bun.BaseModel `bun:"table:chapters"`
	ID            int64      `bun:",pk"`
	CreatedAt     time.Time  `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time  `bun:",nullzero,notnull,default:current_timestamp"`
	MangaID       int64      `form:"manga"`
	UploaderID    int64      `form:"uploader"`
	Title         string     `form:"title"`
	Description   string     `form:"description"`
	Pages         []string   `bun:",array"`
	Comments      []*Comment `bun:"rel:has-many,join:id=chapter_id"`
}
