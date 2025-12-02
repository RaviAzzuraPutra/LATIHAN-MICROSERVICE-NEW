CREATE EXTENSION IF NOT EXISTS "pgcrypto";

DROP TABLE IF EXISTS orders;

CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    productid UUID,
    quantity INTEGER
);