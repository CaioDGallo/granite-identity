-- Write your migrate up statements here

CREATE TYPE account_type AS ENUM ('Savings', 'Checking', 'Business', 'Credit', 'Investment');
CREATE TYPE account_status AS ENUM ('Active', 'Suspended', 'Closed', 'Pending');
CREATE TABLE IF NOT EXISTS accounts (
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL,
    balance DECIMAL(18, 2) NOT NULL DEFAULT 0.00,
    currency VARCHAR(3) NOT NULL DEFAULT 'BRL',
    status account_status NOT NULL DEFAULT 'Active',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    account_type account_type NOT NULL DEFAULT 'Checking',
    account_number VARCHAR(50) NOT NULL,
    last_activity TIMESTAMP NOT NULL DEFAULT NOW()
);

---- create above / drop below ----

DROP TABLE IF EXISTS accounts;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.

