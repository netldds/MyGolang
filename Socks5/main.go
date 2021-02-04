package main

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
	"os"
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
	bigBytes := make([]byte, len(b))
	for n, v := range b {
		bigBytes[len(b)-n-1] = v
	}
	return bigBytes
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
func TransferTraffic(src, dst net.Conn, close chan error) {
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
	defer conn.Close()
	verByte := make([]byte, 1)
	nmethods := make([]byte, 1)
	/*+----+----------+----------+
	|VER | NMETHODS | METHODS |
	+----+----------+----------+
	| 1 | 1 | 1 to 255 |
	+----+----------+----------+
	*/
	n, err := conn.Read(verByte)
	if err != nil {
		panic(err)
	}
	//版本
	if uint(verByte[0]) != VERSION {
		return
	}
	n, err = conn.Read(nmethods)
	if err != nil {
		return
	}
	methods := []int{}
	for i := 0; i < int(nmethods[0]); i++ {
		method := make([]byte, 1)
		n, err = conn.Read(method)
		if err != nil {
			return
		}
		methods = append(methods, int(method[0]))
	}
	b := make([]byte, 0)
	b = append(b, byte(5))
	//reply
	/*
		+----+--------+
		 |VER | METHOD |
		 +----+--------+
		 | 1 | 1 |
		 +----+--------+
	*/
	for _, v := range methods {
		if v == 0 {
			fmt.Println("NO AUTHEN")
			b = append(b, byte(0))
			conn.Write(b)
			break
		}
		if v == 1 {
			fmt.Println("GSSAPI")
			b = append(b, byte(255))
			conn.Write(b)
			return
		}
		if v == 2 {
			fmt.Println("USERNAME/PASSWORD")
			b = append(b, byte(2))
			//reply
			conn.Write(b)
			closeBytes := make([]byte, 0)
			closeBytes = append(closeBytes, byte(1))
			//验证
			if !VerifyPassword(conn) {
				closeBytes = append(closeBytes, byte(1))
				conn.Write(closeBytes)
				os.Exit(1)
			}
			//successed
			closeBytes = append(closeBytes, byte(0))
			n, err = conn.Write(closeBytes)
			if err != nil || n == 0 {
				return
			}
			break
		}
		b = append(b, ErrMethod)
		conn.Write(b)
		return
	}
	//request
	/*
		+----+-----+-------+------+----------+----------+
		 |VER | CMD | RSV | ATYP | DST.ADDR | DST.PORT |
		 +----+-----+-------+------+----------+----------+
		 | 1 | 1 | X’00’ | 1 | Variable | 2 |
		 +----+-----+-------+------+----------+----------+
	*/
	headBytes := make([]byte, 4)
	n, err = conn.Read(headBytes)
	if err != nil || n == 0 {
		return
	}
	ver, cmd, _, atyp := int(headBytes[0]), int(headBytes[1]), int(headBytes[2]), int(headBytes[3])
	if ver != VERSION {
		return
	}

	var dstIP *net.IP
	domain := ""
	switch atyp {
	case 1:

		fmt.Println("IP V4 address")
		dstAddrBytes := make([]byte, 4)
		n, err = conn.Read(dstAddrBytes)
		if err != nil || n == 0 {
			panic(err)
		}
		d := net.IP(dstAddrBytes)
		dstIP = &d
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
		if err != nil {
			panic(err)
		}
		ipAddr, _ := net.ResolveIPAddr("tcp", addrs[0])
		dstIP = &ipAddr.IP
	case 4:
		fmt.Println("IP V6 address")
		dstAddrBytes := make([]byte, 16)
		n, err = conn.Read(dstAddrBytes)
		if err != nil || n == 0 {
			panic(err)
		}
		d := net.IP(dstAddrBytes)
		dstIP = &d
	default:
		panic(atyp)
	}
	dstPortBytes := make([]byte, 2)
	n, err = conn.Read(dstPortBytes)
	if err != nil || n == 0 {
		panic(err)
	}
	dstPort := bytesToint(dstPortBytes)

	//reply
	/*
		+----+-----+-------+------+----------+----------+
		 |VER | REP | RSV | ATYP | BND.ADDR | BND.PORT |
		 +----+-----+-------+------+----------+----------+
		 | 1 | 1 | X’00’ | 1 | Variable | 2 |
		 +----+-----+-------+------+----------+----------+
	*/
	switch cmd {
	case 1:
		fmt.Println("CONNECT")
		addr := fmt.Sprintf("%v:%v", dstIP.String(), dstPort)
		targetConn, err := net.Dial("tcp", addr)
		b := make([]byte, 0)
		b = append(b, byte(5)) //VER
		if err != nil {
			b = append(b, byte(1))
			conn.Write(b)
			os.Exit(1)
		}
		b = append(b, byte(0))    //REP
		b = append(b, byte(0))    //RSV
		b = append(b, byte(atyp)) //ATYP

		remoteAddr, _ := net.ResolveTCPAddr("tcp", targetConn.RemoteAddr().String())
		b = append(b, byte(remoteAddr.IP[0]))
		b = append(b, byte(remoteAddr.IP[1]))
		b = append(b, byte(remoteAddr.IP[2]))
		b = append(b, byte(remoteAddr.IP[3]))            //BND.ADDR
		b = append(b, intTobytes(remoteAddr.Port, 2)...) //BND.PORT

		conn.Write(b)
		closeConn := make(chan error)
		TransferTraffic(conn, targetConn, closeConn)
		//todo 超时 time_wait
		fmt.Println(<-closeConn)
		conn.Close()
	case 2:
		fmt.Println("BIND")
		//BIND之前需要有CONNECT连接验证
		//建立监听 给目标服务用 例如 FTP 的数据传输
		dstAddr := "0.0.0.0"
		dstPort := rand.Int31n(999) + 1024
		listener, err := net.Listen("tcp", fmt.Sprintf("%v:%v", dstAddr, dstPort))
		//first reply
		b := make([]byte, 0)
		b = append(b, byte(5)) //VER
		if err != nil {
			b = append(b, byte(1))
			conn.Write(b)
			os.Exit(1)
		}
		b = append(b, byte(0))                            //REP
		b = append(b, byte(0))                            //RSV
		b = append(b, byte(atyp))                         //ATYP
		b = append(b, byte(0), byte(0), byte(0), byte(0)) //BND.ADDR
		b = append(b, intTobytes(int(dstPort), 2)...)     //BND.PORT
		conn.Write(b)
		//server -> client
		targetConn, err := listener.Accept()
		//sec reply
		b = make([]byte, 0)
		b = append(b, byte(5)) //VER
		if err != nil {
			b = append(b, byte(1))
			conn.Write(b)
			os.Exit(1)
		}
		b = append(b, byte(0))    //REP
		b = append(b, byte(0))    //RSV
		b = append(b, byte(atyp)) //ATYP
		remoteAddr, _ := net.ResolveTCPAddr("tcp", targetConn.RemoteAddr().String())
		b = append(b, byte(remoteAddr.IP[0]))
		b = append(b, byte(remoteAddr.IP[1]))
		b = append(b, byte(remoteAddr.IP[2]))
		b = append(b, byte(remoteAddr.IP[3]))            //BND.ADDR
		b = append(b, intTobytes(remoteAddr.Port, 2)...) //BND.PORT
		conn.Write(b)
		closeChn := make(chan error)
		TransferTraffic(conn, targetConn, closeChn)
		fmt.Println(<-closeChn)
		targetConn.Close()
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
