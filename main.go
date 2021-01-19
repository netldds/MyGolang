package main

import "fmt"

func main() {
	i := []string{"a", "b", "c", "d"}
	index := 1
	i = append(i[:index], i[index+1:]...)
	fmt.Println(i)
}
