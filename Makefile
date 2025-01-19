.PHONY: up up-services up-kafka down-services down-kafka

up:
	docker compose -f docker-compose-kafka.yaml up -d --build && \
	docker compose -f docker-compose-clickhouse.yaml up -d --build && \
	docker compose -f docker-compose-services.yaml up -d --build

up-services:
	docker compose -f docker-compose-services.yaml up -d --build

up-kafka:
	docker compose -f docker-compose-kafka.yaml up -d --build

up-clickhouse:
	docker compose -f docker-compose-clickhouse.yaml up -d --build

down:
	docker compose -f docker-compose-services.yaml down -v && \
	docker compose -f docker-compose-clickhouse.yaml down -v && \
	docker compose -f docker-compose-kafka.yaml down -v

down-services:
	docker compose -f docker-compose-services.yaml down -v

down-kafka:
	docker compose -f docker-compose-kafka.yaml down -v

down-clickhouse:
	docker compose -f docker-compose-clickhouse.yaml down -v
