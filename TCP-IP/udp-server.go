package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	udpaddr, err := net.ResolveUDPAddr("udp", ":10001")
	con, err := net.ListenUDP("udp", udpaddr)
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	buff := make([]byte, 1024)
	for {
		n, addr, err := con.ReadFrom(buff)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Receive:", string(buff[:n]), "from", addr)
		con.WriteTo([]byte("Hello from client."), addr)
	}

}

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

//func main() {
//	/* Lets prepare a address at any address at port 10001*/
//	ServerAddr, err := net.ResolveUDPAddr("udp", ":10001")
//	CheckError(err)
//
//	/* Now listen at selected port */
//	ServerConn, err := net.ListenUDP("udp", ServerAddr)
//	CheckError(err)
//	defer ServerConn.Close()
//
//	buf := make([]byte, 1024)
//
//	for {
//		n, addr, err := ServerConn.ReadFromUDP(buf)
//		fmt.Println("Received ", string(buf[0:n]), " from ", addr)
//
//		if err != nil {
//			fmt.Println("Error: ", err)
//		}
//	}
//}
