DROP INDEX IF EXISTS idx_exchanges_name;

ALTER TABLE orders DROP CONSTRAINT IF EXISTS fk_orders_exchange_id;
ALTER TABLE balances DROP CONSTRAINT IF EXISTS fk_balances_exchange_id;
ALTER TABLE order_books DROP CONSTRAINT IF EXISTS fk_order_books_exchange_id;

DROP TABLE IF EXISTS exchanges;