package main

//#include "worker.h"
import "C"
import "fmt"

func main() {

	num := 10
	added := C.add(num, 2)
	fmt.Println(added)

}
