package main

import (
	"MyGolang/GORoutine"
	"github.com/streadway/amqp"
	"log"
)

var data = `[{"discount":"0.9","discountedPrice":11720,"id":1,"originPrice":13020,"specValue":"1年","title":"精英版标准套餐"},{"id":2,"parent":1,"specValue":"5人","title":"同时操作人数"},{"id":3,"parent":1,"specValue":"200G","title":"应用空间"},{"discountedPrice":0,"id":4,"originPrice":600,"specValue":"1年","title":"模型装配服务"},{"discountedPrice":0,"id":5,"originPrice":240,"specValue":"1年","title":"二维图纸预览服务"}]`

func main() {
	GORoutine.CheckCost()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func AA() {
	conn, err := amqp.Dial("amqp://dx:dx123@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
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
	//failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,                // queue name
		"status",              // routing key
		"OrderStatusExchange", // exchange
		false,
		nil,
	)
	log.Println(q.Name)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
