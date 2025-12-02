CREATE EXTENSION IF NOT EXISTS "pgcrypto";

DROP TABLE IF EXISTS stocks;

CREATE TABLE stocks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT,
    quantity INTEGER
);