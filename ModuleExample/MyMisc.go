package ModuleExample

import (
	"MyGolang/Struct/Matrix"
	"bufio"
	"fmt"
	"net"
	"os"
	"reflect"
)

func InputLoop() {
	rd := bufio.NewReader(os.Stdin)
	for {
		str, err := rd.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(str)
	}

}
func reflect_example_t() {
	method := func(i string) map[string]string {
		fmt.Println(i)
		return map[string]string{"id": "id1"}
	}
	var v reflect.Value

	v = reflect.ValueOf(method)
	vs := []reflect.Value{reflect.ValueOf("myStrings")}
	rvs := v.Call(vs)
	res := rvs[0].MapIndex(reflect.ValueOf("id")).String()
	fmt.Println(res)
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
				fmt.Printf("server local addr  %v\n server remote addr %v \n", conn.LocalAddr(), conn.RemoteAddr())
				for {
					str, err := rd.ReadString('\n')
					if err != nil {
						fmt.Println(err)
						break
					}
					fmt.Println(str)
				}
			}()
		}
	}()
	//client
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:6666")
	conn, _ := net.DialTCP("tcp", nil, addr)
	fmt.Printf("client local addr %v \n client remote addr %v \n ", conn.LocalAddr(), conn.RemoteAddr())
	conn2 := conn
	fmt.Printf("client local addr %v \n client remote addr %v \n ", conn2.LocalAddr(), conn2.RemoteAddr())

}

type MyError struct {
	id          int
	description string
}

func test() error {
	return &MyError{99, "qq"}
}

func (e *MyError) Error() string {
	e.id++
	return fmt.Sprintf("id:%d  description:%s", e.id, e.description)
}

type tFace interface {
	Prt() string
	Check() bool
}

type tFaceStruct struct {
}

func (s *tFaceStruct) Prt() string {
	return fmt.Sprint("return a string")
}

func (s *tFaceStruct) Check() bool {
	return true
}

//返回接口类型
func testTinterface(t tFace) tFace {

	return &tFaceStruct{}

}

func RunMatrixMutiply() {
	mt := [][]float32{
		[]float32{1, 2, 3},
		[]float32{4, 5, 6},
		[]float32{7, 8, 9}}
	mv := [][]float32{
		{1, 2},
		{1, 2},
		{1, 2}}
	mt1, _ := matrix.New(mt)
	mt2, _ := matrix.New(mv)
	result, _ := mt1.Multiply(mt2)
	fmt.Println(*result)
}
