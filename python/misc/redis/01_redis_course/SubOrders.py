import redis
import datetime
import time

r = redis.Redis(password='b5rpiZ9tAKp7tjd9OM48HcllEMs7khMZ', host='redis-17206.c16.us-east-1-3.ec2.cloud.redislabs.com', port=17206, decode_responses=True)

while True:
	received = r.xread({"orders": '$'}, None, 0)

	print(received)

	for result in received:
		data = result[1]
		for tuple in data:
			orderDict = tuple[1];
			print(orderDict)

#	time.sleep(1)

