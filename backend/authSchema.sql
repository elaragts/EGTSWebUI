CREATE TABLE IF NOT EXISTS "users"
(
    Baid         INTEGER
        primary key,
    username     TEXT not null,
    passwordHash TEXT not null
);
