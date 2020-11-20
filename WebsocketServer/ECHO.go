package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
)

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

// This example demonstrates a trivial echo server.
func main() {
	http.Handle("/123", gin.New())
	http.ListenAndServe(":9999", nil)
	return
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		fmt.Println(request.Form)
	})
	http.ListenAndServe(":9999", nil)
	return
	http.Handle("/echo", websocket.Handler(EchoServer))

	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
