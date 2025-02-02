-- Migration Down: create-table-customers

BEGIN;
DROP TABLE IF NOT EXISTS customers;
COMMIT;
