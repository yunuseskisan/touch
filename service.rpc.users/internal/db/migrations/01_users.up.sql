CREATE TABLE users
(
    "id"                        TEXT PRIMARY KEY,
    "created_at"                TIMESTAMPTZ NOT NULL,
    "updated_at"                TIMESTAMPTZ,

    "first_name"                TEXT        NOT NULL,
    "last_name"                 TEXT        NOT NULL
);
