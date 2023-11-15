docker-up:
	docker compose -f ./dockerfiles/docker-compose.yml up -d

docker-down:
	docker compose -f ./dockerfiles/docker-compose.yml down

docker-build:
	docker compose -f ./dockerfiles/docker-compose.yml build

docker-rebuild:
	docker compose -f ./dockerfiles/docker-compose.yml up -d --build

docker-logs:
	docker compose -f ./dockerfiles/docker-compose.yml logs -f

docker-ps:
	docker compose -f ./dockerfiles/docker-compose.yml ps

docker-exec:
	docker compose -f ./dockerfiles/docker-compose.yml exec $(service) bash
