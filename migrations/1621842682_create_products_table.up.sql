CREATE TABLE IF NOT EXISTS public.products
(
    id bigserial NOT NULL,
    name name COLLATE pg_catalog."C" NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    created_by bigserial NOT NULL,
    CONSTRAINT products_pkey PRIMARY KEY (id),
    CONSTRAINT user_key FOREIGN KEY (created_by)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
);