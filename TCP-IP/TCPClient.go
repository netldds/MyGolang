package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
)

func CheckErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}

func main() {

	//rpcclient()

	tcpaddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	CheckErr(err)
	tcpconn, err2 := net.DialTCP("tcp", nil, tcpaddr)
	CheckErr(err2)
	defer tcpconn.Close()

	go handleMessageFromServer(tcpconn)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">")
		text, _ := reader.ReadString('\n')
		//tcpconn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		_, err := tcpconn.Write([]byte(text))
		if text == "q\n" {
			os.Exit(0)
		}
		if err != nil {
			fmt.Println("Error writing to stream.")
			break
		}
	}

}
func handleMessageFromServer(conn *net.TCPConn) {
	reply := make([]byte, 1024)
	for {
		_, err := conn.Read(reply)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Print(reply)
	}
}
func rpcclient() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8080")
	CheckErr(err)
	var reply int
	err = client.Call("Arith.Multiply", &Args{A: 11, B: 11}, &reply)
	CheckErr(err)
	fmt.Printf("Arith: %d*%d=%d\n", 11, 11, reply)
}
