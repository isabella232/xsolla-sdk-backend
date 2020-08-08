flatten:
	rm api/flatten_swagger.yaml || true
	swagger flatten api/swagger.yaml > api/flatten_swagger.yaml

validate: flatten
	swagger validate api/flatten_swagger.yaml

gen: validate
	swagger generate server -f api/flatten_swagger.yaml -t internal/server --exclude-main

build: gen
	rm cmd/web-api/web-api || true
	go build -o cmd/web-api xsolla-sdk-backend/cmd/web-api

run: build
	scripts/run.sh

debug: build
	scripts/debug.sh

run-tests:
	scripts/run-tests.sh

lint:
	golangci-lint run ./...

# Check migrations and fixtures
migrations-test:
	rm tools/migrate/migrate || true
	go build -o tools/migrate/migrate tools/migrate/migrate.go
	./tools/migrate/migrate -user root -password root -host 127.0.0.1 -port 3307 -dir migrations/migrations
	./tools/migrate/migrate -user root -password root -host 127.0.0.1 -port 3307 -dir migrations/fixtures