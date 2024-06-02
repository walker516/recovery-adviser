.PHONY: up down migrate seed up-prod down-prod rebuild

up:
	docker-compose -f docker/docker-compose.dev.yml up -d

down:
	docker-compose -f docker/docker-compose.dev.yml down

migrate:
	docker exec -i recovery_mysql_dev mysql recovery < docker/mysql/init/01_create_tables.sql

seed:
	docker exec -i recovery_mysql_dev mysql recovery < docker/mysql/init/02_insert_sample_data.sql

up-prod:
	docker-compose -f docker/docker-compose.prod.yml up -d

down-prod:
	docker-compose -f docker/docker-compose.prod.yml down

clean:
	docker-compose -f docker/docker-compose.dev.yml down --volumes --remove-orphans
	docker system prune -f --volumes
rebuild: clean
	docker-compose -f docker/docker-compose.dev.yml build --no-cache
	docker-compose -f docker/docker-compose.dev.yml up -d
