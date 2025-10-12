createdb:
	docker exec -it postgres_container createdb -U admin simple_bank
dropdb:
	docker exec -it postgres_container dropdb -U admin simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://admin:password@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://admin:password@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dropdb migrateup migratedown sqlc test