package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {

	//ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10001")
	//LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	//Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10001")
	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)

	if err != nil {
		log.Fatal(err)
	}
	//conn.Write([]byte("hello from UDP client..."))
	defer Conn.Close()
	i := 0
	for {
		_, err := Conn.Write([]byte(strconv.Itoa(i)))
		i++
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * 1)
	}
	//func(conn2 net.Conn) {
	//	buf := bufio.NewReader(conn2)
	//	for {
	//		content, err := buf.ReadString('\n')
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		fmt.Println(content)
	//	}
	//}(conn)
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

//
//func main() {
//	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10001")
//	CheckError(err)
//
//	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
//	CheckError(err)
//
//	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
//	CheckError(err)
//
//	defer Conn.Close()
//	i := 0
//	for {
//		msg := strconv.Itoa(i)
//		i++
//		buf := []byte(msg)
//		_, err := Conn.Write(buf)
//		if err != nil {
//			fmt.Println(msg, err)
//		}
//		time.Sleep(time.Second * 1)
//	}
//}
