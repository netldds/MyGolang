package main

import (
	"bufio"
	"fmt"
	"github.com/golang/glog"
	"log"
	"net/http"
	"os"
	"time"
)

//GIN FrameWork APP
func main() {

	//DefaultAPI()
	//	Middleware()
	//renderData()
	//ClockStream()
	//ServingData()
	//htmlRendering()
	//middleRun()
	//UsingBasicAuth()
	//Goroutinesmiddleware()
	//Encrypt()
	//AutoCert()
	//RunMultiServices()
	//GracefulShutdown()
	//BindFormData()
	//ServerPush()
	//queryDB()
	//fmt.Println(len("9f1431e28609f22fb5e6fcd9f713e8d6"))
	//WebSocketImplementation()
	//StartSrv()
	//input()
	//ForwardProxyStart()
	//APITestforPJT()
	DownloadFile()
}
func input() {
	rd := bufio.NewReader(os.Stdin)
	for {
		delim, err := rd.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(delim)
	}
}
func tt() {
	http.HandleFunc("/", func(res http.ResponseWriter, request *http.Request) {
		res.Write([]byte("ok\n"))
		res.WriteHeader(200)
	})
	srv := http.Server{
		Addr:    ":9000",
		Handler: http.DefaultServeMux,
	}
	//ctx, cancel := context.WithCancel(context.Background())
	go func() {
		//ctx,cancel:=context.WithTimeout(context.Background(), time.Second*2)
		ticker := time.NewTicker(time.Second * 5)
		for {
			select {
			case t := <-ticker.C:
				fmt.Println(t)
				if err := srv.Shutdown(context.Background()); err != nil {
					// Error from closing listeners, or context timeout:
					glog.Infof("HTTP server Shutdown: %v", err)
				}
			}
		}
	}()
	srv.ListenAndServe()
}
