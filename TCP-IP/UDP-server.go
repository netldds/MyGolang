package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

/*https://colobu.com/2016/10/19/Go-UDP-Programming/
 */
var remote *net.UDPAddr

func UDPRead(listenConn *net.UDPConn) {
	go UDPWrite(listenConn)
	for {
		b := make([]byte, 1<<10)
		n, remoteAddr, err := listenConn.ReadFromUDP(b)
		if err != nil {
			panic(err)
		}
		remote = remoteAddr
		fmt.Printf("remote:%v -> %v\n", remote, string(b[:n]))
	}
}

func UDPWrite(listenCoon *net.UDPConn) {
	scaner := bufio.NewScanner(os.Stdin)
	for {
		if scaner.Scan() && remote != nil {
			b := scaner.Bytes()
			fmt.Printf("local->:%v remote:%v\n", string(b), remote)
			b = append(b, byte(10))
			_, err := listenCoon.WriteToUDP(b, remote)
			if err != nil {
				panic(err)
			}
		}
	}
}

func Server() {
	/*
		测试 udp 包行为
	*/
	addr, _ := net.ResolveUDPAddr("udp", ":9988")
	listenConn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	go func() {
		UDPRead(listenConn)
	}()
	select {}
}
