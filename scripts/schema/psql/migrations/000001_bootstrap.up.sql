-- DROP SCHEMA IF EXISTS public CASCADE;
CREATE SCHEMA IF NOT EXISTS public;

CREATE TABLE IF NOT EXISTS public.users
(
    id serial NOT NULL,
    email text,
    pwd text,
    created_at timestamp without time zone default (now() at time zone 'utc'),
    last_logged_in timestamp without time zone default (now() at time zone 'utc'),
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_email UNIQUE (email)

);

