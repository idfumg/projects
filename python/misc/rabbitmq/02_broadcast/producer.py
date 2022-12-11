import pika

connection = pika.BlockingConnection(pika.URLParameters('amqp://user:password@localhost:5672'))
channel = connection.channel()

channel.exchange_declare(exchange='br_exchange', exchange_type='fanout')

for i in range(5):
    msg = 'Hello ' + str(i)
    channel.basic_publish(exchange='br_exchange', routing_key='', body=msg)
    print('[x] Sent %r', msg)

# channel.exchange_delete(exchange='br_exchange', if_unused=False)
connection.close()