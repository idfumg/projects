run:
	ADMIN_PASSWORD=secret go run main.go

r1:
	curl -v localhost:8080/coasters

r2:
	curl -v localhost:8080/coasters -H "Content-Type: application/json" -X POST -d '{"name": "Taron", "in_park": "Phantasialand", "height": 30, "manufacturer": "Intamin"}'

r3:
	curl -v localhost:8080/coasters -H "Content-Type: application/json" -X PUT -d '{"name": "Taron2", "in_park": "Phantasialand", "height": 30, "manufacturer": "Intamin", "id": "1659970442854052000"}'

r4:
	curl -v localhost:8080/admin -u admin:secret