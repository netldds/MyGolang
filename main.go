package main

import (
	"fmt"
)

func main() {
	b := []byte{9}
	c := b[0]
	b = nil
	fmt.Println(b)
	fmt.Println(c)

}

func int2bytes(n int, length int) (b []byte) {
	b = make([]byte, length)
	for i := length - 1; i >= 0; i-- {
		b[i] = byte(n & 0xff)
		n = n >> 8
	}
	return
}

//大端
func bytes2int(b []byte) (sum int64) {
	length := len(b)
	for i := length - 1; i >= 0; i-- {
		sum += int64(b[i]) << uint(8*i)
	}
	return
}
