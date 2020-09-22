package main

import (
	"MyGolang/RabbitMQ"
	"fmt"
	"strings"
)

func AA() {
	data := `/p/:id/roles/:rid/permissions`
	slices := strings.Split(data, "/")
	regexpPath := ""
	for _, v := range slices {
		if v == "" {
			continue
		}
		if strings.Contains(v, ":") {
			regexpPath = fmt.Sprintf("%v/.+", regexpPath)
		} else {
			regexpPath = fmt.Sprintf("%v/%v", regexpPath, v)
		}
	}
	regexpPath = fmt.Sprintf("^%v", regexpPath)
	fmt.Println(slices)
	fmt.Println(regexpPath)
}
func main() {
	RabbitMQ.Client()
}
