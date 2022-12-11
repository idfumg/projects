import redis
import datetime
import time
from pydantic import ValidationError
from Schema import  *

r = redis.Redis(password='b5rpiZ9tAKp7tjd9OM48HcllEMs7khMZ', host='redis-17206.c16.us-east-1-3.ec2.cloud.redislabs.com', port=17206, decode_responses=True)

while True:
	received = r.xread({"orders": '$'}, None, 0)

	for result in received:
		data = result[1]
		for tuple in data:
			orderDict = tuple[1]
			print(orderDict)

			try:
				item = Product(
					StockCode=orderDict['StockCode'],
					Description=orderDict['Description'],
					UnitPrice=orderDict['UnitPrice']
				)

				order = Order(
					InvoiceNo=orderDict['InvoiceNo'],
					Item = item,
					Quantity=orderDict['Quantity'],
					InvoiceDate=datetime.datetime.strptime(orderDict['InvoiceDate'], '%m/%d/%Y %H:%M'),
					CustomerID=orderDict['CustomerID'],
					Country=orderDict['Country']
				)

			except ValidationError as e:
				print(e)
				continue

			print(order.key())
			order.save()

#	time.sleep(1)

