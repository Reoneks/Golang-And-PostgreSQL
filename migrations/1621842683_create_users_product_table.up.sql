CREATE TABLE IF NOT EXISTS public.users_products
(
    user_id bigserial NOT NULL,
    product_id bigserial NOT NULL,
    CONSTRAINT product_key FOREIGN KEY (product_id)
        REFERENCES public.products (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT user_key FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
);