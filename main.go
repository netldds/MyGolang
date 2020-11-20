package main

import "fmt"

func main() {
	mp := make(map[string]interface{})
	mp["a"] = nil
	if _, ok := mp["a"]; ok {
		fmt.Println("ok")
	}
}
