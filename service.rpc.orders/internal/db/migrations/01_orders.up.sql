CREATE TABLE orders
(
    "id"          TEXT PRIMARY KEY,
    "created_at"  TIMESTAMPTZ NOT NULL,
    "updated_at"  TIMESTAMPTZ,

    "asset_type"  INT        NOT NULL,
    "instruction" INT         NOT NULL,
    "amount"      INT         NOT NULL,
    "currency"    INT         NOT NULL
);
