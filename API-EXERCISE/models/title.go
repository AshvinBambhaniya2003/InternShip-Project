package models

import (
	"database/sql"
	"strings"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/rs/xid"

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

func (model *TitleModel) Insert(title Title) (Title, error) {
	title.ID = xid.New().String()

	time := time.Now().Format("2006-01-02T15:04:05.999999Z")

	_, err := model.db.Insert(TitleTable).Rows(
		Title{
			ID:                title.ID,
			Title:             title.Title,
			Type:              title.Type,
			Description:       title.Description,
			ReleaseYear:       title.ReleaseYear,
			AgeCertification:  title.AgeCertification,
			Runtime:           title.Runtime,
			Genres:            title.Genres,
			ProductionCountry: title.ProductionCountry,
			Seasons:           title.Seasons,
			IMDBID:            title.IMDBID,
			IMDBScore:         title.IMDBScore,
			IMDBVotes:         title.IMDBVotes,
			TMDBPopularity:    title.TMDBPopularity,
			TMDBScore:         title.TMDBScore,
			CreatedAt:         time,
			UpdatedAt:         time,
		},
	).Executor().Exec()

	if err != nil {
		return title, err
	}

	title, err = model.GetById(title.ID)

	return title, err
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

func (model *TitleModel) Update(id string, title Title) error {
	_, err := model.db.Update(TitleTable).Set(title).Where(goqu.Ex{"id": id}).Executor().Exec()
	return err
}
