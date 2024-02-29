package models

import (
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/rs/xid"
)

// CreditTable represent table name
const CreditTable = "credits"

// Credit model

type Credit struct {
	Id        string `json:"id" db:"id"`
	PersonID  int    `json:"person_id" db:"person_id"`
	TitleID   string `json:"title_id" db:"title_id" validate:"required"`
	Name      string `json:"name" db:"name" validate:"required" `
	Character string `json:"character" db:"character"`
	Role      string `json:"role" db:"role" validate:"required"`
	CreatedAt string `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty" db:"updated_at,omitempty"`
}

// CreditModel implements credit related database operations
type CreditModel struct {
	db *goqu.Database
}

// InitCreditModel Init model
func InitCreditModel(goqu *goqu.Database) (*CreditModel, error) {
	return &CreditModel{
		db: goqu,
	}, nil
}

func (model *CreditModel) Insert(credit Credit) error {
	id := xid.New().String()

	currentTime := time.Now().Format("2006-01-02T15:04:05.999999Z")

	_, err := model.db.Insert(CreditTable).Rows(
		Credit{
			Id:        id,
			PersonID:  credit.PersonID,
			TitleID:   credit.TitleID,
			Name:      credit.Name,
			Character: credit.Character,
			Role:      credit.Role,
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		},
	).Executor().Exec()

	return err
}
