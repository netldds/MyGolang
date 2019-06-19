package main

import (
	"flag"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"time"
	"unsafe"

	_ "github.com/go-sql-driver/mysql"
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

type Item struct {
	Value int
	Name  string
}

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	str := make([]string, 2)
	fmt.Printf("%p\n", &str[0])
	fmt.Printf("%p\n", &str[1])

	items := make([]Item, 2)
	items[1] = Item{Name: "xxx"}
	fmt.Printf("%T,%d\n", items, unsafe.Sizeof(items))

	fmt.Printf("%p\n", &items[0])
	fmt.Printf("%p\n", &items[1])

	i1Ptr := unsafe.Pointer(&items[0])
	fmt.Println(i1Ptr)
	i2Ptr := unsafe.Pointer(uintptr(i1Ptr) + uintptr(24))
	fmt.Println(i2Ptr)

	fmt.Println(*(*Item)(i2Ptr))

}
