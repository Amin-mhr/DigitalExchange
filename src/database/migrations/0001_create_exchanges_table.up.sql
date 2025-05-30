CREATE TABLE exchanges (
                           id SERIAL PRIMARY KEY,
                           name VARCHAR(255) NOT NULL,
                           api_key VARCHAR(255) NOT NULL,
                           api_secret VARCHAR(255) NOT NULL,
                           created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_exchanges_name ON exchanges(name);