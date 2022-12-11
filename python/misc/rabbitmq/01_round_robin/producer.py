import pika

connection = pika.BlockingConnection(pika.URLParameters('amqp://user:password@localhost:5672'))
channel = connection.channel()
channel.queue_declare(queue='hello')

channel.basic_publish(exchange='', routing_key='hello', body='hello_world msg #5')
print('[x] Sent Hello World')

connection.close()