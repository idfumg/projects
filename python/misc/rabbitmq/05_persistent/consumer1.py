import pika, os, sys

def main():
    connection = pika.BlockingConnection(pika.URLParameters('amqp://user:password@localhost:5672'))
    channel = connection.channel()

    channel.exchange_declare(exchange='pers_exchange', exchange_type='direct', durable=True)

    queue_name = 'task_queue'
    result = channel.queue_declare(queue=queue_name, durable=True)
    print('Subscriber queue name:', queue_name)

    severity = ['Error', 'Warning', 'Info', 'Other']
    for s in severity:
        channel.queue_bind(exchange='pers_exchange', queue=queue_name, routing_key=s)

    def callback(ch, method, properties, body):
        print('[x] received', str(body))
        channel.basic_ack(delivery_tag=method.delivery_tag)

    channel.basic_qos(prefetch_count=1)
    channel.basic_consume(queue=queue_name, on_message_callback=callback)
    print('[*] waiting for messages. To exit press Ctrl-C')
    channel.start_consuming()

if __name__ == '__main__':
    try:
        main()
    except KeyboardInterrupt:
        print('Interrupted')
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)