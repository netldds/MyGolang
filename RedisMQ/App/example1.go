package App

import (
	"encoding/json"
	"fmt"
	"github.com/adjust/rmq"
	"net/http"
	"time"
)

type PayLoad struct {
	Header string `json:"header"`
	Time   string `json:"claim"`
}

func FuncOne() {

	//wg := &sync.WaitGroup{}
	//wg.Add(1)

	if false {
		go func() {
			publishConnection := rmq.OpenConnection("Publisher", "tcp", "127.0.0.1:6379", 1)

			dataQueue := publishConnection.OpenQueue("dataQueue")

			//otherQueue := publishConnection.OpenQueue("otherQueue")

			for {
				time.Sleep(time.Millisecond * 2)
				message, _ := json.Marshal(PayLoad{
					Header: "header_data",
					Time:   time.Now().String(),
				})

				dataQueue.PublishBytes(message)
			}

		}()
	}

	if false {
		go func() {
			consumerConnect := rmq.OpenConnection("Consumer", "tcp", "127.0.0.1:6379", 1)

			dataQueue := consumerConnect.OpenQueue("dataQueue")

			dataQueue.StartConsuming(100, time.Second)

			for i := 0; i < 10; i++ {
				name := fmt.Sprintf("consumerID:%d", i)
				dataQueue.AddConsumerFunc(name, func(delivery rmq.Delivery) {
					time.Sleep(time.Second * 2)
					var pd PayLoad
					if err := json.Unmarshal([]byte(delivery.Payload()), &pd); err != nil {
						fmt.Println(err)
						delivery.Reject()
						return
					}
					fmt.Println(pd)
					delivery.Ack()

				})
			}

		}()
	}

	connect := rmq.OpenConnection("handler", "tcp", "127.0.0.1:6379", 1)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		layout := request.FormValue("layout")
		//refresh := request.FormValue("refresh")

		queues := connect.GetOpenQueues()
		stats := connect.CollectStats(queues)
		fmt.Fprint(writer, stats.GetHtml(layout, "1"))
	})

	http.ListenAndServe(":3333", nil)
	//wg.Wait()

}
