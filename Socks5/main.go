package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"net"
	"os"
	"strconv"
	"strings"
)

/*
2021.02.01
zhangliu
https://www.rfc-editor.org/rfc/pdfrfc/rfc1928.txt.pdf
https://www.rfc-editor.org/info/rfc1929
*/

const (
	VERSION = 5
)

func bytesToint(b []byte) int {
	unit := 256
	sum := 0
	for n := len(b); n > 0; n-- {
		sum += int(b[len(b)-n]) * int(math.Pow(float64(unit), float64(n-1)))
	}
	return sum
}
func intTobytes(n int, length int) []byte {
	unit := 256
	b := make([]byte, length)
	b[0] = byte(n % unit)
	for i := n / unit; i >= unit; i /= unit {
		b = append(b, byte(i%unit))
	}
	bigB := make([]byte, len(b))
	for n, v := range b {
		bigB[len(b)-n-1] = v
	}
	return bigB
}

var ErrMethod = byte(255)

func VerifyPassword(conn net.Conn) bool {
	verByte := make([]byte, 1)
	n, err := conn.Read(verByte)
	if err != nil || n == 0 {
		panic(err)
	}
	if uint(verByte[0]) != 1 {
		panic(verByte)
	}
	ulenByte := make([]byte, 1)
	n, err = conn.Read(ulenByte)
	if err != nil || n == 0 {
		panic(err)
	}
	if uint(ulenByte[0]) < 1 {
		panic(ulenByte)
	}
	unameByte := make([]byte, uint(ulenByte[0]))
	n, err = conn.Read(unameByte)
	if err != nil || n == 0 {
		panic(err)
	}
	uname := string(unameByte)
	plen := make([]byte, 1)
	n, err = conn.Read(plen)
	if err != nil || n == 0 {
		panic(err)
	}
	if uint(plen[0]) < 1 {
		panic(plen)
	}
	passwdByte := make([]byte, uint(plen[0]))
	n, err = conn.Read(passwdByte)
	if err != nil || n == 0 {
		panic(err)
	}
	passwd := string(passwdByte)
	fmt.Printf("user:%v\rpassed:%v\n", uname, passwd)

	return true
}
func TransformTraffic(src, dst net.Conn, close chan error) {
	go func() {
		for {
			n, err := io.Copy(dst, src)
			if err != nil || n == 0 {
				close <- err
			}
		}
	}()
	go func() {
		for {
			n, err := io.Copy(src, dst)
			if err != nil || n == 0 {
				close <- err
			}
		}
	}()
}
func HandleConn(conn net.Conn) {
	verByte := make([]byte, 1)
	nmethods := make([]byte, 1)
	n, err := conn.Read(verByte)
	if err != nil {
		panic(err)
	}
	//版本
	if uint(verByte[0]) != VERSION {
		panic(verByte[0])
	}
	n, err = conn.Read(nmethods)
	if err != nil {
		panic(err)
	}
	methods := []int{}
	for i := 0; i < int(nmethods[0]); i++ {
		method := make([]byte, 1)
		n, err = conn.Read(method)
		if err != nil {
			panic(err)
		}
		methods = append(methods, int(method[0]))
	}
	flag := false
	for _, v := range methods {
		switch v {
		case 0:
			fmt.Println("NO AUTHEN")
		case 1:
			fmt.Println("GSSAPI")
		case 2:
			fmt.Println("USERNAME/PASSWORD")
		default:
			panic(verByte)
		}
		if v == 2 {
			flag = true
		}
	}
	b := make([]byte, 0)
	b = append(b, byte(5))
	if !flag {
		fmt.Println("flag")
		b = append(b, ErrMethod)
		conn.Write(b)
		os.Exit(1)
	}
	//选择验证
	b = append(b, byte(2))
	conn.Write(b)
	closeBytes := make([]byte, 0)
	closeBytes = append(closeBytes, byte(1))
	//验证
	if !VerifyPassword(conn) {
		closeBytes = append(closeBytes, byte(1))
		conn.Write(closeBytes)
		os.Exit(1)
	}
	//验证成功
	closeBytes = append(closeBytes, byte(0))
	n, err = conn.Write(closeBytes)
	if err != nil || n == 0 {
		panic(n)
	}
	headBytes := make([]byte, 4)
	n, err = conn.Read(headBytes)
	if err != nil || n == 0 {
		panic(err)
	}
	ver, cmd, _, atyp := int(headBytes[0]), int(headBytes[1]), int(headBytes[2]), int(headBytes[3])
	if ver != VERSION {
		panic(ver)
	}
	addrStr := ""
	domain := ""
	switch atyp {
	case 1:
		fmt.Println("IP V4 address")
		dstAddrBytes := make([]byte, 4)
		n, err = conn.Read(dstAddrBytes)
		if err != nil || n == 0 {
			panic(err)
		}
		addrStr = fmt.Sprintf("%v.%v.%v.%v", dstAddrBytes[0], dstAddrBytes[1], dstAddrBytes[2], dstAddrBytes[3])
	case 3:
		fmt.Println("DOMAINNAME")
		hostLenByte := make([]byte, 1)
		n, err = conn.Read(hostLenByte)
		if err != nil || n == 0 {
			panic(err)
		}
		hostBytes := make([]byte, int(hostLenByte[0]))
		n, err = conn.Read(hostBytes)
		if err != nil || n == 0 {
			panic(err)
		}
		domain = string(hostBytes)
		addrs, err := net.LookupHost(domain)
		if len(addrs) == 0 || err != nil {
			panic(err)
		}
		addrStr = addrs[0]
	case 4:
		fmt.Println("IP V6 address")
		dstAddrBytes := make([]byte, 16)
		n, err = conn.Read(dstAddrBytes)
		if err != nil || n == 0 {
			panic(err)
		}
		addrStr = "["
		for i := 0; i < len(dstAddrBytes)/2; i++ {
			addrStr += fmt.Sprintf("%v%v:", hex.EncodeToString(dstAddrBytes[:i+1]))
		}
		addrStr += "]"
	default:
		panic(atyp)
	}
	dstPortBytes := make([]byte, 2)
	n, err = conn.Read(dstPortBytes)
	if err != nil || n == 0 {
		panic(err)
	}
	dstPort := bytesToint(dstPortBytes)
	switch cmd {
	case 1:
		fmt.Println("CONNECT")
		addr := fmt.Sprintf("%v:%v", addrStr, dstPort)
		targetConn, err := net.Dial("tcp", addr)
		b := make([]byte, 0)
		b = append(b, byte(5)) //VER
		if err != nil {
			b = append(b, byte(1))
			conn.Write(b)
			os.Exit(1)
		}
		b = append(b, byte(0))    //succedd
		b = append(b, byte(0))    //RSV
		b = append(b, byte(atyp)) //ATYP

		addrRemote := strings.Split(targetConn.RemoteAddr().String(), ".")
		a1, _ := strconv.Atoi(addrRemote[0])
		a2, _ := strconv.Atoi(addrRemote[1])
		a3, _ := strconv.Atoi(addrRemote[2])
		a4, _ := strconv.Atoi(addrRemote[3])
		b = append(b, byte(a1), byte(a2), byte(a3), byte(a4))
		b = append(b, intTobytes(dstPort, 2)...)

		conn.Write(b)
		closeConn := make(chan error)
		TransformTraffic(conn, targetConn, closeConn)
		fmt.Println(<-closeConn)
		conn.Close()
	case 2:
		fmt.Println("BIND")
	case 3:
		fmt.Println("UDP ASSOCIATE")
	default:
		panic(cmd)
	}
}
func Server() {
	listener, err := net.Listen("tcp", ":1090")
	if err != nil {
		os.Exit(1)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			os.Exit(1)
		}
		go HandleConn(conn)
	}

}

func Client() {

}
func main() {
	Server()
}
