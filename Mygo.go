package main

import (
	"MyGolang/DataBaseOperation"
	"bufio"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net"
	"os"
	"reflect"
	"time"
)

const (
	file1 = "6.9MB.jpg"
	file2 = "26MB.jpg"
	file3 = "1MB.gif"
	file4 = "3.7MB.tga"
	file5 = "13.4MB.pdf"
)

var counterPool = make(map[string]time.Time)

const rootpath = "/home/dx/GoWorkBench/src/dx/taishan/data/comment_files"
const hextable = "0123456789abcdef"

const (
	ChildStatus = 1 << (32 - 1 - iota)
	MasterStatus
	BothStatus        = ChildStatus | MasterStatus
	UnavailableStatus = -1
)

type Item struct {
	Value int
}

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()
	DataBaseOperation.RungOrm()



}

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
func reflect_example() {
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

const data = `
{"aloneEnabled":false,"antialiasEnabled":true,"backgroundColor":[1,1,1],"camera":{"inOrthographicMode":false,"inPerspectiveMode":true,"position":{"x":23.439013672736447,"y":-12.438647461798944,"z":20.09784407701805},"up":{"x":0,"y":0,"z":1},"zoom":1},"cameraTarget":{"x":6.000091552734375,"y":5.000274658203125,"z":2.658921957015991},"canvas":{"objects":[{"angle":0,"backgroundColor":"","clipTo":null,"fill":"rgb(0,0,0)","fillRule":"nonzero","flipX":false,"flipY":false,"globalCompositeOperation":"source-over","height":115.3,"left":813.25,"objects":[{"angle":0,"backgroundColor":"","clipTo":null,"fill":false,"fillRule":"nonzero","flipX":false,"flipY":false,"globalCompositeOperation":"source-over","height":87,"left":6.75,"opacity":1,"originX":"center","originY":"center","rx":0,"ry":0,"scaleX":1,"scaleY":1,"shadow":null,"skewX":0,"skewY":0,"stroke":"#f00","strokeDashArray":null,"strokeLineCap":"butt","strokeLineJoin":"miter","strokeMiterLimit":10,"strokeWidth":3,"top":12.65,"transformMatrix":null,"type":"rect","visible":true,"width":156},{"angle":0,"backgroundColor":"","charSpacing":0,"clipTo":null,"fill":"#f00","fillRule":"nonzero","flipX":false,"flipY":false,"fontFamily":"Times New Roman","fontSize":20,"fontStyle":"","fontWeight":"normal","globalCompositeOperation":"source-over","height":22.6,"left":-86.25,"lineHeight":1.16,"opacity":1,"originX":"left","originY":"center","scaleX":1,"scaleY":1,"shadow":null,"skewX":0,"skewY":0,"stroke":null,"strokeDashArray":null,"strokeLineCap":"butt","strokeLineJoin":"miter","strokeMiterLimit":10,"strokeWidth":1,"text":"1","textAlign":"left","textBackgroundColor":"","textDecoration":"","top":-45.85,"transformMatrix":null,"type":"text","visible":true,"width":10}],"opacity":1,"originX":"center","originY":"center","scaleX":1,"scaleY":1,"shadow":null,"skewX":0,"skewY":0,"stroke":null,"strokeDashArray":null,"strokeLineCap":"butt","strokeLineJoin":"miter","strokeMiterLimit":10,"strokeWidth":0,"top":291.85,"transformMatrix":null,"type":"group","visible":true,"width":172.5},{"angle":0,"backgroundColor":"","clipTo":null,"fill":"rgb(0,0,0)","fillRule":"nonzero","flipX":false,"flipY":false,"globalCompositeOperation":"source-over","height":128.3,"left":480.25,"objects":[{"angle":0,"backgroundColor":"","clipTo":null,"fill":false,"fillRule":"nonzero","flipX":false,"flipY":false,"globalCompositeOperation":"source-over","height":100,"left":6.75,"opacity":1,"originX":"center","originY":"center","rx":0,"ry":0,"scaleX":1,"scaleY":1,"shadow":null,"skewX":0,"skewY":0,"stroke":"#f00","strokeDashArray":null,"strokeLineCap":"butt","strokeLineJoin":"miter","strokeMiterLimit":10,"strokeWidth":3,"top":12.65,"transformMatrix":null,"type":"rect","visible":true,"width":116},{"angle":0,"backgroundColor":"","charSpacing":0,"clipTo":null,"fill":"#f00","fillRule":"nonzero","flipX":false,"flipY":false,"fontFamily":"Times New Roman","fontSize":20,"fontStyle":"","fontWeight":"normal","globalCompositeOperation":"source-over","height":22.6,"left":-66.25,"lineHeight":1.16,"opacity":1,"originX":"left","originY":"center","scaleX":1,"scaleY":1,"shadow":null,"skewX":0,"skewY":0,"stroke":null,"strokeDashArray":null,"strokeLineCap":"butt","strokeLineJoin":"miter","strokeMiterLimit":10,"strokeWidth":1,"text":"2","textAlign":"left","textBackgroundColor":"","textDecoration":"","top":-52.35,"transformMatrix":null,"type":"text","visible":true,"width":10}],"opacity":1,"originX":"center","originY":"center","scaleX":1,"scaleY":1,"shadow":null,"skewX":0,"skewY":0,"stroke":null,"strokeDashArray":null,"strokeLineCap":"butt","strokeLineJoin":"miter","strokeMiterLimit":10,"strokeWidth":0,"top":426.35,"transformMatrix":null,"type":"group","visible":true,"width":132.5}]},"curBackgroundIndex":"color_0","curWhiteBalance":0.65,"height":777,"hiddenIds":"[]","hiddenTreeIds":"{}","isSetting":false,"isWebgl":true,"lastSectionBox":[-3,-3,-1,15.00018310546875,13.00054931640625,6.317843914031982],"markupEnabled":true,"outlineEnabled":false,"path":"taishan/8fde0d690ebd4f2186c0e58e9614a44f/2df8c864c900428ab856b9177ef6ea30/0817be36dee84146a21ebcc4eee7bdbb.ifc","pivot":{"x":6.000091552734375,"y":5.000274658203125,"z":2.658921957015991},"positions":[{"flag":7,"point":[6.75520630540732,3.75652129528034,5.599975372473269]},{"flag":4,"point":[-1.6048683728586965,-3.0000000000000018,-0.5766086093877547]}],"projectionType":"Pers","propertyEnabled":false,"sceneBBox":[-3,-3,-1,15.00018310546875,13.00054931640625,6.317843914031982],"selectionDisplayMode":0,"selectionIds":"[]","selectionTreeIds":"{}","shadowEnabled":false,"treeEnabled":false,"width":1471}

`
