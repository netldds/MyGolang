package main

import (
	"fmt"
	"time"
)

func main() {
	//var s1 string
	//s2 := ""
	fmt.Println(time.Now().Unix())
	u := time.Now().Unix()
	fmt.Println(time.Unix(u, 0))
}
