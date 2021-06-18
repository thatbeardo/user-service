CREATE TABLE user_types (
    id      INT PRIMARY KEY,
    disc    TEXT NOT NULL
);
CREATE TABLE users (
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
    FOREIGN KEY (user_type_id) REFERENCES user_types(id)
);
CREATE TABLE creators (
    fanfit_user_id      SERIAL     NOT NULL,
    payment_info        TEXT    NOT NULL,
    logo_picture        TEXT    NOT NULL,
    background_picture  TEXT    NOT NULL,
    PRIMARY KEY (fanfit_user_id),
    FOREIGN KEY (fanfit_user_id) REFERENCES users(id)
);
CREATE TABLE clients (
    fanfit_user_id      SERIAL     NOT NULL,
    PRIMARY KEY (fanfit_user_id),
    FOREIGN KEY (fanfit_user_id) REFERENCES users(id)
);


INSERT INTO user_types(id, disc) VALUES(0, 'Creator');
INSERT INTO user_types(id, disc) VALUES(1, 'Clients');