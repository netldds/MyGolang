package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

//https://golang.org/cmd/cgo/
var data = strings.NewReader("data from strings")

func main() {
	go func() {
		for {
			time.Sleep(time.Second)
			printUsage()
		}
	}()
	//var overall [][]int
	//for i := 0; i < 4; i++ {
	//	a := make([]int, 999999)
	//	overall = append(overall, a)
	//	printUsage()
	//	time.Sleep(time.Millisecond * 2)
	//}
	reader := getReader()
	log.Println(reader)
	time.Sleep(time.Second * 2)
	reader = nil
	select {}
}
func printUsage() {
	stats := runtime.MemStats{}
	runtime.ReadMemStats(&stats)
	//fmt.Printf("Sys MB %v	", stats.Sys>>20)
	//fmt.Printf("PauseTotalNs %v	", stats.PauseTotalNs)
	//fmt.Printf("Alloc MB %v	", stats.Alloc>>20)
	//fmt.Printf("Lookups %v	", stats.Lookups)
	//fmt.Printf("HeapObjects %v	", stats.HeapObjects)
	//fmt.Printf("live object %v\n", stats.Mallocs-stats.Frees)
	//fmt.Printf("live object %v\n", stats.HeapInuse)

	n, ok := runtime.MemProfile(nil, false)
	var p []runtime.MemProfileRecord
	for {
		p = make([]runtime.MemProfileRecord, n+50)
		n, ok = runtime.MemProfile(p, false)
		if ok {
			p = p[0:n]
			break
		}
	}
	var total runtime.MemProfileRecord
	for i := range p {
		r := &p[i]
		total.AllocBytes += r.AllocBytes
		total.AllocObjects += r.AllocObjects
		total.FreeBytes += r.FreeBytes
		total.FreeObjects += r.FreeObjects
	}

	fmt.Printf("%d in use objects (%d in use bytes) | Alloc: %d TotalAlloc: %d\n",
		total.InUseObjects(), total.InUseBytes(), stats.Alloc, stats.TotalAlloc)
}
func getReader() io.Reader {
	f, err := os.OpenFile("tmp/1MB.gif", os.O_RDONLY, 0775)
	if err != nil {
		log.Println(err)
	}
	return f
}
func serve() {
	g := gin.Default()
	g.GET("/", func(context *gin.Context) {
		//p, _ := os.Getwd()
		//application/octet-stream
		//context.File(path.Join(p, "taishan"))
		context.DataFromReader(http.StatusOK, data.Size(), "application/octet-stream", data, make(map[string]string))
		//context.Data()
	})
	g.Run(":8080")

}
func bb() {
	if true {
		fmt.Println("1")
		if true {

		}
		fmt.Println("1")
	}
	fmt.Println()
}
