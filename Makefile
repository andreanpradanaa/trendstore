DB_URL=postgresql://root:secret@localhost:5432/trendstore?sslmode=disable

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root trendstore

dropdb:
	docker exec -it postgres dropdb trendstore

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

.PHONY: new_migration postgres createdb dropdb