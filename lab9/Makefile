start_db:
	docker compose up --build -d

stop_db:
	docker compose down

drop_db: stop_db
	docker volume rm lab9_lab9_db_data

run:
	cd app && go run .
	clear