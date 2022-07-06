package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/batt0s/mangajoy/database"
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

func CreateChapterTable() error {
	ctx := context.Background()
	_, err := database.DB.NewCreateTable().Model((*Chapter)(nil)).Exec(ctx)
	return err
}

// This should be called before first use of model
func InitChapterModel() {
	database.DB.RegisterModel((*Chapter)(nil))
}

// Chapter<ID, Title, MangaID, UploaderID, CreatedAt, UpdatedAt>
func (c *Chapter) String() string {
	return fmt.Sprintf("Chapter<%d, %s, %d, %d, %v, %v>",
		c.ID, c.Title, c.MangaID, c.UploaderID, c.CreatedAt, c.UpdatedAt)
}

func (c *Chapter) Create() error {
	ctx := context.Background()
	if !c.IsValid() {
		return errors.New("chapter not valid")
	}
	_, err := database.DB.NewInsert().Model(c).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (c *Chapter) Update() error {
	ctx := context.Background()
	if !c.IsValid() {
		return errors.New("chapter not valid")
	}
	_, err := database.DB.NewUpdate().Model(c).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (c *Chapter) IsValid() bool {
	if len(c.Title) < 4 || len(c.Title) > 128 {
		return false
	}
	if len(c.Description) < 8 || len(c.Description) > 512 {
		return false
	}
	return true
}

func (c *Chapter) Delete() error {
	ctx := context.Background()
	_, err := database.DB.NewDelete().Model(c).WherePK().Exec(ctx)
	return err
}

func GetChapterByID(id int64) (*Chapter, error) {
	chapter := new(Chapter)
	ctx := context.Background()
	err := database.DB.NewSelect().Model(chapter).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return chapter, nil
}
