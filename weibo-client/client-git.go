package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"mime/multipart"
	"net/http"
	"net/url"
)

const (
	ClientId     = "f7350ec9a0beb046ccae"
	ClientSecret = "5ee266a0def8e65b58e880a113b74cc506a30616"
	authorizeUrl = "https://github.com/login/oauth/authorize"
	tokenUrl     = "https://github.com/login/oauth/access_token"
	apiURLBase   = "https://api.github.com/"
	CallbackUrl  = "http://127.0.0.1:8080/callback"
)

//code for token
var Code string
var Token string

//发送带client id的请求到github
//重定向地址到本地 带code
//Client带code 等 请求到github 换区token
//Client带Token使用github api
func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	go func() {
		req := apiRequest(authorizeUrl,
			map[string]string{
				"client_id":    ClientId,
				"redirect_uri": CallbackUrl,
			},
			false,
			nil,
		)
		_, err := http.DefaultClient.Do(req)
		if err != nil {
			glog.Info(err)
			return
		}
	}()

	http.HandleFunc("/callback", callbackHandle)
	glog.Info(http.ListenAndServe(":8080", nil))
}
func callbackHandle(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		Code = req.URL.Query().Get("code")
	}
	go GetToken(Code)
	resp.Write([]byte("ok"))
}
func GetToken(code string) {

	req := apiRequest("POST", nil, true, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	rawJson, _ := bufio.NewReader(resp.Body).ReadBytes('\n')
	glog.Info(string(rawJson))
	var unmarshalJson map[string]interface{}
	json.Unmarshal(rawJson, &unmarshalJson)
	Token = unmarshalJson["access_token"].(string)

	glog.Info(Token)

	//test

	u, _ := url.Parse("https://api.github.com/user")

	req = apiRequest(u.String(), nil, false, map[string]string{
		"Authorization": "token" + " " + Token,
		"Accept":        "application/json",
	})
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		glog.Info(err)
	}
	rawJson, _ = bufio.NewReader(resp.Body).ReadBytes('\n')
	glog.Info(string(rawJson))
}
func apiRequest(urls string, parameters map[string]string, post bool, headers map[string]string) *http.Request {
	var err error
	reqUrl, _ := url.Parse(urls)

	var req *http.Request
	if post {
		glog.Info(reqUrl.String())
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		writer.WriteField("client_id", ClientId)
		writer.WriteField("client_secret", ClientSecret)
		writer.WriteField("code", Code)
		writer.Close()
		glog.Info(body.String())
		req, err = http.NewRequest("POST", tokenUrl, body)
		if err != nil {
			glog.Info(req)
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("Accept", "application/json")

	} else {
		values := url.Values{}
		for i, v := range parameters {
			values.Add(i, v)
		}
		reqUrl.RawQuery = values.Encode()
		req, err = http.NewRequest("GET", reqUrl.String(), nil)
		glog.Info(reqUrl.String())
	}
	if err != nil {
		glog.Warning(err)
		return nil
	}
	for i, v := range headers {
		req.Header.Add(i, v)
	}
	return req
}
