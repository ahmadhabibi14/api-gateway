docker-dev:
	docker-compose -f docker-compose.dev.yml up -d

docker-prod:
	docker-compose -f docker-compose.yml up -d