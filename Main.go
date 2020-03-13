package main

import (
	"MyGolang/Cache"
	"errors"
	"fmt"
)

type Calendar struct {
	d map[string][]byte
}

func (c *Calendar) Get(key string) ([]byte, error) {
	if v, ok := c.d[key]; ok {
		return v, nil
	}
	return nil, errors.New("failed")
}
func main() {
	data := Calendar{d: map[string][]byte{}}
	data.d["January"] = []byte("一月")
	data.d["February"] = []byte("二月")
	data.d["March"] = []byte("三月")
	data.d["April"] = []byte("四月")

	g := Cache.NewGroup("Calendar", 22, Cache.GetterFunc(data.Get))
	fmt.Println(g.Get("January"))
	fmt.Println(g.Get("February"))
	fmt.Println(g.Get("March"))
	fmt.Println(g.Get("April"))

}
