package main

import (
	"fmt"
	"math/rand"
)

func main() {
	T := &BST{Key: 80}
	for i := 0; i < 10; i++ {
		TreeInsert(T, &BST{
			Key: int(rand.Int31n(99)),
		})
	}
	TreeDelete(T, T.Left.Right)
	fmt.Print(T)
}
