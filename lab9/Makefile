.PHONY: coverage

start_db:
	docker compose up --build -d

stop_db:
	docker compose down

drop_db: stop_db
	docker volume rm lab9_lab9_db_data

run:
	cd app && go run .
	clear

doc:
	cd app && gomarkdoc -o DOCUMENTATION.md --exclude-dirs ./tests/... ./...

test:
	cd app && go test app/tests

coverage:
	mkdir -p coverage
	cd app && go test \
  		-coverpkg=./config,./database,./internal,./internal/logger,./internal/person,./internal/user \
  		-coverprofile=../coverage/coverage.out \
  		./tests
	cd app && go tool cover -html=../coverage/coverage.out -o ../coverage/coverage.html