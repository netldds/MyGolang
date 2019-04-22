package main

import (
	"flag"
	"github.com/golang/glog"
	"net/http"
	"net/url"
)

const (
	ClientId     = "f7350ec9a0beb046ccae"
	ClientSecret = "5ee266a0def8e65b58e880a113b74cc506a30616"
	authorizeUrl = "https://github.com/login/oauth/authorize"
	tokenUrl     = "https://github.com/login/oauth/access_token"
	apiURLBase   = "https://api.github.com/"
	CallbackUrl  = "http://127.0.0.1/callback"
)

var Code string

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	req := apiRequest(authorizeUrl,
		map[string]string{
			"client_id":    ClientId,
			"redirect_uri": CallbackUrl,
		},
		false,
		nil,
	)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		glog.Warning(err)
	}
	defer resp.Body.Close()
	//data, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	glog.Warning(err)
	//}
	http.HandleFunc("/callback", callbackHandle)
	glog.Info(http.ListenAndServe(":80", nil))
}
func callbackHandle(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		Code = req.URL.Query().Get("code")
	}
	resp.Write([]byte("ok"))
}
func apiRequest(urls string, parameters map[string]string, post bool, headers map[string]string) *http.Request {
	var err error
	reqUrl, _ := url.Parse(urls)
	values := url.Values{}
	for i, v := range parameters {
		values.Add(i, v)
	}
	reqUrl.RawQuery = values.Encode()
	if err != nil {
		glog.Warning(err)
		return nil
	}

	var req *http.Request
	if post {
		glog.Info(reqUrl.String())
		req, err = http.NewRequest("POST", reqUrl.String(), nil)
	} else {
		glog.Info(reqUrl.String())
		req, err = http.NewRequest("GET", reqUrl.String(), nil)
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
