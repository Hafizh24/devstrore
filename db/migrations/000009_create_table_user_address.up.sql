CREATE TABLE IF NOT EXISTS user_address  (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    address_line text,
    country VARCHAR(50),
    city VARCHAR(50),
    postal_code VARCHAR(50),
    telephone VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    user_id bigint NOT NULL
);