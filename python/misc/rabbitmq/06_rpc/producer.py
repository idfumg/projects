import pika
import uuid

class RPCConsumer:
    def __init__(self) -> None:
        self.connection = pika.BlockingConnection(pika.URLParameters('amqp://user:password@localhost:5672'))
        self.channel = self.connection.channel()
        self.producer_queue_name = 'rpc_producer_queue'
        self.consumer_queue_name = 'rpc_consumer_queue'
        self.channel.queue_declare(queue=self.producer_queue_name, exclusive=True)
        self.channel.basic_consume(queue=self.producer_queue_name, on_message_callback=self.on_response, auto_ack=True)
        self.response = None
        self.correlation_id = str(uuid.uuid4())

    def on_response(self, ch, method, props, body):
        if self.correlation_id == props.correlation_id:
            self.response = body

    def call(self, s):
        self.channel.basic_publish(exchange='', routing_key=self.consumer_queue_name, 
            properties=pika.BasicProperties(
                reply_to=self.producer_queue_name,
                correlation_id=self.correlation_id,
            ),
            body=s)

        while self.response is None:
            self.connection.process_data_events()

        return self.response

rpc = RPCConsumer()
msg = '222'
print('Requesting', msg)
resp = rpc.call(msg)
print('Got the response', str(resp))