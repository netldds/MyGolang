package main

import "fmt"

type Mulit map[int]Body

func main() {
	m := make(Mulit)
	m[1] = Body{
		Name:       "",
		OCCUPATION: "",
	}
	fmt.Println(len(m))
}

type Body struct {
	Name       string
	OCCUPATION string
}
