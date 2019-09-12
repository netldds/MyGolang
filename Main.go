package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

//https://golang.org/cmd/cgo/
var data = strings.NewReader("data from strings")

func main() {
	g := gin.New()
	pwd, _ := os.Getwd()
	g.Static("/res", pwd)
	g.Run(":8888")

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
