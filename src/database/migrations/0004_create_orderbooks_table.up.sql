CREATE TABLE order_books (
                             id SERIAL PRIMARY KEY,
                             exchange_id INTEGER NOT NULL,
                             symbol VARCHAR(50) NOT NULL,
                             bid_price FLOAT NOT NULL,
                             bid_quantity FLOAT NOT NULL,
                             ask_price FLOAT NOT NULL,
                             ask_quantity FLOAT NOT NULL,
                             timestamp TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             FOREIGN KEY (exchange_id) REFERENCES exchanges(id) ON DELETE CASCADE
);

CREATE INDEX idx_order_books_exchange_id_symbol ON order_books(exchange_id, symbol);