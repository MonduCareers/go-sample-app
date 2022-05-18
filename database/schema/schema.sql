CREATE TABLE IF NOT EXISTS accounts
(
    id         UUID PRIMARY KEY   DEFAULT gen_random_uuid(),
    balance    BIGINT    NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
