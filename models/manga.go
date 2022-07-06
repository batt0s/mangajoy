package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/batt0s/mangajoy/database"
	"github.com/uptrace/bun"
)

type Manga struct {
	bun.BaseModel    `bun:"table:mangas"`
	ID               int64     `bun:",pk"`
	CreatedAt        time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt        time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	ArtistID         int64     `form:"artist"`
	Title            string    `form:"title"`
	AlternativeTitle string    `form:"alternative_title"`
	Description      string    `form:"description"`
	Tags             []string  `bun:",array" form:"tags"`
	Cover            string
	Chapters         []*Chapter `bun:"rel:has-many,join:id=manga_id"`
}

func CreateMangaTable() error {
	ctx := context.Background()
	_, err := database.DB.NewCreateTable().Model((*Manga)(nil)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

// This should be done before first use of the model
func InitMangaModel() {
	database.DB.RegisterModel((*Manga)(nil))
}

// Manga<ID, ArtistID, Title, CreatedAt, UpdatedAt>
func (m *Manga) String() string {
	return fmt.Sprintf("Manga<%d, %d, %s, %v, %v>",
		m.ID, m.ArtistID, m.Title, m.CreatedAt, m.UpdatedAt)
}

func (m *Manga) Create() error {
	ctx := context.Background()
	if !m.IsValid() {
		return errors.New("manga not valid")
	}
	_, err := database.DB.NewInsert().Model(m).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m *Manga) Update() error {
	ctx := context.Background()
	if !m.IsValid() {
		return errors.New("manga not valid")
	}
	_, err := database.DB.NewUpdate().Model(m).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m *Manga) IsValid() bool {
	if len(m.Title) > 128 || len(m.Title) < 4 {
		return false
	}
	if len(m.AlternativeTitle) > 128 || len(m.AlternativeTitle) < 4 {
		return false
	}
	if len(m.Description) > 512 || len(m.Description) < 32 {
		return false
	}
	return true
}

func (m *Manga) Delete() error {
	ctx := context.Background()
	_, err := database.DB.NewDelete().Model(m).WherePK().Exec(ctx)
	return err
}

func GetMangaByID(id int64) (*Manga, error) {
	manga := &Manga{}
	ctx := context.Background()
	err := database.DB.NewSelect().Model(manga).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return manga, nil
}

func GetLastNManga(n int) ([]Manga, error) {
	mangas := make([]Manga, n)
	ctx := context.Background()
	err := database.DB.NewSelect().Model((*Manga)(nil)).Limit(n).Scan(ctx, &mangas)
	if err != nil {
		return []Manga{}, err
	}
	return mangas, nil
}
