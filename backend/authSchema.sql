CREATE TABLE IF NOT EXISTS users
(
    Baid         INTEGER PRIMARY KEY,
    username     TEXT not null,
    passwordHash TEXT not null,
    customTitleOn INTEGER NOT NULL DEFAULT 0
);