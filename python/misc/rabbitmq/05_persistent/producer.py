import pika, random

connection = pika.BlockingConnection(pika.URLParameters('amqp://user:password@localhost:5672'))
channel = connection.channel()
channel.confirm_delivery()

channel.exchange_declare(exchange='pers_exchange', exchange_type='direct', durable=True)

severity = ['Error', 'Warning', 'Info', 'Other']
messages = ['Emsg', 'Wmsg', 'Imsg', 'Omsg']

for i in range(20):
    randomNum = random.randint(0, len(severity) - 1)
    msg = messages[randomNum]
    rk = severity[randomNum]
    try:
        channel.basic_publish(exchange='pers_exchange', routing_key=rk, body=rk, properties=pika.BasicProperties(
            delivery_mode=2,
        ))
        print('[x] Sent ', rk)
    except pika.exceptions.ChannelClosed:
        print('Channel closed')
    except pika.exceptions.ConnectionClosed:
        print('Connection closed')

# channel.exchange_delete(exchange='br_exchange', if_unused=False)
connection.close()