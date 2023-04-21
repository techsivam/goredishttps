build:
	docker-compose build --build-arg REDIS_HOST=redis --build-arg REDIS_PORT=6379 --build-arg REDIS_PASSWORD= --build-arg REDIS_DB=0
	docker push techsivam16/goredishttps:latest

up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f
