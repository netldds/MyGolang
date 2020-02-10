package main

import (
	"encoding/json"
	"fmt"
	"github.com/adjust/rmq"
	"io/ioutil"
	"path"
	"time"
)

type Payload struct {
	FilePath  string `json:"file_path"`
	OutputDir string `json:"output_dir"`
}

var directory = "/home/dx/GoWorkBench/src/thumbnailerDEV/testdata2"
var filename []string

func main() {
	findAllFiles(&filename)

	if true {
		go func() {
			connection := rmq.OpenConnection("Publisher", "tcp", "127.0.0.1:6379", 1)
			TGAQueue := connection.OpenQueue("tga_queue")
			for _, i := range filename {
				data, _ := json.Marshal(Payload{FilePath: i, OutputDir: path.Dir(i)})
				TGAQueue.PublishBytes(data)
			}
		}()
	}
	if false {
		go func() {
			connection := rmq.OpenConnection("Consumer", "tcp", "127.0.0.1:6379", 1)
			tgaqueue := connection.OpenQueue("tga_queue")
			tgaqueue.StartConsuming(20, time.Second)
			for i := 0; i < 10; i++ {
				tgaqueue.AddConsumerFunc(fmt.Sprintf("consumerID%d", i), func(delivery rmq.Delivery) {
					var py Payload
					json.Unmarshal([]byte(delivery.Payload()), &py)
					fmt.Println(py)
					delivery.Ack()
				})

			}
		}()
	}
	select {}
}
func findAllFiles(fn *[]string) {

	finfo, err := ioutil.ReadDir(directory)
	if err != nil {
		return
	}
	for _, v := range finfo {
		//fmt.Println(v.Name())
		*fn = append(*fn, path.Join(directory, v.Name()))
	}
}
