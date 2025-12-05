CREATE TABLE if not exists users
(
    id         BIGSERIAL PRIMARY KEY,
    username   VARCHAR(255) UNIQUE         not null,
    email      citext UNIQUE               not null,
    password   pg_catalog.bytea            not null,
    created_at timestamp(0) with time zone not null default now()
);