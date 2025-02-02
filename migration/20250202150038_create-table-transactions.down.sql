-- Migration Down: create-table-transactions

BEGIN;
DROP TABLE IF EXISTS transactions;
COMMIT;
