alter table products
ADD CONSTRAINT products_category_id_fk FOREIGN KEY (category_id) REFERENCES categories  (id) ON DELETE CASCADE ;

alter table cart_items 
ADD CONSTRAINT cart_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products  (id) ON DELETE CASCADE ,
ADD CONSTRAINT cart_items_shopping_cart_id_fk FOREIGN KEY (shopping_cart_id) REFERENCES shopping_carts (id) ON DELETE CASCADE ;

alter table order_items  
ADD CONSTRAINT order_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
ADD CONSTRAINT order_items_order_id_fk FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE;

alter table shopping_carts 
ADD CONSTRAINT shopping_carts_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE

alter table orders
ADD CONSTRAINT orders_payment_id_fk FOREIGN KEY (payment_id) REFERENCES order_payments  (id) ON DELETE CASCADE,
ADD CONSTRAINT orders_users_id_fk FOREIGN KEY (user_id) REFERENCES users  (id) ON DELETE CASCADE;

alter table order_payments 
ADD CONSTRAINT orders_payment_order_id_fk FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE;

alter user_address 
ADD CONSTRAINT user_adress_user_id_fk FOREIGN KEY (user_id) REFERENCES users  (id) ON DELETE CASCADE;

CREATE DOMAIN stat VARCHAR(20) CHECK (UPPER(VALUE) IN ('PAID', 'WAITING', 'UNPAID'));