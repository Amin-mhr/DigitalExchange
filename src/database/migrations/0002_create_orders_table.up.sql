CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        exchange_id INTEGER NOT NULL,
                        client_order_id VARCHAR(255) NOT NULL,
                        symbol VARCHAR(50) NOT NULL,
                        side VARCHAR(10) NOT NULL,
                        type VARCHAR(20) NOT NULL,
                        quantity FLOAT NOT NULL,
                        price FLOAT NOT NULL,
                        time_in_force VARCHAR(10) NOT NULL,
                        status VARCHAR(20) NOT NULL,
                        exchange_order_id VARCHAR(255),
                        created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        FOREIGN KEY (exchange_id) REFERENCES exchanges(id) ON DELETE CASCADE
);

CREATE INDEX idx_orders_exchange_id_symbol ON orders(exchange_id, symbol);