-- Migration Down: create-table-accounts

BEGIN;
DROP TABLE IF EXISTS accounts;
COMMIT;
