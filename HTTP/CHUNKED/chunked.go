package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	r, w := io.Pipe()
	m := multipart.NewWriter(b)
	go func() {
		defer w.Close()
		defer m.Close()
		part, err := m.CreateFormFile("file", "test")
		if err != nil {
			fmt.Println(err)
		}
		file, err := os.Open("test")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		if _, err := io.Copy(part, file); err != nil {
			return
		}
	}()
	resp, err := http.Post("http://localhost:8090/api/v1/files", m.FormDataContentType(), r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
