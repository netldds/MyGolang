package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}

func checkErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
func main() {
	fmt.Println(os.Getwd())

	//target_url := "http://127.0.0.1:9090/upload"
	//filename := "t.dat"
	//fmt.Println(postFile(filename, target_url))

	rpcClient()
}
func postFile(filename string, target_url string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer.")
		return err
	}
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, err := http.Post(target_url, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}
func rpcClient() {

	//RPC
	//=============================
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:80")
	checkErr(err)

	var reply int
	args := &Args{A: 2, B: 2}
	err = client.Call("Arith.Multiply", args, &reply)
	fmt.Printf("Arith: %d*%d=%d\n", 2, 2, reply)
	checkErr(err)

	var reply2 int = 1
	i := 5
	err = client.Call("Brith.Add", &i, &reply2)
	fmt.Printf("Brith: %d+%d=%d\n", i, 11, reply)
	//=============================
}
