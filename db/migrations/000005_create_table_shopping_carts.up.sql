CREATE TABLE IF NOT EXISTS shopping_carts  (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    total int,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    user_id bigint NOT NULL
);