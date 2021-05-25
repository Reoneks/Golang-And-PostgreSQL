CREATE TABLE IF NOT EXISTS public.products
(
    id integer NOT NULL,
    name name COLLATE pg_catalog."default" NOT NULL,
    created_at date NOT NULL,
    updated_at date NOT NULL,
    created_by integer,
    CONSTRAINT products_pkey PRIMARY KEY (id),
    CONSTRAINT user_key FOREIGN KEY (created_by)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID
);