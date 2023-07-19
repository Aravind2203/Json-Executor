package main

import (
	"context"
	"engine/util"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	conn, err := amqp.Dial("amqps://ohwbxpft:Qyp6xxRvgDEIX4VmBQQJgKKjWgihB3vU@armadillo.rmq.cloudamqp.com/ohwbxpft")
	util.HandleError(err)
	defer conn.Close()

	ch, err := conn.Channel()
	util.HandleError(err)
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"rpc_send3",
		false,
		false,
		false,
		false,
		nil,
	)
	util.HandleError(err)

	msg, _ := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	go func() {
		for d := range msg {
			fmt.Println(string(d.Body))
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			resp := util.MainExecute(d.Body)
			err := ch.PublishWithContext(ctx,
				"",
				d.ReplyTo,
				false,
				false,
				amqp.Publishing{
					ContentType:   "application/text",
					Body:          []byte(resp),
					CorrelationId: "1",
				},
			)
			util.HandleError(err)
			cancel()
		}
	}()

	var forever chan struct{}
	<-forever
}
