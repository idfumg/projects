import csv
import redis

r = redis.Redis(password='b5rpiZ9tAKp7tjd9OM48HcllEMs7khMZ', host='redis-17206.c16.us-east-1-3.ec2.cloud.redislabs.com', port=17206, decode_responses=True)


with open("OnlineRetail.csv", encoding='utf-8-sig') as csvf:
	csvReader = csv.DictReader(csvf)

	for row in csvReader:
		r.xadd("orders", row)
