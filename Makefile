postgres:
	docker run --name postgres14-bank -p:5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres14-bank createdb --username=root --owner=root bank

dropdb:
	docker exec -it postgres14-bank dropdb bank

migrateup:
	./migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose up

migratedown:
	./migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose down

sqlc:
	./sqlc generate

test:
	go test -v -cover ./...


.PHONY: postgres createdb dropdb migrateup migratedown sqlc test