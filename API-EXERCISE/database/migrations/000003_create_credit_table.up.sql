-- +migrate Up
CREATE TABLE IF NOT EXISTS credits (
    id CHAR (20) PRIMARY KEY,
    person_id INTEGER NOT NULL,
    title_id CHAR (20) REFERENCES titles(id),
    name VARCHAR(255),
    character VARCHAR(255),
    role VARCHAR(255),
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);