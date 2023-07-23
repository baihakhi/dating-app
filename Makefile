SHELL := /bin/bash
.PHONY: run migrate-up migrate-down test mock

run:
	go run cmd/dating-app/main.go

migrate-create:
	migrate create -ext sql -dir internal/databases/schemas -seq $(action)

migrate-up:
	migrate -path internal/databases/schemas -database "postgresql://root:secret@localhost:5432/dating_app?sslmode=disable" -verbose up

migrate-down:
	migrate -path internal/databases/schemas -database "postgresql://root:secret@localhost:5432/dating_app?sslmode=disable" -verbose down

migrate-force:
	migrate -path internal/databases/schemas -database "postgresql://root:secret@localhost:5432/dating_app?sslmode=disable" force $(version)

test:
	go test -v ./internal/${package}/tests -coverpkg=./internal/${package}

mock:
	mockery --dir=internal/${package} --name=$${package^} --filename=${package}_mock.go --output=internal/${package}/mocks --outpkg=${package}Mock

help:
	@echo "run: run main.go"
	@echo "test: run test by package"
	@echo "test: generate mock files"