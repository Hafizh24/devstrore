CREATE TABLE IF NOT EXISTS cart_items  (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    quantity int,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    product_id bigint NOT NULL,
    shopping_cart_id bigint NOT NULL
);
