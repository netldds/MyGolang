package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"time"
)

func Checkerr(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(0)
	}
}

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}

//支持的RPC运算方法，所有方法以结构为基础
type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero.")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
func main() {
	arith := new(Arith)
	rpc.Register(arith)
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8080")

	// accept connection on port
	defer ln.Close()
	// run loop forever (or until ctrl-c)
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		//go handleConnection(conn)
		//rpc
		//go rpc.ServeConn(conn)
		go jsonrpc.ServeConn(conn)
	}

}
func handleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("remote client from" + remoteAddr)

	reader := bufio.NewReader(conn)
	for {
		data, err3 := reader.ReadString('\n')
		if err3 != nil {
			break
		}
		log.Printf("%v -> %s", conn.RemoteAddr(), data)
	}
	//for {
	//	message := make([]byte, 1024)
	//	i, err2 := conn.Read(message)
	//	if err2 != nil {
	//		break
	//	}
	//	log.Printf("%v -> %s",i,string(message))
	//}
	//scanner := bufio.NewScanner(conn)
	//for {
	//	ok := scanner.Scan()
	//	if !ok {
	//		break
	//	}
	//	fmt.Println(scanner.Text())
	//}

}
func tcp_test() {
	//server
	go func() {
		net, _ := net.Listen("tcp", ":6666")
		fmt.Println(net.Addr())
		for {
			conn, err := net.Accept()
			fmt.Println(" server conn accepted")
			if err != nil {
				fmt.Println(err)
				continue
			}
			go func() {
				rd := bufio.NewReader(conn)
				fmt.Printf("server local addr  %v, server remote addr %v \n", conn.LocalAddr(), conn.RemoteAddr())
				for {
					bs := make([]byte, 1024)
					n, err := rd.Read(bs)
					if err != nil {
						fmt.Println(err)
						break
					}
					fmt.Println(string(bs[:n]))
				}
			}()
		}
	}()
	//client
	time.Sleep(time.Second)
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:6666")
	conn, _ := net.DialTCP("tcp", nil, addr)
	fmt.Printf("client local addr %v , client remote addr %v \n ", conn.LocalAddr(), conn.RemoteAddr())
	var count int
	for {
		n, err := conn.Write([]byte("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n"))
		count += n
		fmt.Printf("write %v bytes,err :%v,count:%v \n", n, err, count)
		time.Sleep(time.Second)
	}
	select {}
}
