-- Migration Up: create-table-customers
BEGIN;
CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    identity_number VARCHAR(16) UNIQUE NOT NULL,
    phone_number VARCHAR(15) UNIQUE NOT NULL,
    account_number VARCHAR(10) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
COMMIT;
