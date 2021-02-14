include app.env
export $(shell sed 's/=.*//' app.env)

.EXPORT_ALL_VARIABLES:
PKG ?=-v /vendor
ALL_PACKAGES=$(shell go list ./... | grep $(PKG))
CVPKG:=$(shell go list ./... | grep -v mocks | tr '\n' ',' | sed 's/,$$//')

list-pkg:
	: $(ALL_PACKAGES)

gen-mocks:
	mockery --dir=domain --all --output=domain/mock

gen-migration:
	migrate create -ext sql -dir db/migrations -seq ${migration-name}

db.create:
	@echo "creating database $(DB_NAME)..."
	@PGPASSWORD=${DB_PASSWORD} createdb -h $(DB_HOST) -U $(DB_USER) -w -Eutf8 $(DB_NAME)

db.drop:
	@echo "dropping database $(DB_NAME)..."
	@PGPASSWORD=${DB_PASSWORD} dropdb --if-exists -h $(DB_HOST) -U $(DB_USER) -w $(DB_NAME)

db.migrate:
	migrate -database "postgres://$(DB_USER):${DB_PASSWORD}@$(DB_HOST)/$(DB_NAME)?sslmode=disable" -lock-timeout 30 -path db/migrations up

db.rollback:
	migrate -database "postgres://$(DB_USER):${DB_PASSWORD}@$(DB_HOST)/$(DB_NAME)?sslmode=disable" -lock-timeout 30 -path db/migrations down

testdb.create:
	@echo "creating database $(TEST_DB_NAME)..."
	@PGPASSWORD=${TEST_DB_PASSWORD} createdb -h $(TEST_DB_HOST) -U $(TEST_DB_USER) -w -Eutf8 $(TEST_DB_NAME)

testdb.drop:
	@echo "dropping database $(TEST_DB_NAME)..."
	@PGPASSWORD=${TEST_DB_PASSWORD} dropdb --if-exists -h $(TEST_DB_HOST) -U $(TEST_DB_USER) -w $(TEST_DB_NAME)

testdb.migrate:
	migrate -database "postgres://$(TEST_DB_USER):${TEST_DB_PASSWORD}@$(TEST_DB_HOST)/$(TEST_DB_NAME)?sslmode=disable" -lock-timeout 30 -path db/migrations up

testdb.rollback:
	migrate -database "postgres://$(TEST_DB_USER):${TEST_DB_PASSWORD}@$(TEST_DB_HOST)/$(TEST_DB_NAME)?sslmode=disable" -lock-timeout 30 -path db/migrations down

test-full: testdb.drop testdb.create testdb.migrate
	@APP_ENV=test go test -v ./... -cover -race -coverprofile=coverage.out -covermode=atomic
	@go tool cover -html=coverage.out

run-watch:
	reflex -r "\.go" -s -- sh -c "go run main.go"