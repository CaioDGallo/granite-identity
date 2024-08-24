-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS transactions (
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    account_id uuid NOT NULL REFERENCES accounts(id),
    type VARCHAR(50) NOT NULL,
    amount DECIMAL(18, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'BRL',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    reference_id VARCHAR(50) NOT NULL,
    description TEXT
);

---- create above / drop below ----

DROP TABLE IF EXISTS transactions;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.

