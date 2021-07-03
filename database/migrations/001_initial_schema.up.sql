BEGIN;

CREATE TABLE IF NOT EXISTS user_types (
    id      INT PRIMARY KEY,
    description    TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id              SERIAL     PRIMARY KEY,
    user_type_id    INT     NOT NULL,
    first_name      TEXT    NOT NULL,
    last_name       TEXT    NOT NULL,
    email           TEXT    NOT NULL UNIQUE,
    created_date    TIMESTAMP NOT NULL DEFAULT Now(),
    username        TEXT    NULL UNIQUE,
    phone_no        TEXT    NULL UNIQUE,
    gender          TEXT    NULL,
    profile_picture VARCHAR(1000) NULL,
    bio             VARCHAR(1000) NULL,
    background_picture TEXT         NULL,
    FOREIGN KEY (user_type_id) REFERENCES user_types(id)
);

CREATE TABLE IF NOT EXISTS clients (
    fanfit_user_id      INT     NOT NULL,
    temp_field          TEXT       NULL,
    PRIMARY KEY (fanfit_user_id),
    FOREIGN KEY (fanfit_user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS creators (
    fanfit_user_id      INT        NOT NULL,
    payment_info        TEXT       NULL,
    logo_picture        TEXT       NULL,
    PRIMARY KEY (fanfit_user_id),
    FOREIGN KEY (fanfit_user_id) REFERENCES users(id)
);

COMMIT;