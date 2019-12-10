package main

import (
	"fmt"
)

type JOB struct {
	Id string `json:"id"`
}

func (e *JOB) Execute() {
	fmt.Println("execute")
}
func (e *JOB) Description() string {

	return "one job"
}
func (e *JOB) Key() int {
	return 1
}
func main() {
}
