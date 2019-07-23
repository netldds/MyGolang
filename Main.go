package main

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"net/http"
	"time"
)
/*
#cgo CFLAGS: -I${SRCDIR}/libs
#cgo LDFLAGS: ${SRCDIR}/libs/static_lib.a -lstdc++
#include "static_lib.h"
#include <stdlib.h>
*/
import "C"
//https://golang.org/cmd/cgo/
func main() {
	num := C.add(1, 2)
	fmt.Println(num)
}
func tt() {
	http.HandleFunc("/", func(res http.ResponseWriter, request *http.Request) {
		res.Write([]byte("ok\n"))
		res.WriteHeader(200)
	})
	srv := http.Server{
		Addr:    ":9000",
		Handler: http.DefaultServeMux,
	}
	//ctx, cancel := context.WithCancel(context.Background())
	go func() {
		//ctx,cancel:=context.WithTimeout(context.Background(), time.Second*2)
		ticker := time.NewTicker(time.Second * 5)
		for {
			select {
			case t := <-ticker.C:
				fmt.Println(t)
				if err := srv.Shutdown(context.Background()); err != nil {
					// Error from closing listeners, or context timeout:
					glog.Infof("HTTP server Shutdown: %v", err)
				}
			}
		}
	}()
	srv.ListenAndServe()
}
