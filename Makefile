new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

.PHONY: new_migration