.PHONY: help up down clean rebuild exec create-exe create-zip network

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  up           Start the development environment"
	@echo "  down         Stop the development environment"
	@echo "  clean        Stop and remove containers, volumes, and images created by up"
	@echo "  rebuild      Rebuild the development environment without using the cache"
	@echo "  exec         Open a bash shell in the running app container"
	@echo "  create-exe   Build the executable file using pyinstaller"
	@echo "  create-zip   Create a zip file for distribution"
	@echo "  network      Create the common external network"

up: network
	docker-compose -f docker/docker-compose.yml up -d

down:
	docker-compose -f docker/docker-compose.yml down

clean:
	docker-compose -f docker/docker-compose.yml down -v --rmi all
	docker system prune -f
	docker volume prune -f

rebuild: clean network
	docker-compose -f docker/docker-compose.yml build --no-cache
	docker-compose -f docker/docker-compose.yml up -d 

exec:
	docker-compose -f docker/docker-compose.yml exec recovery_app_client /bin/bash

create-exe:
	docker-compose -f docker/docker-compose.yml exec recovery_app_client pyinstaller recovery_adviser.spec

create-zip:
	docker-compose -f docker/docker-compose.yml exec recovery_app_client python scripts/create_zip.py

network:
	docker network create recovery-net || true
