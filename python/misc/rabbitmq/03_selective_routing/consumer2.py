import pika, os, sys

def main():
    connection = pika.BlockingConnection(pika.URLParameters('amqp://user:password@localhost:5672'))
    channel = connection.channel()

    channel.exchange_declare(exchange='logs_exchange', exchange_type='direct')

    result = channel.queue_declare(queue='', exclusive=True)
    queue_name = result.method.queue
    print('Subscriber queue name:', queue_name)

    severity = ['Error', 'Warning', 'Info']

    channel.queue_bind(exchange='logs_exchange', queue=queue_name, routing_key='Warning')

    def callback(ch, method, properties, body):
        print('[x] received', body)

    channel.basic_consume(queue=queue_name, on_message_callback=callback, auto_ack=True)
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