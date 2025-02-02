-- Migration Up: drop-customers-account_number
BEGIN;
ALTER TABLE customers DROP COLUMN account_number;

COMMIT;
