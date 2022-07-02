package models

import (
	"fmt"
	"time"

	"github.com/batt0s/mangajoy/database"
	"github.com/uptrace/bun"
)

type Manga struct {
	bun.BaseModel `bun:"table:mangas"`
	ID            int64      `bun:",pk"`
	CreatedAt     time.Time  `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time  `bun:",nullzero,notnull,default:current_timestamp"`
	ArtistID      int64      `form:"artist"`
	Title         string     `form:"title"`
	Description   string     `form:"description"`
	Chapters      []*Chapter `bun:"rel:has-many,join:manga_id"`
}

// This should be done before first use of the model
func (m *Manga) Init() {
	database.DB.RegisterModel((*Manga)(nil))
}

func (m *Manga) String() string {
	return fmt.Sprintf("Manga<%d, %d, %s, %v, %v>",
		m.ID, m.ArtistID, m.Title, m.CreatedAt, m.UpdatedAt)
}
