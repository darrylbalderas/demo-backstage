lint:
	pre-commit run --all-files -v
run: start_db
	yarn dev
start_db:
	docker-compose up -d
stop_db:
	docker-compose down
