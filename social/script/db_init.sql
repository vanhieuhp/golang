create database social;

--
-- migrate -path=./cmd/migrate/migrations -database="postgres://admin:password@localhost/social?sslmode=disable" up