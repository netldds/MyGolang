package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type File string

func (f *File) String() {
	fmt.Printf("%s + end.", *f)
}

//https://golang.org/cmd/cgo/
func main() {
	dir, _ := os.Getwd()
	f, err := ioutil.TempFile(dir, "PATTERN")
	fmt.Println(err)
	f.Write([]byte("abc"))
	fmt.Println(filepath.Base(f.Name()))
	f.Close()
}
