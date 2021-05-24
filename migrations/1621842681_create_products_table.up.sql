CREATE TABLE IF NOT EXISTS public.products
(
    id integer NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    created_at date NOT NULL,
    updated_at date NOT NULL,
    CONSTRAINT products_pkey PRIMARY KEY (id)
);