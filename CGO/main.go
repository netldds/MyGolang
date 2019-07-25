package main

import "fmt"

/*
//头文件
#cgo CFLAGS: -I${SRCDIR}/
//库文件
#cgo LDFLAGS: ${SRCDIR}/worker.a
#include "worker.h"
*/
import "C"

func main() {
	added := C.add(1, 2)
	fmt.Println(added)
}
