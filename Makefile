lint:
	pre-commit run --all-files -v
run:
	yarn dev
start_db:
	docker-compose up -d
stop_db:
	docker-compose down
