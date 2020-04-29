package main

import (
	address "MyGolang/proto3"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/proto"
)

func main() {
	msg := new(address.Person)
	msg.Name = "jacy~~~~"
	ret, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)

	var resp address.Person
	err = proto.Unmarshal(ret, &resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
