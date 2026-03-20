-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id                      SERIAL NOT NULL PRIMARY KEY,
    email                   VARCHAR(255) NOT NULL UNIQUE, 
    nickname                VARCHAR(100) NOT NULL,
    password_hash           CHAR(60) NOT NULL,
    refresh_token_version   TIMESTAMP(6) WITHOUT TIME ZONE
);

-- +goose Down
DROP TABLE IF EXISTS users; 
