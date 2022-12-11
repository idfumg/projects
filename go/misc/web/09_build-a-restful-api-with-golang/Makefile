run:
	docker-compose up

test:
	go test ./... -v

build:
	docker build -t comments-api .

install-k8s:
	brew install kind # k8s

show-k8s:
	kubectl cluster-info --context kind-kind

run-k8s:
	cat config/deployment.yml | kubectl apply -f -

run-pg:
	docker run \
	-d \
	-e POSTGRES_HOST_AUTH_METHOD=trust \
	-e POSTGRES_USER=postgres \
	-e POSTGRES_PASSWORD=postgres \
	-e POSTGRES_DB=dbname \
	-p 5432:5432 \
	--name pg_instance \
	postgres:12.5-alpine

stop-pg:
	docker stop pg_instance
	docker rm pg_instance