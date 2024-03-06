package models

import (
	"database/sql"
	"strings"

	"github.com/doug-martin/goqu/v9"

	"github.com/Improwised/golang-api/constants"
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
	CreatedAt         string  `json:"created_at,omitempty" db:"created_at" goqu:"omitempty"`
	UpdatedAt         string  `json:"updated_at,omitempty" db:"updated_at" goqu:"omitempty"`
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

func (model *TitleModel) List(queries map[string]string, page int) ([]Title, error) {
	var titles []Title

	query := model.db.From(TitleTable)

	for key, value := range queries {
		if key == constants.TitleName {
			query = query.Where(goqu.Func("LOWER", goqu.I("title")).Like("%" + strings.ToLower(value) + "%"))
		}

		if key == constants.TitleType {
			query = query.Where(goqu.Ex{"type": value})
		}
	}

	//add offset
	offset := (page - 1) * constants.PageSize
	if offset >= 0 {
		query = query.Offset(uint(offset))
	}

	//add limit
	query = query.Limit(uint(constants.PageSize))

	sql, args, err := query.ToSQL()
	if err != nil {
		return titles, err
	}

	err = model.db.ScanStructs(&titles, sql, args...)
	if err != nil {
		return titles, err
	}

	return titles, err
}

func (model *TitleModel) Create(title Title) (Title, error) {

	_, err := model.db.Insert(TitleTable).Rows(title).Executor().Exec()
	if err != nil {
		return title, err
	}

	return model.GetById(title.ID)
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

func (model *TitleModel) Update(id string, title Title) (Title, error) {

	query := goqu.Update(TitleTable).Set(goqu.Ex{
		"id":                   title.ID,
		"title":                title.Title,
		"type":                 title.Type,
		"description":          title.Description,
		"release_year":         title.ReleaseYear,
		"age_certification":    title.AgeCertification,
		"runtime":              title.Runtime,
		"genres":               title.Genres,
		"production_countries": title.ProductionCountry,
		"seasons":              title.Seasons,
		"imdb_id":              title.IMDBID,
		"imdb_score":           title.IMDBScore,
		"imdb_votes":           title.IMDBVotes,
		"tmdb_popularity":      title.TMDBPopularity,
		"tmdb_score":           title.TMDBScore,
		"updated_at":           goqu.L("CURRENT_TIMESTAMP"),
	},
	).Where(goqu.Ex{"id": id})

	sql, args, err := query.ToSQL()
	if err != nil {
		return Title{}, err
	}

	_, err = model.db.Exec(sql, args...)
	if err != nil {
		return Title{}, err
	}

	return model.GetById(title.ID)
}
