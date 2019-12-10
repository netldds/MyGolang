package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
)

/*
https://www.rabbitmq.com/tutorials/tutorial-three-go.html
*/
var data = `[{"discount":"0.9","discountedPrice":11720,"id":1,"originPrice":13020,"specValue":"1年","title":"精英版标准套餐"},{"id":2,"parent":1,"specValue":"5人","title":"同时操作人数"},{"id":3,"parent":1,"specValue":"200G","title":"应用空间"},{"discountedPrice":0,"id":4,"originPrice":600,"specValue":"1年","title":"模型装配服务"},{"discountedPrice":0,"id":5,"originPrice":240,"specValue":"1年","title":"二维图纸预览服务"}]`

func Client() {
	conn, err := amqp.Dial("amqp://dx:dx123@192.168.99.102:5672/")
	fmt.Println(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	fmt.Println(err, "Failed to open a channel")
	defer ch.Close()

	//err = ch.ExchangeDeclare(
	//	"logs",   // name
	//	"fanout", // type
	//	true,     // durable
	//	false,    // auto-deleted
	//	false,    // internal
	//	false,    // no-wait
	//	nil,      // arguments
	//)
	//fmt.Println(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	fmt.Println(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,                // queue name
		"status",              // routing key
		"OrderStatusExchange", // exchange
		false,
		nil,
	)
	fmt.Println(q.Name)
	fmt.Println(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	fmt.Println(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf(" [x] %s", d.Body)
		}
	}()

	fmt.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
