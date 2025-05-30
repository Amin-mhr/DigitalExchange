CREATE TABLE balances (
                          id SERIAL PRIMARY KEY,
                          exchange_id INTEGER NOT NULL,
                          asset VARCHAR(50) NOT NULL,
                          free FLOAT NOT NULL,
                          locked FLOAT NOT NULL,
                          updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          FOREIGN KEY (exchange_id) REFERENCES exchanges(id) ON DELETE CASCADE
);

CREATE INDEX idx_balances_exchange_id_asset ON balances(exchange_id, asset);