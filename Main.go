package main

import (
	"compress/bzip2"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(2)
	}()
	defer func() {
		fmt.Println(3)
	}()

}
func R() {
	name := "tmp001183035"
	f, err := os.OpenFile(name, os.O_RDWR, 0755)
	if err != nil {
		log.Println(err)
	}
	for {
		b := make([]byte, 1024)
		time.Sleep(time.Second)
		n, _ := f.Read(b)
		fmt.Print(string(b[:n]))
	}
}
func W() {
	name := "tmp001183035"
	f, err := os.OpenFile(name, os.O_RDWR, 0755)
	if err != nil {
		log.Println(err)
	}
	for {
		time.Sleep(time.Second)
		s := fmt.Sprintf("%v\n", time.Now().String())
		f.WriteString(s)
		fmt.Print(s)
	}
}

type myFileType struct {
	io.Reader
	io.Closer
}

func OpenFile(name string) io.ReadCloser {

	file, err := os.Open(name)

	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(name, ".gz") {

		gzip, gerr := gzip.NewReader(file)
		if gerr != nil {
			log.Fatal(gerr)
		}
		return gzip

	} else if strings.Contains(name, ".bz2") {

		bzip2 := bzip2.NewReader(file)
		return myFileType{bzip2, file}

	} else {
		return file
	}
}
