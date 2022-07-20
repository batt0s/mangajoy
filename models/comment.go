package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/batt0s/mangajoy/database"
	"github.com/uptrace/bun"
)

type Comment struct {
	bun.BaseModel `bun:"table:comments"`
	ID            int64     `bun:"id,pk,autoincrement"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	ChapterID     int64     `form:"chapter"`
	UserID        int64     `form:"author"`
	Content       string    `form:"comment"`
}

func CreateCommentTable() error {
	ctx := context.Background()
	_, err := database.DB.NewCreateTable().Model((*Comment)(nil)).Exec(ctx)
	return err
}

// This should be called before first use of model
func InitCommentModel() {
	database.DB.RegisterModel((*Comment)(nil))
}

// Comment<ID, UserID, ChapterID, CreatedAt, UpdatedAt>
func (c *Comment) String() string {
	return fmt.Sprintf("Comment<%d, %d, %d, %v, %v>",
		c.ID, c.UserID, c.ChapterID, c.CreatedAt, c.UpdatedAt)
}

func (c *Comment) Create() error {
	ctx := context.Background()
	if !c.IsValid() {
		return errors.New("comment not valid")
	}
	_, err := database.DB.NewInsert().Model(c).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (c *Comment) Update() error {
	ctx := context.Background()
	if !c.IsValid() {
		return errors.New("comment not valid")
	}
	_, err := database.DB.NewUpdate().Model(c).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (c *Comment) IsValid() bool {
	if len(c.Content) > 512 || len(c.Content) < 4 {
		return false
	}
	return true
}

func (c *Comment) Delete() error {
	ctx := context.Background()
	_, err := database.DB.NewDelete().Model(c).WherePK().Exec(ctx)
	return err
}

func GetCommentByID(id int64) (*Comment, error) {
	comment := new(Comment)
	ctx := context.Background()
	err := database.DB.NewSelect().Model(comment).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
