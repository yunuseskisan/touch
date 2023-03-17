CREATE TABLE assets
(
    "id"         TEXT PRIMARY KEY,
    "created_at" TIMESTAMPTZ           NOT NULL,
    "updated_at" TIMESTAMPTZ,

    "user_id"    TEXT REFERENCES users NOT NULL,
    "type"       INT                   NOT NULL,
    "units"      INT                   DEFAULT 0,

    UNIQUE ("user_id", "type")
);
