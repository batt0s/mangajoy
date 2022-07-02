package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Artist struct {
	bun.BaseModel `bun:"table:artists"`
	ID            int64     `bun:"id,pk,autoincrement"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	Name          string    `form:"name"`
	About         string    `form:"about"`
	Mangas        []*Manga  `bun:"rel:has-many,join:artist_id"`
}
