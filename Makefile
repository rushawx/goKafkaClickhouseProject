.PHONY: up-services up-kafka down-services down-kafka

up-services:
	docker compose -f docker-compose-services.yaml up -d --build

up-kafka:
	docker compose -f docker-compose-kafka.yaml up -d --build

down-services:
	docker compose -f docker-compose-services.yaml down -v

down-kafka:
	docker compose -f docker-compose-kafka.yaml down -v
