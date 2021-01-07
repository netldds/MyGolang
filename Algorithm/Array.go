package Algorithm

import "fmt"

//数组名地址本身就是数组第一个元素地址
func parseFirstArrayIndex() {
	ar := []byte("1")
	ary := [3]byte{'a', 'b'}
	fmt.Printf("%p\n", &ary[0])
	//ar[0]=2
	fmt.Printf("%p\n", &ary)
	pass(ar)
	fmt.Println(ar)
}
func pass(ar []byte) {
	fmt.Println("pass:", ar)
	ar[0] = 3
	fmt.Println("pass:", ar)
}
