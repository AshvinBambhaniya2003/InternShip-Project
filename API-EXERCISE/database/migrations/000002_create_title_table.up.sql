-- +migrate Up
CREATE TABLE IF NOT EXISTS titles (
    id CHAR (20) PRIMARY KEY,
    title VARCHAR(255),
    type VARCHAR(20),
    description TEXT,
    release_year INTEGER,
    age_certification VARCHAR(10),
    runtime INTEGER,
    genres TEXT,
    production_countries TEXT,
    seasons INTEGER,
    imdb_id VARCHAR(20),
    imdb_score FLOAT,
    imdb_votes FLOAT,
    tmdb_popularity FLOAT,
    tmdb_score FLOAT,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
