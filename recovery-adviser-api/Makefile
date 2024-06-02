.PHONY: help up down migrate seed up-prod down-prod clean rebuild

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  up           Start the development environment"
	@echo "  down         Stop the development environment"
	@echo "  migrate      Run database migrations in the development environment"
	@echo "  seed         Seed the database with sample data in the development environment"
	@echo "  up-prod      Start the production environment"
	@echo "  down-prod    Stop the production environment"
	@echo "  clean        Stop and remove containers, networks, volumes, and images created by up and up-prod"
	@echo "  rebuild      Rebuild the development environment without using the cache"

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
