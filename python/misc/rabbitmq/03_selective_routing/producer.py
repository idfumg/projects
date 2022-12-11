import pika, random

connection = pika.BlockingConnection(pika.URLParameters('amqp://user:password@localhost:5672'))
channel = connection.channel()

channel.exchange_declare(exchange='logs_exchange', exchange_type='direct')

severity = ['Error', 'Warning', 'Info', 'Other']
messages = ['EMsg', 'WMsg', 'IMsg', 'OMsg']

for i in range(10):
    randomNum = random.randint(0, len(severity) - 1)
    print(randomNum)
    msg = messages[randomNum]
    rk = severity[randomNum]
    channel.basic_publish(exchange='logs_exchange', routing_key=rk, body=msg)
    print('[x] Sent ', msg)

# channel.exchange_delete(exchange='br_exchange', if_unused=False)
connection.close()