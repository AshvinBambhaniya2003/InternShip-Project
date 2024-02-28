package models

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

// TitleTable represent table name
const TitleTable = "titles"

// Title model

type Title struct {
	ID                string  `json:"id" db:"id"`
	Title             string  `json:"title" db:"title" validate:"required"`
	Type              string  `json:"type" db:"type" validate:"required"`
	Description       string  `json:"description" db:"description"`
	ReleaseYear       int     `json:"release_year" db:"release_year" validate:"required"`
	AgeCertification  string  `json:"age_certification" db:"age_certification" validate:"required"`
	Runtime           int     `json:"runtime" db:"runtime" validate:"required"`
	Genres            string  `json:"genres" db:"genres" validate:"required"`
	ProductionCountry string  `json:"production_countries" db:"production_countries" validate:"required"`
	Seasons           int     `json:"seasons" db:"seasons"`
	IMDBID            string  `json:"imdb_id" db:"imdb_id"`
	IMDBScore         float64 `json:"imdb_score" db:"imdb_score"`
	IMDBVotes         float64 `json:"imdb_votes" db:"imdb_votes"`
	TMDBPopularity    float64 `json:"tmdb_popularity" db:"tmdb_popularity"`
	TMDBScore         float64 `json:"tmdb_score" db:"tmdb_score"`
	CreatedAt         string  `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt         string  `json:"updated_at,omitempty" db:"updated_at,omitempty"`
}

// TitleModel implements title related database operations
type TitleModel struct {
	db *goqu.Database
}

// InitTitleModel Init model
func InitTitleModel(goqu *goqu.Database) (*TitleModel, error) {
	return &TitleModel{
		db: goqu,
	}, nil
}

func (model *TitleModel) List() ([]Title, error) {
	var titles []Title
	if err := model.db.From(TitleTable).ScanStructs(&titles); err != nil {
		return nil, err
	}
	return titles, nil
}

// GetById get title by id
func (model *TitleModel) GetById(id string) (Title, error) {
	title := Title{}
	found, err := model.db.From(TitleTable).Where(goqu.Ex{"id": id}).ScanStruct(&title)

	if err != nil {
		return title, err
	}

	if !found {
		return title, sql.ErrNoRows
	}

	return title, err
}

// DeleteTitle deletes a title by its ID
func (model *TitleModel) Delete(id string) error {
	_, err := model.db.Delete(TitleTable).Where(goqu.Ex{"id": id}).Executor().Exec()
	return err
}
