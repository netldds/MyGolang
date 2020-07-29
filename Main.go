package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type B struct {
	Name string `json:"name"`
}

func (s B) MarshalJSON() ([]byte, error) {
	s.Name = "aaa"
	return json.Marshal(s)
}
func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Unix())
	fmt.Println(t.UTC())
	fmt.Println(t.UTC().Unix())
	stamp := t.UTC().Unix()
	u := time.Unix(stamp, 0)
	fmt.Println(u)
	fmt.Println(u.UTC())
}
