build:
	@go build -o bin/app

run: build
	@./bin/app

redis-start:
	docker run -d -p 6379:6379 --name myredis redis

redis-stop:
	docker stop myredis
	docker rm myredis