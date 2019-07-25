package main

import (
	"MyGolang/Misc"
	"fmt"
)

/*
#cgo CFLAGS: -I${SRCDIR}/libs/include
#cgo LDFLAGS: ${SRCDIR}/libs/static_lib.a -lstdc++
#include "static_lib.h"
#include <stdlib.h>
*/
import "C"

//https://golang.org/cmd/cgo/
func main() {
	num := C.add(1, 2)
	fmt.Println(num)
	Misc.Add(1, 2)
}
