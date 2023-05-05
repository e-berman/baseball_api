include .env

build:
	docker-compose up --build

run:
	docker compose exec rest_api ./main

db:
	@docker exec -it baseball_api_db_1 psql -U $(POSTGRES_USER) $(POSTGRES_DB)

get_path:
	docker compose exec rest_api readlink -f stats.csv 
