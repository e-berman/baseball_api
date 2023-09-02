include .env

build:
	docker-compose up --build

db:
	@docker exec -it baseball_api_db_1 psql -U $(POSTGRES_USER) $(POSTGRES_DB)
