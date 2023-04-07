CREATE TABLE IF NOT EXISTS order_payments  (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    status stat,
    amount bigint,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    order_id bigint NOT NULL 
   );