import pika, random

def on_request(ch, method, props, body):
    reply_queue_name = props.reply_to
    correlation_id = props.correlation_id
    response = 'Response is ' + str(body)

    ch.basic_publish(exchange='', routing_key=reply_queue_name, 
        properties=pika.BasicProperties(
            correlation_id=correlation_id,
        ),
        body=str(response))
    ch.basic_ack(delivery_tag=method.delivery_tag)

connection = pika.BlockingConnection(pika.URLParameters('amqp://user:password@localhost:5672'))
channel = connection.channel()
queue_name = 'rpc_consumer_queue'
result = channel.queue_declare(queue=queue_name, durable=True)
channel.basic_qos(prefetch_count=1)
channel.basic_consume(queue=queue_name, on_message_callback=on_request)
print('Awaiting RPC requests')
channel.start_consuming()