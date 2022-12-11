run-mq:
	docker run -d --hostname rabbitmq --name test-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management

stop-mq:
	docker stop test-rabbit && docker rm test-rabbit