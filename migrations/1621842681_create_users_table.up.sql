CREATE TABLE IF NOT EXISTS public.users
(
    id integer NOT NULL,
    first_name name COLLATE pg_catalog."default" NOT NULL,
    last_name name COLLATE pg_catalog."default" NOT NULL,
    email text COLLATE pg_catalog."default" NOT NULL,
    password text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);