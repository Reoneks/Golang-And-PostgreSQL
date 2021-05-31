CREATE TABLE IF NOT EXISTS public.users
(
    id bigint NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    first_name name COLLATE pg_catalog."C" NOT NULL,
    last_name name COLLATE pg_catalog."C" NOT NULL,
    email character(45) COLLATE pg_catalog."default" NOT NULL,
    password text COLLATE pg_catalog."default" NOT NULL,
    status smallint NOT NULL DEFAULT 1,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);