CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description text,
    currency VARCHAR(50),
    price bigint,
    total_stock int NOT NULL ,
    is_active boolean default false,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    category_id bigint NOT NULL
);