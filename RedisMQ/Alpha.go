package main

import (
	"github.com/adjust/rmq"
	"github.com/gin-gonic/gin"
	"time"
)

func Run() {

	conn := rmq.OpenConnection("tag", "tcp", "127.0.0.1:6379", 2)
	queue := conn.OpenQueue("test0")
	queue.StartConsuming(10, time.Millisecond*100)
	queue.AddConsumerFunc("c1", func(delivery rmq.Delivery) {
		delivery.Reject()
	})
	go func() {
		for range time.Tick(time.Second * 10) {
			queue.ReturnAllRejected()
		}
	}()
	go func() {
		for range time.Tick(time.Millisecond * 100) {
			queue.PublishBytes([]byte("zzz..."))
		}
	}()

	server := gin.Default()
	server.GET("/", func(context *gin.Context) {
		queues := conn.GetOpenQueues()
		state := conn.CollectStats(queues)
		//context.String(0, "", state.GetHtml("", "1"))
		context.Writer.Write([]byte(state.GetHtml("", "1")))
	})
	server.Run(":4444")
}
