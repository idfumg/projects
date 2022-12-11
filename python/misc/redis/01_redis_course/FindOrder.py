import datetime
from Schema import *

results = Order.find(Order.Quantity > 200).all()

for result in results:
	print(result)
	print("\n")
