package structs

// All request sturcts
// Request struct have Req prefix

type ReqRegisterUser struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Roles     string `json:"roles" validate:"required"`
}

type ReqLoginUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ReqRegisterTitle struct {
	Title             string  `json:"title" validate:"required"`
	Type              string  `json:"type" validate:"required"`
	Description       string  `json:"description" db:"description"`
	ReleaseYear       int     `json:"release_year" validate:"required"`
	AgeCertification  string  `json:"age_certification" validate:"required"`
	Runtime           int     `json:"runtime" validate:"required"`
	Genres            string  `json:"genres" db:"genres" validate:"required"`
	ProductionCountry string  `json:"production_countries" db:"production_countries" validate:"required"`
	Seasons           int     `json:"seasons" validate:"required"`
	IMDBID            string  `json:"imdb_id" validate:"required"`
	IMDBScore         float64 `json:"imdb_score" validate:"required"`
	IMDBVotes         float64 `json:"imdb_votes" validate:"required"`
	TMDBPopularity    float64 `json:"tmdb_popularity" validate:"required"`
	TMDBScore         float64 `json:"tmdb_score" validate:"required"`
}

type ReqRegisterCredit struct {
	PersonID  int    `json:"person_id" db:"person_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Character string `json:"character" db:"character"`
	Role      string `json:"role" db:"role" validate:"required"`
}
