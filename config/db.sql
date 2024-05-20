-- Active: 1696507852877@@127.0.0.1@5432@auth_go@public

CREATE Table users (username VARCHAR(100));

DROP TABLE users;

CREATE TABLE users (
    id VARCHAR DEFAULT REPLACE(
        uuid_generate_v4 ()::text,
        '-',
        ''
    ) NOT NULL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE users_hash (
    id VARCHAR DEFAULT REPLACE(
        uuid_generate_v4 ()::text,
        '-',
        ''
    ) NOT NULL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    hashed_password BYTEA NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    deleted_at TIMESTAMP NULL
);

DROP TABLE users_hash;

SELECT * from users_hash;