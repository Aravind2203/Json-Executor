import pika

class RabbitMqConnection:

    def __init__(self):
        self.connection=pika.BlockingConnection(pika.URLParameters("amqps://ohwbxpft:Qyp6xxRvgDEIX4VmBQQJgKKjWgihB3vU@armadillo.rmq.cloudamqp.com/ohwbxpft"))
        self.channel=self.connection.channel()

        self.transmit_queue=self.channel.queue_declare(
            "recv_resp",
        )
        self.resp=self.transmit_queue.method.queue

        self.channel.basic_consume(
            self.resp,
            self.on_message,
            True,
            False,
        )
        self.response=None
    def on_message(self,ch,method,props,body):

        self.response=body
    
    def call(self,json):

        self.channel.basic_publish(
            "",
            "rpc_send3",
            json,
            properties=pika.BasicProperties(
                content_type="application/text",
                correlation_id='1',
                reply_to=self.resp
            )
        )
        self.connection.process_data_events(time_limit=None)
        self.connection.close()
        return str(self.response)
        