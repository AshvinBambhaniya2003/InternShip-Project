package models

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
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
	CreatedAt string `json:"created_at,omitempty" db:"created_at" goqu:"omitempty"`
	UpdatedAt string `json:"updated_at,omitempty" db:"updated_at" goqu:"omitempty"`
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

func (model *CreditModel) Insert(credit Credit) (Credit, error) {

	_, err := model.db.Insert(CreditTable).Rows(credit).Executor().Exec()
	if err != nil {
		return credit, err
	}

	return model.GetById(credit.Id)
}

func (model *CreditModel) ListCredits(id string) ([]Credit, error) {
	var credits []Credit
	err := model.db.From(CreditTable).Where(goqu.Ex{"title_id": id}).ScanStructs(&credits)

	return credits, err
}

func (model *CreditModel) GetById(id string) (Credit, error) {
	credit := Credit{}
	found, err := model.db.From(CreditTable).Where(goqu.Ex{"id": id}).ScanStruct(&credit)

	if err != nil {
		return credit, err
	}

	if !found {
		return credit, sql.ErrNoRows
	}

	return credit, err
}

func (model *CreditModel) Delete(id string) error {
	_, err := model.db.Delete(CreditTable).Where(goqu.Ex{"id": id}).Executor().Exec()
	return err
}
