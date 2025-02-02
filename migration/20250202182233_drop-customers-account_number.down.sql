-- Migration Down: drop-customers-account_number

BEGIN;
ALTER TABLE customers ADD COLUMN identity_number VARCHAR(16) UNIQUE NOT NULL;
COMMIT;
