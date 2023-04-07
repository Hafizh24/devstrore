CREATE TABLE IF NOT EXISTS orders  (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    quantity int,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    payment_id bigint NOT NULL,
    user_id bigint NOT NULL 
);