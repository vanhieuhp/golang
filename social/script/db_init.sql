create database social;

Create extension if not exists citext;

--
-- migrate -path=./cmd/migrate/migrations -database="postgres://admin:password@localhost/social?sslmode=disable" up