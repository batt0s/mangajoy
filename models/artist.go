package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/batt0s/mangajoy/database"
	"github.com/uptrace/bun"
)

type Artist struct {
	bun.BaseModel `bun:"table:artists"`
	ID            int64     `bun:"id,pk,autoincrement"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	Name          string    `form:"name"`
	About         string    `form:"about"`
	Picture       string
	Mangas        []*Manga `bun:"rel:has-many,join:id=artist_id"`
}

func CreateArtistTable() error {
	ctx := context.Background()
	_, err := database.DB.NewCreateTable().Model((*Artist)(nil)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

// This should be called before first use of the model
func InitArtistModel() {
	database.DB.RegisterModel((*Artist)(nil))
}

// Artist<ID, Name, Mangas, CreatedAt, UpdatedAt>
func (a *Artist) String() string {
	return fmt.Sprintf("Artist<%d, %s, %d, %v, %v>",
		a.ID, a.Name, len(a.Mangas), a.CreatedAt, a.UpdatedAt)
}

func (a *Artist) Create() error {
	ctx := context.Background()
	if !a.IsValid() {
		return errors.New("artist not valid")
	}
	_, err := database.DB.NewInsert().Model(a).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (a *Artist) Update() error {
	ctx := context.Background()
	if !a.IsValid() {
		return errors.New("artist not valid")
	}
	_, err := database.DB.NewUpdate().Model(a).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (a *Artist) IsValid() bool {
	if len(a.Name) < 6 || len(a.Name) > 32 {
		return false
	}
	if len(a.About) < 32 || len(a.About) > 512 {
		return false
	}
	return true
}

func (a *Artist) Delete() error {
	ctx := context.Background()
	_, err := database.DB.NewDelete().Model(a).WherePK().Exec(ctx)
	return err
}

func GetArtistByID(id int64) (*Artist, error) {
	artist := new(Artist)
	ctx := context.Background()
	err := database.DB.NewSelect().Model(artist).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return artist, nil
}
