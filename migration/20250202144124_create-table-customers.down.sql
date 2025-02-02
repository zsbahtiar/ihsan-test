-- Migration Down: create-table-customers

BEGIN;
DROP TABLE IF EXISTS customers;
COMMIT;
